import createClient, { type Client } from "openapi-fetch";
import type { paths, components } from "./generated/schema.js";
import { CMSError, NotFoundError } from "./errors.js";
import { paginateSequential, paginateParallel } from "./paginate.js";

export type FetchLike = typeof fetch;

export type Item = components["schemas"]["versionedItem"];
export type Field = components["schemas"]["field"];
export type Model = components["schemas"]["model"];
export type Asset = components["schemas"]["asset"];
export type Comment = components["schemas"]["comment"];
export type AssetEmbedding = components["schemas"]["assetEmbedding"];

export interface AssetUpload {
  url?: string;
  token?: string;
  contentType?: string;
  contentEncoding?: string;
  contentLength?: number;
  next?: string;
}

/** Acceptable body types for asset data. */
export type AssetData = Blob | ArrayBuffer | ArrayBufferView | ReadableStream<Uint8Array> | string;

/** Compression algorithms supported by the browser's `CompressionStream`. */
export type Compression = "gzip" | "deflate" | "deflate-raw";

/**
 * Compress data and return a `ReadableStream` of compressed bytes.
 * Streams end-to-end — no intermediate buffering.
 *
 * Available on Cloudflare Workers, Node.js 18+, Bun, and Deno.
 */
export function compressToStream(data: AssetData, algo: Compression): ReadableStream<Uint8Array> {
  const stream = new Response(data as BodyInit).body;
  if (!stream) throw new Error("failed to read data stream");
  return stream.pipeThrough(new CompressionStream(algo));
}

/**
 * Compress data and buffer the result into a `Blob`.
 * Use this when the consumer needs a `Blob` (e.g. multipart form parts).
 * For raw PUT uploads, prefer `compressToStream` to avoid buffering.
 */
export async function compressData(data: AssetData, algo: Compression): Promise<Blob> {
  return new Response(compressToStream(data, algo)).blob();
}

function isReadableStream(x: unknown): x is ReadableStream<Uint8Array> {
  return typeof ReadableStream !== "undefined" && x instanceof ReadableStream;
}

export interface CMSOptions {
  /** Base URL of the Re:Earth CMS server. `/api` is appended automatically. */
  baseURL: string;
  /** Bearer token for the Integration API. */
  token: string;
  /**
   * Workspace ID or alias. When set, the client uses the new workspace-scoped
   * API (`/api/{workspace}/projects/...`). When omitted, the client falls back
   * to the legacy API (`/api/...`) for compatibility with older servers.
   */
  workspace?: string;
  /** Default project ID or alias. Can be overridden per call. */
  project?: string;
  /** Request timeout in milliseconds. */
  timeout?: number;
  /** Custom fetch implementation (defaults to global `fetch`). */
  fetch?: FetchLike;
}

export interface ListResponse<T> {
  results: T[];
  page: number;
  perPage: number;
  totalCount: number;
}

export interface GetItemOptions {
  project?: string;
  /** Required on the new API (when `workspace` is set). Ignored on the legacy API. */
  model?: string;
  itemId: string;
  asset?: boolean | AssetEmbedding;
  ref?: "latest" | "public";
}

export interface GetItemsOptions {
  project?: string;
  model: string;
  asset?: boolean | AssetEmbedding;
  ref?: "latest" | "public";
  keyword?: string;
  sort?: "createdAt" | "updatedAt";
  dir?: "asc" | "desc";
}

export interface GetItemsPageOptions extends GetItemsOptions {
  page?: number;
  perPage?: number;
}

export interface GetAllItemsOptions extends GetItemsOptions {
  perPage?: number;
  /** When set, fetch pages in parallel with the given concurrency. */
  concurrency?: number;
}

export interface CreateItemOptions {
  project?: string;
  model: string;
  fields?: Field[];
  metadataFields?: Field[];
}

export interface UpdateItemOptions {
  project?: string;
  /** Required on the new API (when `workspace` is set). Ignored on the legacy API. */
  model?: string;
  itemId: string;
  fields?: Field[];
  metadataFields?: Field[];
  asset?: AssetEmbedding;
}

export interface DeleteItemOptions {
  project?: string;
  /** Required on the new API (when `workspace` is set). Ignored on the legacy API. */
  model?: string;
  itemId: string;
}

export interface GetModelOptions {
  /**
   * Project ID or alias. Required on the new API. On the legacy API it is
   * optional: when omitted, the flat `/api/models/{modelId}` endpoint is used
   * (model must be addressed by ID, not key).
   */
  project?: string;
  modelIdOrKey: string;
}

export interface GetModelsPageOptions {
  project?: string;
  page?: number;
  perPage?: number;
  keyword?: string;
  sort?: "createdAt" | "updatedAt";
  dir?: "asc" | "desc";
}

export interface GetAllModelsOptions {
  project?: string;
  perPage?: number;
  concurrency?: number;
  keyword?: string;
  sort?: "createdAt" | "updatedAt";
  dir?: "asc" | "desc";
}

export interface CommentOptions {
  project?: string;
  /** Required on the new API (when `workspace` is set). Ignored on the legacy API. */
  model?: string;
  itemId: string;
  content: string;
}

export interface CommentToAssetOptions {
  project?: string;
  assetId: string;
  content: string;
}

export interface GetAssetOptions {
  project?: string;
  assetId: string;
}

export interface UploadAssetByURLOptions {
  project?: string;
  url: string;
  contentEncoding?: string;
  skipDecompression?: boolean;
}

export interface UploadAssetDirectlyOptions {
  project?: string;
  /** The file data. Blob is recommended because it carries a filename and type. */
  file: Blob | AssetData;
  /** Filename. Required unless `file` is a `File` with a name. */
  name?: string;
  contentType?: string;
  contentEncoding?: string;
  skipDecompression?: boolean;
  /**
   * Compress the data client-side before uploading and automatically set
   * `contentEncoding` to match. Ignored if `contentEncoding` is already set.
   */
  compress?: Compression;
}

export interface CreateAssetUploadOptions {
  project?: string;
  name: string;
  contentType?: string;
  contentEncoding?: string;
  contentLength?: number;
  cursor?: string;
}

export interface CreateAssetByTokenOptions {
  project?: string;
  token: string;
  contentEncoding?: string;
  skipDecompression?: boolean;
}

export interface UploadToAssetUploadOptions {
  /**
   * Compress the data client-side before uploading. The `Content-Encoding`
   * header is derived from `upload.contentEncoding` set at `createAssetUpload`
   * time; pass `compress` here to match (or override) it.
   */
  compress?: Compression;
}

type Query = Record<string, string | number | boolean | undefined>;

interface SendOptions {
  body?: unknown;
  query?: Query;
  contentType?: string;
}

export class CMS {
  readonly options: Readonly<CMSOptions>;
  /**
   * Low-level typed OpenAPI client. Requires `workspace` to be set because
   * the schema only describes the new workspace-scoped API.
   */
  readonly api: Client<paths>;

  constructor(options: CMSOptions) {
    if (!options.baseURL) throw new Error("baseURL is required");
    if (!options.token) throw new Error("token is required");

    this.options = { ...options };

    const baseUrl = options.baseURL.replace(/\/+$/, "") + "/api";
    this.api = createClient<paths>({
      baseUrl,
      fetch: options.fetch,
      headers: { Authorization: `Bearer ${options.token}` },
    });
  }

  /** Returns a new CMS instance with the given project as default. */
  withProject(project: string): CMS {
    return new CMS({ ...this.options, project });
  }

  /** Returns a new CMS instance with the given workspace. */
  withWorkspace(workspace: string): CMS {
    return new CMS({ ...this.options, workspace });
  }

  /** Returns a new CMS instance with the given timeout. */
  withTimeout(timeout: number): CMS {
    return new CMS({ ...this.options, timeout });
  }

  /** True when `workspace` is set; false uses the legacy API. */
  useNewAPI(): boolean {
    return Boolean(this.options.workspace);
  }

  protected requireProject(project?: string): string {
    const p = project ?? this.options.project;
    if (!p) throw new Error("project is required (pass to method or CMS constructor)");
    return p;
  }

  protected requireModel(model?: string): string {
    if (!model) throw new Error("model is required on the new API (set via CMS.workspace)");
    return model;
  }

  protected signal(): AbortSignal | undefined {
    if (!this.options.timeout || this.options.timeout <= 0) return undefined;
    return AbortSignal.timeout(this.options.timeout);
  }

  protected assetParam(v: boolean | AssetEmbedding | undefined): AssetEmbedding | undefined {
    if (v === undefined) return undefined;
    if (typeof v === "boolean") return v ? "true" : "false";
    return v;
  }

  // ---------- Path builders (new vs legacy API) ----------

  private pathModel(modelIdOrKey: string, project?: string): string[] {
    if (this.useNewAPI()) {
      return [
        "api",
        this.options.workspace!,
        "projects",
        this.requireProject(project),
        "models",
        modelIdOrKey,
      ];
    }
    const p = project ?? this.options.project;
    if (p) return ["api", "projects", p, "models", modelIdOrKey];
    return ["api", "models", modelIdOrKey];
  }

  private pathModels(project: string): string[] {
    if (this.useNewAPI()) {
      return ["api", this.options.workspace!, "projects", project, "models"];
    }
    return ["api", "projects", project, "models"];
  }

  private pathItems(model: string, project?: string): string[] {
    if (this.useNewAPI()) {
      return [
        "api",
        this.options.workspace!,
        "projects",
        this.requireProject(project),
        "models",
        model,
        "items",
      ];
    }
    const p = project ?? this.options.project;
    if (p) return ["api", "projects", p, "models", model, "items"];
    return ["api", "models", model, "items"];
  }

  private pathItem(itemId: string, model?: string, project?: string): string[] {
    if (this.useNewAPI()) {
      return [
        "api",
        this.options.workspace!,
        "projects",
        this.requireProject(project),
        "models",
        this.requireModel(model),
        "items",
        itemId,
      ];
    }
    return ["api", "items", itemId];
  }

  private pathItemComments(itemId: string, model?: string, project?: string): string[] {
    return [...this.pathItem(itemId, model, project), "comments"];
  }

  private pathAssets(project: string): string[] {
    if (this.useNewAPI()) {
      return ["api", this.options.workspace!, "projects", project, "assets"];
    }
    return ["api", "projects", project, "assets"];
  }

  private pathAssetUploads(project: string): string[] {
    return [...this.pathAssets(project), "uploads"];
  }

  private pathAsset(assetId: string, project?: string): string[] {
    if (this.useNewAPI()) {
      return [
        "api",
        this.options.workspace!,
        "projects",
        this.requireProject(project),
        "assets",
        assetId,
      ];
    }
    return ["api", "assets", assetId];
  }

  private pathAssetComments(assetId: string, project?: string): string[] {
    return [...this.pathAsset(assetId, project), "comments"];
  }

  // ---------- Raw request ----------

  protected async send<T>(method: string, segs: string[], opts: SendOptions = {}): Promise<T> {
    const base = this.options.baseURL.replace(/\/+$/, "");
    const url = new URL(base + "/" + segs.map((s) => encodeURIComponent(s)).join("/"));
    if (opts.query) {
      for (const [k, v] of Object.entries(opts.query)) {
        if (v !== undefined) url.searchParams.set(k, String(v));
      }
    }

    const headers: Record<string, string> = {
      Authorization: `Bearer ${this.options.token}`,
    };

    let body: BodyInit | undefined;
    let duplex: "half" | undefined;
    const b = opts.body;
    if (b !== undefined) {
      if (b instanceof FormData) {
        // let fetch set the multipart boundary
        body = b;
      } else if (isReadableStream(b)) {
        body = b;
        duplex = "half";
        if (opts.contentType) headers["Content-Type"] = opts.contentType;
      } else if (b instanceof Blob) {
        body = b;
        if (opts.contentType) headers["Content-Type"] = opts.contentType;
      } else if (b instanceof ArrayBuffer || ArrayBuffer.isView(b)) {
        body = b as BodyInit;
        if (opts.contentType) headers["Content-Type"] = opts.contentType;
      } else {
        body = JSON.stringify(b);
        headers["Content-Type"] = opts.contentType ?? "application/json";
      }
    }

    const init: RequestInit & { duplex?: "half" } = {
      method,
      headers,
      body,
      signal: this.signal(),
    };
    if (duplex) init.duplex = duplex;

    const fetchImpl = this.options.fetch ?? fetch;
    const res = await fetchImpl(url.toString(), init);
    if (!res.ok) {
      if (res.status === 404) throw new NotFoundError();
      const text = await res.text().catch(() => undefined);
      throw new CMSError(res.status, `request failed: status=${res.status}`, text);
    }
    const ct = res.headers.get("content-type") ?? "";
    if (!ct.includes("application/json") || res.status === 204) {
      return undefined as unknown as T;
    }
    return (await res.json()) as T;
  }

  // ---------- Models ----------

  async getModel(opts: GetModelOptions): Promise<Model> {
    return this.send<Model>("GET", this.pathModel(opts.modelIdOrKey, opts.project));
  }

  async getModelsPage(opts: GetModelsPageOptions = {}): Promise<ListResponse<Model>> {
    const data = await this.send<{
      models?: Model[];
      page?: number;
      perPage?: number;
      totalCount?: number;
    }>("GET", this.pathModels(this.requireProject(opts.project)), {
      query: {
        page: opts.page,
        perPage: opts.perPage,
        keyword: opts.keyword,
        sort: opts.sort,
        dir: opts.dir,
      },
    });
    return {
      results: data.models ?? [],
      page: data.page ?? opts.page ?? 1,
      perPage: data.perPage ?? opts.perPage ?? 0,
      totalCount: data.totalCount ?? 0,
    };
  }

  async getAllModels(opts: GetAllModelsOptions = {}): Promise<ListResponse<Model>> {
    const perPage = opts.perPage ?? 100;
    const fetchPage = (page: number, per: number) =>
      this.getModelsPage({ ...opts, page, perPage: per });

    if (opts.concurrency && opts.concurrency > 1) {
      return paginateParallel(fetchPage, perPage, opts.concurrency);
    }
    return paginateSequential(fetchPage, perPage);
  }

  // ---------- Items ----------

  async getItem(opts: GetItemOptions): Promise<Item> {
    return this.send<Item>("GET", this.pathItem(opts.itemId, opts.model, opts.project), {
      query: { asset: this.assetParam(opts.asset), ref: opts.ref },
    });
  }

  async getItemsPage(opts: GetItemsPageOptions): Promise<ListResponse<Item>> {
    const data = await this.send<{
      items?: Item[];
      page?: number;
      perPage?: number;
      totalCount?: number;
    }>("GET", this.pathItems(opts.model, opts.project), {
      query: {
        page: opts.page,
        perPage: opts.perPage,
        asset: this.assetParam(opts.asset),
        ref: opts.ref,
        keyword: opts.keyword,
        sort: opts.sort,
        dir: opts.dir,
      },
    });
    return {
      results: data.items ?? [],
      page: data.page ?? opts.page ?? 1,
      perPage: data.perPage ?? opts.perPage ?? 0,
      totalCount: data.totalCount ?? 0,
    };
  }

  async getAllItems(opts: GetAllItemsOptions): Promise<ListResponse<Item>> {
    const perPage = opts.perPage ?? 100;
    const fetchPage = (page: number, per: number) =>
      this.getItemsPage({ ...opts, page, perPage: per });

    if (opts.concurrency && opts.concurrency > 1) {
      return paginateParallel(fetchPage, perPage, opts.concurrency);
    }
    return paginateSequential(fetchPage, perPage);
  }

  async createItem(opts: CreateItemOptions): Promise<Item> {
    return this.send<Item>("POST", this.pathItems(opts.model, opts.project), {
      body: { fields: opts.fields, metadataFields: opts.metadataFields },
    });
  }

  async updateItem(opts: UpdateItemOptions): Promise<Item> {
    return this.send<Item>("PATCH", this.pathItem(opts.itemId, opts.model, opts.project), {
      body: {
        fields: opts.fields,
        metadataFields: opts.metadataFields,
        asset: opts.asset,
      },
    });
  }

  async deleteItem(opts: DeleteItemOptions): Promise<void> {
    await this.send<void>("DELETE", this.pathItem(opts.itemId, opts.model, opts.project));
  }

  // ---------- Assets ----------

  async getAsset(opts: GetAssetOptions): Promise<Asset> {
    return this.send<Asset>("GET", this.pathAsset(opts.assetId, opts.project));
  }

  /**
   * Upload an asset from a public URL. The server fetches and stores it.
   */
  async uploadAsset(opts: UploadAssetByURLOptions): Promise<Asset> {
    return this.send<Asset>("POST", this.pathAssets(this.requireProject(opts.project)), {
      body: {
        url: opts.url,
        contentEncoding: opts.contentEncoding,
        skipDecompression: opts.skipDecompression ?? false,
      },
    });
  }

  /**
   * Upload an asset directly via multipart/form-data.
   * Use this for small-to-medium files. For larger files, prefer
   * `createAssetUpload` + `uploadToAssetUpload` + `createAssetByToken`.
   */
  async uploadAssetDirectly(opts: UploadAssetDirectlyOptions): Promise<Asset> {
    const form = new FormData();
    if (opts.contentType) form.append("contentType", opts.contentType);
    if (opts.skipDecompression) form.append("skipDecompression", "true");

    const name = opts.name ?? (opts.file instanceof File ? opts.file.name : undefined);
    if (!name) throw new Error("name is required when file is not a File");

    // multipart/form-data requires a Blob (known size).
    // A ReadableStream input must be buffered here — use the 2-stage upload
    // (createAssetUpload + uploadToAssetUpload) for true streaming.
    let blob: Blob;
    if (opts.file instanceof Blob) {
      blob = opts.file;
    } else if (isReadableStream(opts.file)) {
      blob = await new Response(opts.file).blob();
    } else {
      blob = new Blob([opts.file as BlobPart]);
    }

    let contentEncoding = opts.contentEncoding;
    if (opts.compress && !contentEncoding) {
      blob = await compressData(blob, opts.compress);
      contentEncoding = opts.compress;
    }
    if (contentEncoding) form.append("contentEncoding", contentEncoding);

    form.append("file", blob, name);

    return this.send<Asset>("POST", this.pathAssets(this.requireProject(opts.project)), {
      body: form,
    });
  }

  /**
   * Create a 2-stage asset upload. The returned `url` accepts a PUT request
   * with the raw body. After uploading, call `createAssetByToken` with the
   * returned `token` to register the asset.
   */
  async createAssetUpload(opts: CreateAssetUploadOptions): Promise<AssetUpload> {
    return this.send<AssetUpload>(
      "POST",
      this.pathAssetUploads(this.requireProject(opts.project)),
      {
        body: {
          name: opts.name,
          contentType: opts.contentType,
          contentEncoding: opts.contentEncoding,
          contentLength: opts.contentLength,
          cursor: opts.cursor,
        },
      },
    );
  }

  /**
   * Upload the raw body to the pre-signed URL returned by `createAssetUpload`.
   * Uses the injected `fetch` if provided (otherwise global `fetch`).
   */
  async uploadToAssetUpload(
    upload: AssetUpload,
    body: AssetData,
    opts: UploadToAssetUploadOptions = {},
  ): Promise<void> {
    if (!upload.url) throw new Error("upload.url is empty");

    // Stream end-to-end: compression pipes through without buffering.
    let data: BodyInit = body as BodyInit;
    if (opts.compress) {
      data = compressToStream(body, opts.compress);
    }

    const headers: Record<string, string> = {};
    if (upload.contentType) headers["Content-Type"] = upload.contentType;
    if (upload.contentEncoding) headers["Content-Encoding"] = upload.contentEncoding;

    // Node's fetch (undici) requires `duplex: "half"` when the body is a stream.
    // Cloudflare Workers / Bun / Deno accept it or tolerate it.
    const init: RequestInit & { duplex?: "half" } = {
      method: "PUT",
      headers,
      body: data,
      signal: this.signal(),
    };
    if (isReadableStream(data)) init.duplex = "half";

    const fetchImpl = this.options.fetch ?? fetch;
    const res = await fetchImpl(upload.url, init);
    if (!res.ok) {
      const text = await res.text().catch(() => undefined);
      throw new CMSError(
        res.status,
        `failed to upload to presigned URL: status=${res.status}`,
        text,
      );
    }
  }

  /**
   * Register an asset after uploading to the pre-signed URL.
   */
  async createAssetByToken(opts: CreateAssetByTokenOptions): Promise<Asset> {
    if (!opts.token) throw new Error("token is required");
    return this.send<Asset>("POST", this.pathAssets(this.requireProject(opts.project)), {
      body: {
        token: opts.token,
        contentEncoding: opts.contentEncoding,
        skipDecompression: opts.skipDecompression ?? false,
      },
    });
  }

  // ---------- Comments ----------

  async commentToItem(opts: CommentOptions): Promise<Comment> {
    return this.send<Comment>(
      "POST",
      this.pathItemComments(opts.itemId, opts.model, opts.project),
      { body: { content: opts.content } },
    );
  }

  async commentToAsset(opts: CommentToAssetOptions): Promise<Comment> {
    return this.send<Comment>("POST", this.pathAssetComments(opts.assetId, opts.project), {
      body: { content: opts.content },
    });
  }
}

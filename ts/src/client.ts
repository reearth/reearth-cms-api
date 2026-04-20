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
  /** Workspace ID or alias. Required for the new (workspace-scoped) API. */
  workspace: string;
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
  model: string;
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
  model: string;
  itemId: string;
  fields?: Field[];
  metadataFields?: Field[];
  asset?: AssetEmbedding;
}

export interface DeleteItemOptions {
  project?: string;
  model: string;
  itemId: string;
}

export interface GetModelOptions {
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
  model: string;
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

export class CMS {
  readonly options: Readonly<CMSOptions>;
  /** Low-level typed OpenAPI client. Exposed for advanced use cases. */
  readonly api: Client<paths>;

  constructor(options: CMSOptions) {
    if (!options.baseURL) throw new Error("baseURL is required");
    if (!options.token) throw new Error("token is required");
    if (!options.workspace) throw new Error("workspace is required");

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

  /** Returns a new CMS instance with the given timeout. */
  withTimeout(timeout: number): CMS {
    return new CMS({ ...this.options, timeout });
  }

  protected requireProject(project?: string): string {
    const p = project ?? this.options.project;
    if (!p) throw new Error("project is required (pass to method or CMS constructor)");
    return p;
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

  protected throwIfError(response: Response, error: unknown): void {
    if (response.ok) return;
    const bodyStr =
      error === undefined ? undefined : typeof error === "string" ? error : JSON.stringify(error);
    if (response.status === 404) throw new NotFoundError(bodyStr);
    throw new CMSError(
      response.status,
      `request failed: status=${response.status}${bodyStr ? ` body=${bodyStr}` : ""}`,
      bodyStr,
    );
  }

  // ---------- Models ----------

  async getModel(opts: GetModelOptions): Promise<Model> {
    const { data, error, response } = await this.api.GET(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            modelIdOrKey: opts.modelIdOrKey,
          },
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  async getModelsPage(opts: GetModelsPageOptions = {}): Promise<ListResponse<Model>> {
    const { data, error, response } = await this.api.GET(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
          },
          query: {
            page: opts.page,
            perPage: opts.perPage,
            keyword: opts.keyword,
            sort: opts.sort,
            dir: opts.dir,
          },
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return {
      results: data!.models ?? [],
      page: data!.page ?? opts.page ?? 1,
      perPage: data!.perPage ?? opts.perPage ?? 0,
      totalCount: data!.totalCount ?? 0,
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
    const { data, error, response } = await this.api.GET(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            modelIdOrKey: opts.model,
            itemId: opts.itemId,
          },
          query: {
            asset: this.assetParam(opts.asset),
            ref: opts.ref,
          },
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  async getItemsPage(opts: GetItemsPageOptions): Promise<ListResponse<Item>> {
    const { data, error, response } = await this.api.GET(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            modelIdOrKey: opts.model,
          },
          query: {
            page: opts.page,
            perPage: opts.perPage,
            asset: this.assetParam(opts.asset),
            ref: opts.ref,
            keyword: opts.keyword,
            sort: opts.sort,
            dir: opts.dir,
          },
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return {
      results: data!.items ?? [],
      page: data!.page ?? opts.page ?? 1,
      perPage: data!.perPage ?? opts.perPage ?? 0,
      totalCount: data!.totalCount ?? 0,
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
    const { data, error, response } = await this.api.POST(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            modelIdOrKey: opts.model,
          },
        },
        body: {
          fields: opts.fields,
          metadataFields: opts.metadataFields,
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  async updateItem(opts: UpdateItemOptions): Promise<Item> {
    const { data, error, response } = await this.api.PATCH(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            modelIdOrKey: opts.model,
            itemId: opts.itemId,
          },
        },
        body: {
          fields: opts.fields,
          metadataFields: opts.metadataFields,
          asset: opts.asset,
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  async deleteItem(opts: DeleteItemOptions): Promise<void> {
    const { error, response } = await this.api.DELETE(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            modelIdOrKey: opts.model,
            itemId: opts.itemId,
          },
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
  }

  // ---------- Assets ----------

  async getAsset(opts: GetAssetOptions): Promise<Asset> {
    const { data, error, response } = await this.api.GET(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{assetId}",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            assetId: opts.assetId,
          },
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  /**
   * Upload an asset from a public URL. The server fetches and stores it.
   */
  async uploadAsset(opts: UploadAssetByURLOptions): Promise<Asset> {
    const { data, error, response } = await this.api.POST(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
          },
        },
        body: {
          url: opts.url,
          contentEncoding: opts.contentEncoding,
          skipDecompression: opts.skipDecompression ?? false,
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
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

    const { data, error, response } = await this.api.POST(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
          },
        },
        // @ts-expect-error: body is FormData instead of the json shape
        body: form,
        bodySerializer: (body: unknown) => body as FormData,
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  /**
   * Create a 2-stage asset upload. The returned `url` accepts a PUT request
   * with the raw body. After uploading, call `createAssetByToken` with the
   * returned `token` to register the asset.
   */
  async createAssetUpload(opts: CreateAssetUploadOptions): Promise<AssetUpload> {
    const { data, error, response } = await this.api.POST(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/uploads",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
          },
        },
        body: {
          name: opts.name,
          contentType: opts.contentType,
          contentEncoding: opts.contentEncoding,
          contentLength: opts.contentLength,
          cursor: opts.cursor,
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
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
    const { data, error, response } = await this.api.POST(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
          },
        },
        body: {
          token: opts.token,
          contentEncoding: opts.contentEncoding,
          skipDecompression: opts.skipDecompression ?? false,
        },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  // ---------- Comments ----------

  async commentToItem(opts: CommentOptions): Promise<Comment> {
    const { data, error, response } = await this.api.POST(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}/comments",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            modelIdOrKey: opts.model,
            itemId: opts.itemId,
          },
        },
        body: { content: opts.content },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }

  async commentToAsset(opts: CommentToAssetOptions): Promise<Comment> {
    const { data, error, response } = await this.api.POST(
      "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{assetId}/comments",
      {
        params: {
          path: {
            workspaceIdOrAlias: this.options.workspace,
            projectIdOrAlias: this.requireProject(opts.project),
            assetId: opts.assetId,
          },
        },
        body: { content: opts.content },
        signal: this.signal(),
      },
    );
    this.throwIfError(response, error);
    return data!;
  }
}

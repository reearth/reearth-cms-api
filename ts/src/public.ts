import type { FetchLike, Asset } from "./client.js";
import { CMSError, NotFoundError } from "./errors.js";
import { paginateSequential, paginateParallel } from "./paginate.js";

export interface PublicAPIOptions {
  /** Base URL of the Re:Earth CMS server. `/api/p` is appended per request. */
  baseURL: string;
  /** Request timeout in milliseconds. */
  timeout?: number;
  /** Custom fetch implementation (defaults to global `fetch`). */
  fetch?: FetchLike;
  /** Optional access token for published projects that require auth. */
  token?: string;
}

export interface PublicListResponse<T> {
  results: T[];
  page: number;
  perPage: number;
  totalCount: number;
}

export interface GetPublicItemsOptions {
  project: string;
  model: string;
}

export interface GetPublicItemsPageOptions extends GetPublicItemsOptions {
  page?: number;
  perPage?: number;
}

export interface GetAllPublicItemsOptions extends GetPublicItemsOptions {
  perPage?: number;
  /** When set, fetch pages in parallel with the given concurrency. */
  concurrency?: number;
}

export interface PublicAsset extends Asset {
  type?: string;
  files?: string[];
}

export class PublicAPIClient<T = unknown> {
  readonly options: Readonly<PublicAPIOptions>;

  constructor(options: PublicAPIOptions) {
    if (!options.baseURL) throw new Error("baseURL is required");
    this.options = { ...options };
  }

  withTimeout(timeout: number): PublicAPIClient<T> {
    return new PublicAPIClient<T>({ ...this.options, timeout });
  }

  /** Reinterpret this client's item type parameter. */
  as<K>(): PublicAPIClient<K> {
    return new PublicAPIClient<K>(this.options);
  }

  private signal(): AbortSignal | undefined {
    if (!this.options.timeout || this.options.timeout <= 0) return undefined;
    return AbortSignal.timeout(this.options.timeout);
  }

  private url(path: string, query?: Record<string, string | number | undefined>): string {
    const base = this.options.baseURL.replace(/\/+$/, "");
    const u = new URL(base + path);
    if (query) {
      for (const [k, v] of Object.entries(query)) {
        if (v !== undefined) u.searchParams.set(k, String(v));
      }
    }
    return u.toString();
  }

  private async send<R>(url: string): Promise<R> {
    const headers: Record<string, string> = {};
    if (this.options.token) headers["Authorization"] = `Bearer ${this.options.token}`;

    const fetchImpl = this.options.fetch ?? fetch;
    const res = await fetchImpl(url, {
      method: "GET",
      headers,
      signal: this.signal(),
    });
    if (!res.ok) {
      if (res.status === 404) throw new NotFoundError();
      const text = await res.text().catch(() => undefined);
      throw new CMSError(res.status, `request failed: status=${res.status}`, text);
    }
    return (await res.json()) as R;
  }

  async getItemsPage(opts: GetPublicItemsPageOptions): Promise<PublicListResponse<T>> {
    const url = this.url(
      `/api/p/${encodeURIComponent(opts.project)}/${encodeURIComponent(opts.model)}`,
      {
        page: opts.page,
        per_page: opts.perPage,
      },
    );
    const r = await this.send<{
      results?: T[];
      page?: number;
      perPage?: number;
      totalCount?: number;
    }>(url);
    return {
      results: r.results ?? [],
      page: r.page ?? opts.page ?? 1,
      perPage: r.perPage ?? opts.perPage ?? 0,
      totalCount: r.totalCount ?? 0,
    };
  }

  async getAllItems(opts: GetAllPublicItemsOptions): Promise<PublicListResponse<T>> {
    const perPage = opts.perPage ?? 100;
    const fetchPage = (page: number, per: number) =>
      this.getItemsPage({ project: opts.project, model: opts.model, page, perPage: per });

    if (opts.concurrency && opts.concurrency > 1) {
      return paginateParallel(fetchPage, perPage, opts.concurrency);
    }
    return paginateSequential(fetchPage, perPage);
  }

  async getItem(project: string, model: string, id: string): Promise<T> {
    const url = this.url(
      `/api/p/${encodeURIComponent(project)}/${encodeURIComponent(model)}/${encodeURIComponent(id)}`,
    );
    return this.send<T>(url);
  }

  async getAsset(project: string, id: string): Promise<PublicAsset> {
    const url = this.url(`/api/p/${encodeURIComponent(project)}/assets/${encodeURIComponent(id)}`);
    return this.send<PublicAsset>(url);
  }
}

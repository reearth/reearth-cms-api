import { describe, it, expect, vi } from "vitest";
import { CMS, NotFoundError, CMSError } from "../src/index.js";

interface CapturedRequest {
  url: string;
  method: string;
  headers: Record<string, string>;
  body: unknown;
}

function mockFetch(handler: (req: CapturedRequest) => Response | Promise<Response>) {
  return vi.fn(async (input: RequestInfo | URL, init?: RequestInit) => {
    let url: string;
    let method: string;
    let headers: Record<string, string>;
    let body: unknown;

    if (input instanceof Request) {
      url = input.url;
      method = input.method;
      headers = {};
      input.headers.forEach((v, k) => {
        headers[k] = v;
      });
      const ct = input.headers.get("content-type") ?? "";
      if (ct.includes("application/json")) body = await input.clone().text();
      else if (ct.includes("multipart/form-data")) body = await input.clone().formData();
      else body = input.body;
    } else {
      url = typeof input === "string" ? input : input.toString();
      method = init?.method ?? "GET";
      headers = {};
      const rawHeaders = (init?.headers as Record<string, string>) ?? {};
      for (const [k, v] of Object.entries(rawHeaders)) {
        headers[k.toLowerCase()] = v;
      }
      body = init?.body;
    }
    return handler({ url, method, headers, body });
  }) as unknown as typeof fetch;
}

function jsonResponse(body: unknown, status = 200): Response {
  return new Response(JSON.stringify(body), {
    status,
    headers: { "content-type": "application/json" },
  });
}

describe("CMS constructor", () => {
  it("requires baseURL and token", () => {
    expect(() => new CMS({ baseURL: "", token: "t", workspace: "w" })).toThrow(/baseURL/);
    expect(() => new CMS({ baseURL: "https://x", token: "", workspace: "w" })).toThrow(/token/);
  });

  it("allows workspace to be omitted (legacy API mode)", () => {
    const cms = new CMS({ baseURL: "https://x", token: "t" });
    expect(cms.useNewAPI()).toBe(false);
  });

  it("uses new API when workspace is set", () => {
    const cms = new CMS({ baseURL: "https://x", token: "t", workspace: "ws" });
    expect(cms.useNewAPI()).toBe(true);
  });

  it("withProject returns a new instance with project set", () => {
    const a = new CMS({ baseURL: "https://x", token: "t", workspace: "w" });
    const b = a.withProject("p");
    expect(a.options.project).toBeUndefined();
    expect(b.options.project).toBe("p");
  });
});

describe("CMS.getItem", () => {
  it("sends GET with correct URL, auth header, and query", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(
        /\/api\/ws\/projects\/proj\/models\/model\/items\/item-1\?asset=true$/,
      );
      expect(req.headers["authorization"]).toBe("Bearer tok");
      expect(req.method).toBe("GET");
      return jsonResponse({ id: "item-1", modelId: "model" });
    });

    const cms = new CMS({
      baseURL: "https://cms.example.com",
      token: "tok",
      workspace: "ws",
      project: "proj",
      fetch,
    });

    const item = await cms.getItem({ model: "model", itemId: "item-1", asset: true });
    expect(item.id).toBe("item-1");
  });

  it("throws NotFoundError on 404", async () => {
    const fetch = mockFetch(() => new Response("", { status: 404 }));
    const cms = new CMS({ baseURL: "https://x", token: "t", workspace: "w", project: "p", fetch });
    await expect(cms.getItem({ model: "m", itemId: "i" })).rejects.toBeInstanceOf(NotFoundError);
  });

  it("throws CMSError on other failures", async () => {
    const fetch = mockFetch(() => jsonResponse({ error: "bad" }, 500));
    const cms = new CMS({ baseURL: "https://x", token: "t", workspace: "w", project: "p", fetch });
    await expect(cms.getItem({ model: "m", itemId: "i" })).rejects.toBeInstanceOf(CMSError);
  });

  it("throws when project is missing", async () => {
    const fetch = mockFetch(() => jsonResponse({}));
    const cms = new CMS({ baseURL: "https://x", token: "t", workspace: "w", fetch });
    await expect(cms.getItem({ model: "m", itemId: "i" })).rejects.toThrow(/project/);
  });
});

describe("CMS.createItem", () => {
  it("sends POST with JSON body", async () => {
    const fetch = mockFetch(async (req) => {
      expect(req.url).toContain("/api/ws/projects/proj/models/model/items");
      expect(req.method).toBe("POST");
      expect(req.headers["content-type"]).toContain("application/json");
      const body = JSON.parse(req.body as string);
      expect(body.fields).toEqual([{ key: "name", type: "text", value: "hello" }]);
      return jsonResponse({ id: "item-created" });
    });

    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });
    const item = await cms.createItem({
      model: "model",
      fields: [{ key: "name", type: "text", value: "hello" }],
    });
    expect(item.id).toBe("item-created");
  });
});

describe("CMS.uploadAssetDirectly", () => {
  it("sends multipart/form-data", async () => {
    const fetch = mockFetch(async (req) => {
      expect(req.url).toContain("/api/ws/projects/proj/assets");
      expect(req.method).toBe("POST");
      expect(req.body).toBeInstanceOf(FormData);
      const fd = req.body as FormData;
      expect(fd.get("contentType")).toBe("text/plain");
      expect(fd.get("file")).toBeInstanceOf(Blob);
      return jsonResponse({
        id: "asset-1",
        projectId: "proj",
        url: "https://x/a",
        createdAt: "",
        updatedAt: "",
        public: false,
      });
    });

    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });
    const asset = await cms.uploadAssetDirectly({
      file: new Blob(["hello"]),
      name: "hello.txt",
      contentType: "text/plain",
    });
    expect(asset.id).toBe("asset-1");
  });

  it("auto-compresses with compress: 'gzip'", async () => {
    let uploadedSize = 0;
    let uploadedEncoding: FormDataEntryValue | null = null;
    const fetch = mockFetch(async (req) => {
      const fd = req.body as FormData;
      uploadedEncoding = fd.get("contentEncoding");
      const file = fd.get("file") as File;
      uploadedSize = file.size;
      return jsonResponse({
        id: "asset-1",
        projectId: "proj",
        url: "https://x/a",
        createdAt: "",
        updatedAt: "",
        public: false,
      });
    });

    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });
    // highly compressible input
    const payload = "hello world ".repeat(1000);
    await cms.uploadAssetDirectly({
      file: new Blob([payload]),
      name: "big.txt",
      compress: "gzip",
    });

    expect(uploadedEncoding).toBe("gzip");
    expect(uploadedSize).toBeLessThan(payload.length);
    expect(uploadedSize).toBeGreaterThan(0);
  });

  it("does not override explicit contentEncoding", async () => {
    let uploadedEncoding: FormDataEntryValue | null = null;
    const fetch = mockFetch(async (req) => {
      uploadedEncoding = (req.body as FormData).get("contentEncoding");
      return jsonResponse({
        id: "asset-1",
        projectId: "proj",
        url: "https://x/a",
        createdAt: "",
        updatedAt: "",
        public: false,
      });
    });
    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });

    // user already compressed the data; library should NOT re-compress
    await cms.uploadAssetDirectly({
      file: new Blob(["already-compressed"]),
      name: "x.gz",
      contentEncoding: "gzip",
      compress: "gzip",
    });

    expect(uploadedEncoding).toBe("gzip");
  });
});

describe("CMS.uploadToAssetUpload", () => {
  it("streams a Blob body without buffering", async () => {
    let receivedBody: BodyInit | null = null;
    const fetch = vi.fn(async (url: string, init: RequestInit) => {
      receivedBody = init.body as BodyInit;
      return new Response("", { status: 200 });
    }) as unknown as typeof globalThis.fetch;

    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });

    const blob = new Blob(["hello"]);
    await cms.uploadToAssetUpload(
      { url: "https://upload.example.com/presigned", contentType: "text/plain" },
      blob,
    );

    // fetch receives the Blob as-is (streamed by fetch internals)
    expect(receivedBody).toBe(blob);
  });

  it("streams compression through without buffering to Blob", async () => {
    let receivedBody: BodyInit | null = null;
    let receivedDuplex: string | undefined;
    const fetch = vi.fn(async (url: string, init: RequestInit & { duplex?: string }) => {
      receivedBody = init.body as BodyInit;
      receivedDuplex = init.duplex;
      // Drain the stream so the test completes
      if (init.body instanceof ReadableStream) {
        await new Response(init.body).arrayBuffer();
      }
      return new Response("", { status: 200 });
    }) as unknown as typeof globalThis.fetch;

    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });

    await cms.uploadToAssetUpload(
      {
        url: "https://upload.example.com/presigned",
        contentType: "text/plain",
        contentEncoding: "gzip",
      },
      new Blob(["hello ".repeat(1000)]),
      { compress: "gzip" },
    );

    expect(receivedBody).toBeInstanceOf(ReadableStream);
    expect(receivedDuplex).toBe("half");
  });

  it("sets Content-Encoding header from upload.contentEncoding", async () => {
    let encoding: string | undefined;
    const fetch = vi.fn(async (url: string, init: RequestInit) => {
      const headers = init.headers as Record<string, string>;
      encoding = headers["Content-Encoding"];
      return new Response("", { status: 200 });
    }) as unknown as typeof globalThis.fetch;

    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });

    await cms.uploadToAssetUpload(
      { url: "https://up/x", contentType: "text/plain", contentEncoding: "gzip" },
      new Blob(["compressed-bytes"]),
    );
    expect(encoding).toBe("gzip");
  });
});

describe("CMS legacy API (no workspace)", () => {
  it("getItem uses flat /api/items/{id} path", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/items\/item-9$/);
      expect(req.method).toBe("GET");
      return jsonResponse({ id: "item-9", modelId: "m" });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", fetch });
    const item = await cms.getItem({ itemId: "item-9" });
    expect(item.id).toBe("item-9");
  });

  it("updateItem uses flat /api/items/{id} PATCH", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/items\/item-9$/);
      expect(req.method).toBe("PATCH");
      return jsonResponse({ id: "item-9", modelId: "m" });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", fetch });
    await cms.updateItem({ itemId: "item-9", fields: [] });
  });

  it("deleteItem uses flat /api/items/{id} DELETE", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/items\/item-9$/);
      expect(req.method).toBe("DELETE");
      return new Response(null, { status: 204 });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", fetch });
    await cms.deleteItem({ itemId: "item-9" });
  });

  it("getAsset uses flat /api/assets/{id} path", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/assets\/asset-1$/);
      return jsonResponse({
        id: "asset-1",
        projectId: "p",
        url: "https://x/a",
        createdAt: "",
        updatedAt: "",
        public: false,
      });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", fetch });
    const asset = await cms.getAsset({ assetId: "asset-1" });
    expect(asset.id).toBe("asset-1");
  });

  it("getModel with project uses /api/projects/{p}/models/{id}", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/projects\/proj\/models\/mdl$/);
      return jsonResponse({ id: "mdl", key: "mdl" });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", project: "proj", fetch });
    const model = await cms.getModel({ modelIdOrKey: "mdl" });
    expect(model.id).toBe("mdl");
  });

  it("getModel without project uses flat /api/models/{id}", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/models\/mdl-id$/);
      return jsonResponse({ id: "mdl-id" });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", fetch });
    const model = await cms.getModel({ modelIdOrKey: "mdl-id" });
    expect(model.id).toBe("mdl-id");
  });

  it("getItemsPage with project uses /api/projects/{p}/models/{m}/items", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toContain("/api/projects/proj/models/mdl/items");
      return jsonResponse({ items: [], page: 1, perPage: 10, totalCount: 0 });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", project: "proj", fetch });
    await cms.getItemsPage({ model: "mdl", perPage: 10 });
  });

  it("getItemsPage without project uses flat /api/models/{m}/items", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toContain("/api/models/mdl-id/items");
      return jsonResponse({ items: [], page: 1, perPage: 10, totalCount: 0 });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", fetch });
    await cms.getItemsPage({ model: "mdl-id", perPage: 10 });
  });

  it("uploadAsset (by URL) uses /api/projects/{p}/assets", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/projects\/proj\/assets$/);
      expect(req.method).toBe("POST");
      return jsonResponse({
        id: "a1",
        projectId: "proj",
        url: "https://x/a",
        createdAt: "",
        updatedAt: "",
        public: false,
      });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", project: "proj", fetch });
    await cms.uploadAsset({ url: "https://remote/file.jpg" });
  });

  it("commentToItem uses flat /api/items/{id}/comments", async () => {
    const fetch = mockFetch((req) => {
      expect(req.url).toMatch(/\/api\/items\/i-1\/comments$/);
      return jsonResponse({ id: "c1", content: "hello" });
    });
    const cms = new CMS({ baseURL: "https://x", token: "t", fetch });
    await cms.commentToItem({ itemId: "i-1", content: "hello" });
  });
});

describe("CMS new API requires model", () => {
  it("throws if model missing on getItem in new API mode", async () => {
    const fetch = mockFetch(() => jsonResponse({}));
    const cms = new CMS({ baseURL: "https://x", token: "t", workspace: "ws", project: "p", fetch });
    await expect(cms.getItem({ itemId: "i" })).rejects.toThrow(/model/);
  });
});

describe("CMS.getAllItems", () => {
  it("paginates sequentially across pages", async () => {
    let calls = 0;
    const total = 23;
    const fetch = mockFetch(async (req) => {
      calls++;
      const u = new URL(req.url);
      const page = Number(u.searchParams.get("page") ?? "1");
      const perPage = Number(u.searchParams.get("perPage") ?? "10");
      const start = (page - 1) * perPage;
      const count = Math.max(0, Math.min(perPage, total - start));
      return jsonResponse({
        items: Array.from({ length: count }, (_, i) => ({ id: `item-${start + i}`, modelId: "m" })),
        page,
        perPage,
        totalCount: total,
      });
    });
    const cms = new CMS({
      baseURL: "https://x",
      token: "t",
      workspace: "ws",
      project: "proj",
      fetch,
    });

    const r = await cms.getAllItems({ model: "m", perPage: 10 });
    expect(r.results).toHaveLength(23);
    expect(calls).toBe(3);
  });
});

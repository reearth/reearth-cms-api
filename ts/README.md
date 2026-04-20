# @reearth/cms-api

Re:Earth CMS Integration API client for TypeScript/JavaScript.

- Works on **Cloudflare Workers**, **Node.js 18+**, **Bun**, and **Deno**.
- Types generated from the OpenAPI schema (`schemas/integration.yml`).
- Thin, dependency-light (`openapi-fetch` only).
- Webhook verification uses the Web Crypto API — no Node `crypto` dependency.

## Install

```sh
npm install @reearth/cms-api
```

On Node.js < 18 the global `fetch` is not available. Install `node-fetch` (or any `fetch`-compatible polyfill) and pass it in:

```ts
import fetch from "node-fetch";
import { CMS } from "@reearth/cms-api";

const cms = new CMS({
  baseURL: "https://cms.example.com",
  token: process.env.CMS_TOKEN!,
  workspace: "my-workspace",
  fetch: fetch as unknown as typeof globalThis.fetch,
});
```

On Node.js 18+, Cloudflare Workers, Bun, and Deno, no polyfill is needed.

## Integration API

```ts
import { CMS } from "@reearth/cms-api";

const cms = new CMS({
  baseURL: "https://cms.example.com",
  token: process.env.CMS_TOKEN!,
  workspace: "my-workspace",
  project: "my-project", // optional default
});

// Items
const item = await cms.getItem({ model: "model-key", itemId: "...", asset: true });
const { results } = await cms.getAllItems({ model: "model-key", concurrency: 5 });
await cms.createItem({
  model: "model-key",
  fields: [{ key: "name", type: "text", value: "hello" }],
});
await cms.updateItem({ model: "model-key", itemId: "...", fields: [...] });
await cms.deleteItem({ model: "model-key", itemId: "..." });

// Models
const model = await cms.getModel({ modelIdOrKey: "model-key" });
const { results: models } = await cms.getAllModels();

// Assets
const asset = await cms.uploadAssetDirectly({
  file: new Blob(["hello"]),
  name: "hello.txt",
  contentType: "text/plain",
});

// 2-stage upload (for large files)
const upload = await cms.createAssetUpload({ name: "big.zip" });
await cms.uploadToAssetUpload(upload, largeBlob);
const registered = await cms.createAssetByToken({ token: upload.token! });

// Auto-compress with gzip (direct upload)
await cms.uploadAssetDirectly({
  file: new Blob([veryLargeJson]),
  name: "data.json",
  contentType: "application/json",
  compress: "gzip", // also sets contentEncoding: "gzip"
});

// Auto-compress with gzip (2-stage upload, streams end-to-end)
const up = await cms.createAssetUpload({
  name: "data.json",
  contentType: "application/json",
  contentEncoding: "gzip",
});
await cms.uploadToAssetUpload(up, largeBlobOrReadableStream, { compress: "gzip" });
await cms.createAssetByToken({ token: up.token! });

// Comments
await cms.commentToItem({ model: "m", itemId: "i", content: "hello" });
```

### Streaming and buffering

- **`uploadToAssetUpload`** streams end-to-end. Pass a `Blob`, `File`, `ReadableStream`, or
  `ArrayBuffer` — it is handed to `fetch` without reading into memory. With `compress`,
  the source is piped through a `CompressionStream` and the compressed stream is piped
  straight to the presigned URL (no intermediate `Blob`).
- **`uploadAssetDirectly`** uses `multipart/form-data`, which requires a sized `Blob` for
  each part. A `ReadableStream` input is buffered into a `Blob` before the request, and
  `compress` produces a `Blob` too. For very large files prefer the 2-stage upload.

`compressToStream(data, "gzip")` is exposed for advanced use.

### Timeouts

Pass `timeout` (milliseconds) to the constructor to enable an `AbortSignal.timeout()` on every request.

```ts
const cms = new CMS({ ...opts, timeout: 30_000 });
```

### Raw access

The low-level typed `openapi-fetch` client is exposed via `cms.api` for endpoints not wrapped by the high-level methods:

```ts
const { data } = await cms.api.GET("/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups", {
  params: {
    path: { workspaceIdOrAlias: "ws", projectIdOrAlias: "proj" },
  },
});
```

## Public API

For published projects' read-only endpoints, use `PublicAPIClient`:

```ts
import { PublicAPIClient } from "@reearth/cms-api/public";

interface MyItem {
  id: string;
  name: string;
}

const pub = new PublicAPIClient<MyItem>({ baseURL: "https://cms.example.com" });
const { results } = await pub.getAllItems({ project: "proj", model: "model" });
```

## Webhooks (Cloudflare Workers)

```ts
import { verifyAndParseWebhook } from "@reearth/cms-api/webhook";

export default {
  async fetch(request: Request, env: { WEBHOOK_SECRET: string }) {
    try {
      const payload = await verifyAndParseWebhook(request, { secret: env.WEBHOOK_SECRET });

      if (payload.type === "item.create") {
        console.log("new item:", payload.itemData.item);
      } else if (payload.type === "asset.decompress") {
        console.log("decompressed asset:", payload.assetData.id);
      }

      return new Response("ok");
    } catch (e) {
      return new Response("unauthorized", { status: 401 });
    }
  },
};
```

## Types

All OpenAPI-derived types are re-exported:

```ts
import type { components, paths, Item, Asset, Model } from "@reearth/cms-api";

type Field = components["schemas"]["field"];
```

## Development

```sh
npm install
npm run generate    # regenerate src/generated/schema.ts from ../schemas/integration.yml
npm run typecheck
npm test
npm run build
```

When the OpenAPI schema changes (via `make update-schema` at the repo root), re-run `npm run generate` and commit the updated `src/generated/schema.ts`. Or run `make gen` at the repo root to update the schema and regenerate Go + TypeScript artifacts in one go.

## License

MIT

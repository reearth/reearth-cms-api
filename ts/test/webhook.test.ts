import { describe, it, expect } from "vitest";
import {
  verifyAndParseWebhook,
  verifySignature,
  SIGNATURE_HEADER,
  WebhookError,
} from "../src/webhook.js";

const SECRET = "super-secret";

async function makeSignature(body: string, timestamp: number, secret = SECRET): Promise<string> {
  const key = await crypto.subtle.importKey(
    "raw",
    new TextEncoder().encode(secret),
    { name: "HMAC", hash: "SHA-256" },
    false,
    ["sign"],
  );
  const message = new TextEncoder().encode(`v1:${timestamp}:${body}`);
  const sig = await crypto.subtle.sign("HMAC", key, message);
  const hex = Array.from(new Uint8Array(sig))
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
  return `v1,t=${timestamp},${hex}`;
}

function makeRequest(body: string, signature: string): Request {
  return new Request("https://example.com/webhook", {
    method: "POST",
    headers: { [SIGNATURE_HEADER]: signature, "content-type": "application/json" },
    body,
  });
}

describe("verifySignature", () => {
  it("returns true for a valid signature", async () => {
    const body = '{"type":"item.create","data":{}}';
    const ts = Math.floor(Date.now() / 1000);
    const sig = await makeSignature(body, ts);
    const ok = await verifySignature(sig, new TextEncoder().encode(body), { secret: SECRET });
    expect(ok).toBe(true);
  });

  it("returns false for wrong secret", async () => {
    const body = '{"type":"item.create"}';
    const ts = Math.floor(Date.now() / 1000);
    const sig = await makeSignature(body, ts, "wrong-secret");
    const ok = await verifySignature(sig, new TextEncoder().encode(body), { secret: SECRET });
    expect(ok).toBe(false);
  });

  it("returns false for tampered body", async () => {
    const body = '{"type":"item.create"}';
    const ts = Math.floor(Date.now() / 1000);
    const sig = await makeSignature(body, ts);
    const tampered = '{"type":"item.delete"}';
    const ok = await verifySignature(sig, new TextEncoder().encode(tampered), { secret: SECRET });
    expect(ok).toBe(false);
  });

  it("returns false for expired timestamp", async () => {
    const body = "{}";
    const ts = Math.floor(Date.now() / 1000) - 3601;
    const sig = await makeSignature(body, ts);
    const ok = await verifySignature(sig, new TextEncoder().encode(body), { secret: SECRET });
    expect(ok).toBe(false);
  });

  it("returns false for malformed signature", async () => {
    const ok = await verifySignature("not-a-signature", new Uint8Array(), { secret: SECRET });
    expect(ok).toBe(false);
  });

  it("returns false for wrong version", async () => {
    const body = "{}";
    const ts = Math.floor(Date.now() / 1000);
    const sig = (await makeSignature(body, ts)).replace(/^v1/, "v2");
    const ok = await verifySignature(sig, new TextEncoder().encode(body), { secret: SECRET });
    expect(ok).toBe(false);
  });
});

describe("verifyAndParseWebhook", () => {
  it("parses an item.create payload", async () => {
    const body = JSON.stringify({
      type: "item.create",
      data: {
        item: { id: "item-1", modelId: "model-1" },
        schema: { projectId: "project-1" },
      },
      operator: { user: { id: "user-1" } },
    });
    const ts = Math.floor(Date.now() / 1000);
    const sig = await makeSignature(body, ts);
    const req = makeRequest(body, sig);

    const payload = await verifyAndParseWebhook(req, { secret: SECRET });

    expect(payload.type).toBe("item.create");
    if (payload.type === "item.create") {
      expect(payload.itemData.item?.id).toBe("item-1");
      expect(payload.operator.user?.id).toBe("user-1");
    }
  });

  it("parses an asset.decompress payload", async () => {
    const body = JSON.stringify({
      type: "asset.decompress",
      data: { id: "asset-1", projectId: "project-1", url: "https://example.com" },
      operator: {},
    });
    const ts = Math.floor(Date.now() / 1000);
    const sig = await makeSignature(body, ts);
    const req = makeRequest(body, sig);

    const payload = await verifyAndParseWebhook(req, { secret: SECRET });
    expect(payload.type).toBe("asset.decompress");
    if (payload.type === "asset.decompress") {
      expect(payload.assetData.id).toBe("asset-1");
    }
  });

  it("throws on missing signature header", async () => {
    const req = new Request("https://example.com/webhook", {
      method: "POST",
      body: "{}",
    });
    await expect(verifyAndParseWebhook(req, { secret: SECRET })).rejects.toThrow(WebhookError);
  });

  it("throws on invalid signature", async () => {
    const req = makeRequest("{}", "v1,t=0,deadbeef");
    await expect(verifyAndParseWebhook(req, { secret: SECRET })).rejects.toThrow(
      /invalid signature/,
    );
  });
});

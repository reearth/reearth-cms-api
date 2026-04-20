import type { components } from "./generated/schema.js";

export const SIGNATURE_HEADER = "Reearth-Signature";

export const ITEM_EVENT_TYPES = ["item.create", "item.update", "item.publish"] as const;
export const ASSET_EVENT_TYPES = ["asset.decompress"] as const;

export type ItemEventType = (typeof ITEM_EVENT_TYPES)[number];
export type AssetEventType = (typeof ASSET_EVENT_TYPES)[number];
export type WebhookEventType = ItemEventType | AssetEventType;

export const EVENT_ITEM_CREATE = "item.create" satisfies ItemEventType;
export const EVENT_ITEM_UPDATE = "item.update" satisfies ItemEventType;
export const EVENT_ITEM_PUBLISH = "item.publish" satisfies ItemEventType;
export const EVENT_ASSET_DECOMPRESS = "asset.decompress" satisfies AssetEventType;

const VERSION = "v1";
const DEFAULT_EXPIRES_SECONDS = 60 * 60;

type Item = components["schemas"]["item"];
type Model = components["schemas"]["model"];
type Schema = components["schemas"]["schema"];
type Asset = components["schemas"]["asset"];
type Field = components["schemas"]["field"];

export interface Operator {
  user?: { id: string };
  integration?: { id: string };
  machine?: Record<string, never>;
}

export interface ItemData {
  item?: Item;
  model?: Model;
  schema?: Schema;
  referencedItems?: Item[];
  changes?: FieldChange[];
}

export interface FieldChange {
  id?: string;
  type: "add" | "update" | "delete";
  previousValue?: unknown;
  currentValue?: unknown;
}

export type AssetData = Asset;

export interface ItemWebhookPayload {
  type: ItemEventType;
  itemData: ItemData;
  assetData?: undefined;
  operator: Operator;
  /** Raw body bytes (for re-verification or logging). */
  body: Uint8Array;
  /** Signature header value. */
  signature: string;
}

export interface AssetWebhookPayload {
  type: AssetEventType;
  itemData?: undefined;
  assetData: AssetData;
  operator: Operator;
  body: Uint8Array;
  signature: string;
}

export type WebhookPayload = ItemWebhookPayload | AssetWebhookPayload;

function isItemEventType(s: string): s is ItemEventType {
  return (ITEM_EVENT_TYPES as readonly string[]).includes(s);
}

function isAssetEventType(s: string): s is AssetEventType {
  return (ASSET_EVENT_TYPES as readonly string[]).includes(s);
}

export interface VerifyOptions {
  /** Webhook secret. Can be a string (utf-8) or raw bytes. */
  secret: string | Uint8Array;
  /** Maximum signature age in seconds. Default: 3600. */
  expiresSeconds?: number;
  /** Override the current time (for testing). */
  now?: () => Date;
}

export class WebhookError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "WebhookError";
  }
}

/**
 * Verify the webhook signature and parse the body of a `Request` (Fetch API).
 * Works on Cloudflare Workers, Node.js 20+, Bun, and Deno.
 *
 * @throws {WebhookError} if the signature is missing, malformed, expired, or invalid.
 */
export async function verifyAndParseWebhook(
  request: Request,
  opts: VerifyOptions,
): Promise<WebhookPayload> {
  const signature = request.headers.get(SIGNATURE_HEADER);
  if (!signature) throw new WebhookError("missing signature header");

  const body = new Uint8Array(await request.arrayBuffer());

  const ok = await verifySignature(signature, body, opts);
  if (!ok) throw new WebhookError("invalid signature");

  return parseBody(body, signature);
}

/**
 * Verify a signature string against a raw body and secret.
 * Does not throw — returns a boolean.
 */
export async function verifySignature(
  signature: string,
  body: Uint8Array,
  opts: VerifyOptions,
): Promise<boolean> {
  const parts = signature.split(",");
  if (parts.length !== 3 || parts[0] !== VERSION) return false;

  const tsPart = parts[1];
  if (!tsPart?.startsWith("t=")) return false;
  const timestamp = Number.parseInt(tsPart.slice(2), 10);
  if (!Number.isFinite(timestamp)) return false;

  const now = (opts.now ?? (() => new Date()))();
  const ageMs = now.getTime() - timestamp * 1000;
  const expiresMs = (opts.expiresSeconds ?? DEFAULT_EXPIRES_SECONDS) * 1000;
  if (ageMs > expiresMs) return false;

  const expected = await sign(body, opts.secret, timestamp);
  return timingSafeEqual(signature, expected);
}

async function sign(
  body: Uint8Array,
  secret: string | Uint8Array,
  timestamp: number,
): Promise<string> {
  const secretBytes = typeof secret === "string" ? new TextEncoder().encode(secret) : secret;
  const key = await crypto.subtle.importKey(
    "raw",
    secretBytes as BufferSource,
    { name: "HMAC", hash: "SHA-256" },
    false,
    ["sign"],
  );

  const prefix = new TextEncoder().encode(`${VERSION}:${timestamp}:`);
  const message = new Uint8Array(prefix.length + body.length);
  message.set(prefix, 0);
  message.set(body, prefix.length);

  const sig = await crypto.subtle.sign("HMAC", key, message as BufferSource);
  const hex = bytesToHex(new Uint8Array(sig));
  return `${VERSION},t=${timestamp},${hex}`;
}

function bytesToHex(bytes: Uint8Array): string {
  let s = "";
  for (let i = 0; i < bytes.length; i++) {
    s += (bytes[i] as number).toString(16).padStart(2, "0");
  }
  return s;
}

function timingSafeEqual(a: string, b: string): boolean {
  if (a.length !== b.length) return false;
  let diff = 0;
  for (let i = 0; i < a.length; i++) {
    diff |= a.charCodeAt(i) ^ b.charCodeAt(i);
  }
  return diff === 0;
}

function parseBody(body: Uint8Array, signature: string): WebhookPayload {
  const text = new TextDecoder().decode(body);
  const json = JSON.parse(text) as {
    type?: string;
    data?: unknown;
    operator?: Operator;
  };

  if (!json.type) throw new WebhookError("invalid payload: missing type");
  const operator = json.operator ?? {};

  if (isItemEventType(json.type)) {
    return {
      type: json.type,
      itemData: (json.data as ItemData) ?? {},
      operator,
      body,
      signature,
    };
  }
  if (isAssetEventType(json.type)) {
    return {
      type: json.type,
      assetData: (json.data as AssetData) ?? ({} as AssetData),
      operator,
      body,
      signature,
    };
  }
  throw new WebhookError(`unknown webhook type: ${json.type}`);
}

/** Return the project ID from a payload (best-effort). */
export function payloadProjectId(p: WebhookPayload): string | undefined {
  if (p.assetData) return p.assetData.projectId;
  return p.itemData?.schema?.projectId;
}

export type { Field };

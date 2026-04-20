export class CMSError extends Error {
  readonly status: number;
  readonly body: string | undefined;

  constructor(status: number, message: string, body?: string) {
    super(message);
    this.name = "CMSError";
    this.status = status;
    this.body = body;
  }
}

export class NotFoundError extends CMSError {
  constructor(body?: string) {
    super(404, "not found", body);
    this.name = "NotFoundError";
  }
}

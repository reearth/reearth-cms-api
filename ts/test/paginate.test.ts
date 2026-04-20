import { describe, it, expect, vi } from "vitest";
import { paginateSequential, paginateParallel } from "../src/paginate.js";

function makeFetcher(totalCount: number, perPage: number, delayMs = 0) {
  return vi.fn(async (page: number, per: number) => {
    if (delayMs) await new Promise((r) => setTimeout(r, delayMs));
    const start = (page - 1) * per;
    const count = Math.max(0, Math.min(per, totalCount - start));
    return {
      results: Array.from({ length: count }, (_, i) => start + i),
      page,
      perPage: per,
      totalCount,
    };
  });
}

describe("paginateSequential", () => {
  it("concatenates multiple pages", async () => {
    const fetcher = makeFetcher(23, 10);
    const r = await paginateSequential(fetcher, 10);
    expect(r.results).toEqual(Array.from({ length: 23 }, (_, i) => i));
    expect(r.totalCount).toBe(23);
    expect(fetcher).toHaveBeenCalledTimes(3);
  });

  it("handles single page", async () => {
    const fetcher = makeFetcher(5, 10);
    const r = await paginateSequential(fetcher, 10);
    expect(r.results).toHaveLength(5);
    expect(fetcher).toHaveBeenCalledTimes(1);
  });

  it("handles empty results", async () => {
    const fetcher = makeFetcher(0, 10);
    const r = await paginateSequential(fetcher, 10);
    expect(r.results).toEqual([]);
    expect(fetcher).toHaveBeenCalledTimes(1);
  });

  it("throws on invalid perPage", async () => {
    const fetcher = vi.fn(async () => ({ results: [], page: 1, perPage: 0, totalCount: 0 }));
    await expect(paginateSequential(fetcher, 10)).rejects.toThrow(/perPage/);
  });
});

describe("paginateParallel", () => {
  it("concatenates in correct order", async () => {
    const fetcher = makeFetcher(57, 10);
    const r = await paginateParallel(fetcher, 10, 3);
    expect(r.results).toEqual(Array.from({ length: 57 }, (_, i) => i));
    expect(r.totalCount).toBe(57);
    expect(fetcher).toHaveBeenCalledTimes(6);
  });

  it("handles single page", async () => {
    const fetcher = makeFetcher(5, 10);
    const r = await paginateParallel(fetcher, 10, 3);
    expect(r.results).toHaveLength(5);
    expect(fetcher).toHaveBeenCalledTimes(1);
  });

  it("respects concurrency limit", async () => {
    let inFlight = 0;
    let maxInFlight = 0;
    const fetcher = vi.fn(async (page: number, per: number) => {
      inFlight++;
      maxInFlight = Math.max(maxInFlight, inFlight);
      await new Promise((r) => setTimeout(r, 5));
      inFlight--;
      return {
        results: [page],
        page,
        perPage: per,
        totalCount: 100,
      };
    });
    await paginateParallel(fetcher, 10, 3);
    expect(maxInFlight).toBeLessThanOrEqual(3);
  });
});

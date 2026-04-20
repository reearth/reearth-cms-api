/**
 * Fetches pages sequentially until a page signals no more.
 * Returns the concatenated results plus the last page's metadata.
 */
export async function paginateSequential<T>(
  fetchPage: (
    page: number,
    perPage: number,
  ) => Promise<{
    results: T[];
    page: number;
    perPage: number;
    totalCount: number;
  }>,
  perPage = 100,
): Promise<{
  results: T[];
  page: number;
  perPage: number;
  totalCount: number;
}> {
  const all: T[] = [];
  let last = { page: 0, perPage, totalCount: 0 };

  for (let p = 1; ; p++) {
    const r = await fetchPage(p, perPage);
    if (!r.perPage || r.perPage <= 0) {
      throw new Error(`invalid response: perPage=${r.perPage}`);
    }
    all.push(...r.results);
    last = { page: r.page, perPage: r.perPage, totalCount: r.totalCount };

    const maxPage = Math.ceil(r.totalCount / r.perPage);
    if (r.page >= maxPage) break;
  }

  return { results: all, ...last };
}

/**
 * Fetches pages in parallel with a concurrency limit.
 * The first page is fetched first to learn the total count,
 * then remaining pages are fetched in parallel.
 */
export async function paginateParallel<T>(
  fetchPage: (
    page: number,
    perPage: number,
  ) => Promise<{
    results: T[];
    page: number;
    perPage: number;
    totalCount: number;
  }>,
  perPage = 100,
  concurrency = 5,
): Promise<{
  results: T[];
  page: number;
  perPage: number;
  totalCount: number;
}> {
  if (concurrency <= 0) concurrency = 5;

  const first = await fetchPage(1, perPage);
  if (!first.perPage || first.perPage <= 0) {
    throw new Error(`invalid response: perPage=${first.perPage}`);
  }

  const maxPage = Math.ceil(first.totalCount / first.perPage);
  if (maxPage <= 1) {
    return {
      results: first.results,
      page: first.page,
      perPage: first.perPage,
      totalCount: first.totalCount,
    };
  }

  const rest: T[][] = new Array(maxPage - 1);
  let next = 2;
  const worker = async (): Promise<void> => {
    while (true) {
      const p = next++;
      if (p > maxPage) return;
      const r = await fetchPage(p, perPage);
      rest[p - 2] = r.results;
    }
  };

  const workers = Array.from({ length: Math.min(concurrency, maxPage - 1) }, worker);
  await Promise.all(workers);

  const results: T[] = [...first.results];
  for (const chunk of rest) {
    if (chunk) results.push(...chunk);
  }

  return {
    results,
    page: first.page,
    perPage: first.perPage,
    totalCount: first.totalCount,
  };
}

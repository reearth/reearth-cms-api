package cms

import (
	"net/url"
	"strconv"
)

func parallel[T any](limit int, fn func(int) (T, int, error)) ([]T, error) {
	r, maxPage, err := fn(0)
	if err != nil {
		return nil, err
	}
	if maxPage == 0 {
		return nil, nil
	}

	res := make([]T, maxPage)
	res[0] = r

	type result struct {
		Page int
		Res  T
		Err  error
	}

	guard := make(chan struct{}, limit)
	results := make(chan result, maxPage)

	for i := 1; i < maxPage; i++ {
		go func(i int) {
			guard <- struct{}{}
			defer func() { <-guard }()

			r, _, err := fn(i)
			results <- result{Page: i, Res: r, Err: err}
		}(i)
	}

	for i := 1; i < maxPage; i++ {
		r := <-results
		if r.Err != nil {
			return nil, r.Err
		}

		res[r.Page] = r.Res
	}

	return res, nil
}

func cloneURL(u *url.URL) *url.URL {
	u2 := *u
	return &u2
}

func paginationQuery(page, perPage int) map[string][]string {
	q := map[string][]string{}
	if page >= 1 {
		q["page"] = []string{strconv.Itoa(page)}
	}
	if perPage >= 1 {
		q["perPage"] = []string{strconv.Itoa(perPage)}
	}
	return q
}

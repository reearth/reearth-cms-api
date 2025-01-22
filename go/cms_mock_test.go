package cms_test

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	cms "github.com/reearth/reearth-cms-api/go"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func mockCMS(t *testing.T, host, token string) func(string) int {
	t.Helper()

	checkHeader := func(next func(req *http.Request) (any, error)) func(req *http.Request) (*http.Response, error) {
		return func(req *http.Request) (*http.Response, error) {
			if t := parseToken(req); t != token {
				return httpmock.NewJsonResponse(http.StatusUnauthorized, "unauthorized")
			}

			if req.Method != "GET" {
				if c := req.Header.Get("Content-Type"); c != "application/json" && !strings.HasPrefix(c, "multipart/form-data") {
					return httpmock.NewJsonResponse(http.StatusUnsupportedMediaType, "unsupported media type")
				}
			}

			res, err := next(req)
			if err != nil {
				return nil, err
			}
			return httpmock.NewJsonResponse(http.StatusOK, res)
		}
	}

	httpmock.RegisterResponder("GET", host+"/api/items/a", checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{
			"id":     "a",
			"fields": []map[string]string{{"id": "f", "type": "text", "value": "t"}},
		}, nil
	}))

	modelResponder := checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{
			"id":           "idid",
			"key":          "mmm",
			"lastModified": time.Date(2023, time.April, 1, 1, 0, 0, 0, time.UTC),
		}, nil
	})

	modelsResponder := checkHeader(func(r *http.Request) (any, error) {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
		if page < 1 {
			page = 1
		}
		if perPage < 1 {
			perPage = 50
		}
		max := page * perPage
		if max > len(testModels) {
			max = len(testModels)
		}

		return map[string]any{
			"models":     testModels[(page-1)*perPage : max],
			"page":       page,
			"perPage":    perPage,
			"totalCount": len(testModels),
		}, nil
	})

	itemsResponder := checkHeader(func(r *http.Request) (any, error) {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
		if page < 1 {
			page = 1
		}
		if perPage < 1 {
			perPage = 50
		}
		max := page * perPage
		if max > len(testItems) {
			max = len(testItems)
		}

		return map[string]any{
			"items":      testItems[(page-1)*perPage : max],
			"page":       page,
			"perPage":    perPage,
			"totalCount": len(testItems),
		}, nil
	})

	emptyItemsResponder := checkHeader(func(r *http.Request) (any, error) {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
		if page < 1 {
			page = 1
		}
		if perPage < 1 {
			perPage = 50
		}

		return map[string]any{
			"items":      nil,
			"page":       page,
			"perPage":    perPage,
			"totalCount": 0,
		}, nil
	})

	httpmock.RegisterResponder("GET", host+"/api/projects/ppp/models/mmm", modelResponder)
	httpmock.RegisterResponder("GET", host+"/api/models/mmm", modelResponder)
	httpmock.RegisterResponder("GET", host+"/api/projects/ppp/models", modelsResponder)
	httpmock.RegisterResponder("GET", host+"/api/projects/ppp/models/mmm/items", itemsResponder)
	httpmock.RegisterResponder("GET", host+"/api/projects/ppp/models/empty/items", emptyItemsResponder)
	httpmock.RegisterResponder("GET", host+"/api/models/mmm/items", itemsResponder)
	httpmock.RegisterResponder("GET", host+"/api/models/empty/items", emptyItemsResponder)

	httpmock.RegisterResponder("POST", host+"/api/projects/ppp/models/mmm/items", checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{
			"id":     "a",
			"fields": []map[string]string{{"id": "f", "type": "text", "value": "t"}},
		}, nil
	}))

	httpmock.RegisterResponder("PATCH", host+"/api/items/a", checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{
			"id":     "a",
			"fields": []map[string]string{{"id": "f", "type": "text", "value": "t"}},
		}, nil
	}))

	httpmock.RegisterResponder("POST", host+"/api/models/a/items", checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{
			"id":     "a",
			"fields": []map[string]string{{"id": "f", "type": "text", "value": "t"}},
		}, nil
	}))

	httpmock.RegisterResponder("DELETE", host+"/api/items/a", checkHeader(func(r *http.Request) (any, error) {
		return nil, nil
	}))

	httpmock.RegisterResponder("POST", host+"/api/projects/ppp/assets", checkHeader(func(r *http.Request) (any, error) {
		if c := r.Header.Get("Content-Type"); strings.HasPrefix(c, "multipart/form-data") {
			f, fh, err := r.FormFile("file")
			if err != nil {
				return nil, err
			}
			defer func() {
				_ = f.Close()
			}()
			d, _ := io.ReadAll(f)
			assert.Equal(t, "datadata", string(d))
			assert.Equal(t, "file.txt", fh.Filename)
		}

		return map[string]any{
			"id": "idid",
		}, nil
	}))

	httpmock.RegisterResponder("POST", host+"/api/items/itit/comments", checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{}, nil
	}))

	httpmock.RegisterResponder("POST", host+"/api/assets/c/comments", checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{}, nil
	}))

	httpmock.RegisterResponder("GET", host+"/api/assets/a", checkHeader(func(r *http.Request) (any, error) {
		return map[string]any{
			"id":  "a",
			"url": "url",
		}, nil
	}))

	return func(p string) int {
		b, a, _ := strings.Cut(p, " ")
		return httpmock.GetCallCountInfo()[b+" "+host+a]
	}
}

func parseToken(r *http.Request) string {
	aut := r.Header.Get("Authorization")
	_, token, found := strings.Cut(aut, "Bearer ")
	if !found {
		return ""
	}
	return token
}

var testItems = lo.Map(lo.Range(500), func(i, _ int) cms.Item {
	return cms.Item{
		ID:     strconv.Itoa(i),
		Fields: []*cms.Field{{ID: "f", Type: "text", Value: "t"}},
	}
})

var testModels = lo.Map(lo.Range(500), func(i, _ int) cms.Model {
	return cms.Model{
		ID: strconv.Itoa(i),
	}
})

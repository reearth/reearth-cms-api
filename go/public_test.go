package reearthcmsapi

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestPublicAPIClient_GetAllItems(t *testing.T) {
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm", "page=1&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results": []any{
			map[string]any{"id": "a"},
		},
		"totalCount": 101,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm", "page=2&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results": []any{
			map[string]any{"id": "b"},
		},
		"totalCount": 101,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm2", "", lo.Must(httpmock.NewJsonResponder(http.StatusNotFound, nil)))

	c, err := NewPublicAPIClient[any](nil, "https://example.com/")
	assert.NoError(t, err)
	res, err := c.GetAllItems(ctx, "ppp", "mmm")
	assert.NoError(t, err)
	assert.Equal(t, []any{
		map[string]any{"id": "a"},
		map[string]any{"id": "b"},
	}, res)

	res, err = c.GetAllItems(ctx, "ppp", "mmm2")
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, res)
}

func TestPublicAPIClient_GetAllItemsInParallel(t *testing.T) {
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm", "page=1&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results": []any{
			map[string]any{"id": "a"},
		},
		"totalCount": 101,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm", "page=2&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results": []any{
			map[string]any{"id": "b"},
		},
		"totalCount": 101,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm2", "", lo.Must(httpmock.NewJsonResponder(http.StatusNotFound, nil)))

	c, err := NewPublicAPIClient[any](nil, "https://example.com/")
	assert.NoError(t, err)
	res, err := c.GetAllItemsInParallel(ctx, "ppp", "mmm", 2)
	assert.NoError(t, err)
	assert.Equal(t, []any{
		map[string]any{"id": "a"},
		map[string]any{"id": "b"},
	}, res)

	res, err = c.GetAllItemsInParallel(ctx, "ppp", "mmm2", 0)
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, res)
}

func TestPublicAPIClient_GetItems(t *testing.T) {
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm", "page=1&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results": []any{
			map[string]any{"id": "a"},
			map[string]any{"id": "b"},
		},
		"totalCount": 2,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/api/p/ppp/mmm2", "", lo.Must(httpmock.NewJsonResponder(http.StatusNotFound, nil)))

	c, err := NewPublicAPIClient[any](nil, "https://example.com/")
	assert.NoError(t, err)
	res, err := c.GetItems(ctx, "ppp", "mmm", 1, 100)
	assert.NoError(t, err)
	assert.Equal(t, &PublicAPIListResponse[any]{
		Results: []any{
			map[string]any{"id": "a"},
			map[string]any{"id": "b"},
		},
		PerPage:    100,
		Page:       1,
		TotalCount: 2,
	}, res)

	res, err = c.GetItems(ctx, "ppp", "mmm2", 0, 0)
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, res)
}

func TestPublicAPIClient_GetItem(t *testing.T) {
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("GET", "https://example.com/api/p/ppp/mmm/iii", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"id": "a",
	})))
	httpmock.RegisterResponder("GET", "https://example.com/api/p/ppp/mmm/iii2", lo.Must(httpmock.NewJsonResponder(http.StatusNotFound, nil)))

	c, err := NewPublicAPIClient[any](nil, "https://example.com/")
	assert.NoError(t, err)
	res, err := c.GetItem(ctx, "ppp", "mmm", "iii")
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"id": "a"}, res)

	res, err = c.GetItem(ctx, "ppp", "mmm", "iii2")
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, res)
}

func TestPublicAPIClient_GetAsset(t *testing.T) {
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("GET", "https://example.com/api/p/ppp/assets/aaa", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"id":    "aaa",
		"url":   "https://example.com",
		"files": []string{"https://example.com/a.txt", "https://example.com/b.txt"},
	})))
	httpmock.RegisterResponder("GET", "https://example.com/api/p/ppp/assets/aaa2", lo.Must(httpmock.NewJsonResponder(http.StatusNotFound, nil)))

	c, err := NewPublicAPIClient[any](nil, "https://example.com/")
	assert.NoError(t, err)
	res, err := c.GetAsset(ctx, "ppp", "aaa")
	assert.NoError(t, err)
	assert.Equal(t, &PublicAsset{
		ID:  "aaa",
		URL: "https://example.com",
		Files: []string{
			"https://example.com/a.txt",
			"https://example.com/b.txt",
		},
	}, res)

	res, err = c.GetAsset(ctx, "ppp", "aaa2")
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, res)
}

func TestPublicAPIListResponse_HasNext(t *testing.T) {
	assert.True(t, PublicAPIListResponse[any]{Page: 1, PerPage: 50, TotalCount: 100}.HasNext())
	assert.False(t, PublicAPIListResponse[any]{Page: 2, PerPage: 50, TotalCount: 100}.HasNext())
	assert.True(t, PublicAPIListResponse[any]{Page: 1, PerPage: 10, TotalCount: 11}.HasNext())
	assert.False(t, PublicAPIListResponse[any]{Page: 2, PerPage: 10, TotalCount: 11}.HasNext())
}

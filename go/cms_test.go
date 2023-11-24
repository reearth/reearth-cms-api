package cms

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var _ Interface = (*CMS)(nil)

func TestCMS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	ctx := context.Background()

	// valid
	call := mockCMS(t, "http://cms.example.com", "TOKEN")
	c := lo.Must(New("http://cms.example.com", "TOKEN"))

	model, err := c.GetModel(ctx, "mmm")
	assert.Equal(t, 1, call("GET /api/models/mmm"))
	assert.NoError(t, err)
	assert.Equal(t, &Model{
		ID:           "idid",
		Key:          "mmm",
		LastModified: time.Date(2023, time.April, 1, 1, 0, 0, 0, time.UTC),
	}, model)

	model, err = c.GetModelByKey(ctx, "ppp", "mmm")
	assert.Equal(t, 1, call("GET /api/projects/ppp/models/mmm"))
	assert.NoError(t, err)
	assert.Equal(t, &Model{
		ID:           "idid",
		Key:          "mmm",
		LastModified: time.Date(2023, time.April, 1, 1, 0, 0, 0, time.UTC),
	}, model)

	item, err := c.GetItem(ctx, "a", false)
	assert.Equal(t, 1, call("GET /api/items/a"))
	assert.NoError(t, err)
	assert.Equal(t, &Item{
		ID:     "a",
		Fields: []*Field{{ID: "f", Type: "text", Value: "t"}},
	}, item)

	item, err = c.CreateItem(ctx, "a", nil, nil)
	assert.Equal(t, 1, call("POST /api/models/a/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Item{
		ID:     "a",
		Fields: []*Field{{ID: "f", Type: "text", Value: "t"}},
	}, item)

	item, err = c.CreateItemByKey(ctx, "ppp", "mmm", nil, nil)
	assert.Equal(t, 1, call("POST /api/projects/ppp/models/mmm/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Item{
		ID:     "a",
		Fields: []*Field{{ID: "f", Type: "text", Value: "t"}},
	}, item)

	item, err = c.UpdateItem(ctx, "a", nil, nil)
	assert.Equal(t, 1, call("PATCH /api/items/a"))
	assert.NoError(t, err)
	assert.Equal(t, &Item{
		ID:     "a",
		Fields: []*Field{{ID: "f", Type: "text", Value: "t"}},
	}, item)

	err = c.DeleteItem(ctx, "a")
	assert.Equal(t, 1, call("DELETE /api/items/a"))
	assert.NoError(t, err)

	a, err := c.Asset(ctx, "a")
	assert.Equal(t, 1, call("GET /api/assets/a"))
	assert.NoError(t, err)
	assert.Equal(t, &Asset{ID: "a", URL: "url"}, a)

	assetID, err := c.UploadAsset(ctx, "ppp", "aaa")
	assert.Equal(t, 1, call("POST /api/projects/ppp/assets"))
	assert.NoError(t, err)
	assert.Equal(t, "idid", assetID)

	assetID, err = c.UploadAssetDirectly(ctx, "ppp", "file.txt", strings.NewReader("datadata"))
	assert.Equal(t, 2, call("POST /api/projects/ppp/assets"))
	assert.NoError(t, err)
	assert.Equal(t, "idid", assetID)

	assert.NoError(t, c.CommentToAsset(ctx, "c", "comment"))
	assert.Equal(t, 1, call("POST /api/assets/c/comments"))

	// invalid token
	httpmock.Reset()
	call = mockCMS(t, "http://cms.example.com", "TOKEN")
	c = lo.Must(New("http://cms.example.com", "TOKEN2"))

	model, err = c.GetModel(ctx, "mmm")
	assert.Equal(t, 1, call("GET /api/models/mmm"))
	assert.Nil(t, model)
	assert.ErrorContains(t, err, "failed to request: code=401")

	model, err = c.GetModelByKey(ctx, "ppp", "mmm")
	assert.Equal(t, 1, call("GET /api/projects/ppp/models/mmm"))
	assert.Nil(t, model)
	assert.ErrorContains(t, err, "failed to request: code=401")

	item, err = c.GetItem(ctx, "a", false)
	assert.Equal(t, 1, call("GET /api/items/a"))
	assert.Nil(t, item)
	assert.ErrorContains(t, err, "failed to request: code=401")

	items, err := c.GetItemsPartially(ctx, "mmm", 1, 1, false)
	assert.Equal(t, 1, call("GET /api/models/mmm/items"))
	assert.Nil(t, items)
	assert.ErrorContains(t, err, "failed to request: code=401")

	items, err = c.GetItemsPartiallyByKey(ctx, "ppp", "mmm", 1, 1, false)
	assert.Equal(t, 1, call("GET /api/projects/ppp/models/mmm/items"))
	assert.Nil(t, items)
	assert.ErrorContains(t, err, "failed to request: code=401")

	items, err = c.GetItems(ctx, "mmm", false)
	assert.Equal(t, 2, call("GET /api/models/mmm/items"))
	assert.Nil(t, items)
	assert.ErrorContains(t, err, "failed to request: code=401")

	items, err = c.GetItemsByKey(ctx, "ppp", "mmm", false)
	assert.Equal(t, 2, call("GET /api/projects/ppp/models/mmm/items"))
	assert.Nil(t, items)
	assert.ErrorContains(t, err, "failed to request: code=401")

	item, err = c.CreateItemByKey(ctx, "ppp", "mmm", nil, nil)
	assert.Equal(t, 1, call("POST /api/projects/ppp/models/mmm/items"))
	assert.Nil(t, item)
	assert.ErrorContains(t, err, "failed to request: code=401")

	item, err = c.CreateItem(ctx, "a", nil, nil)
	assert.Equal(t, 1, call("POST /api/models/a/items"))
	assert.Nil(t, item)
	assert.ErrorContains(t, err, "failed to request: code=401")

	item, err = c.UpdateItem(ctx, "a", nil, nil)
	assert.Equal(t, 1, call("PATCH /api/items/a"))
	assert.Nil(t, item)
	assert.ErrorContains(t, err, "failed to request: code=401")

	err = c.DeleteItem(ctx, "a")
	assert.Equal(t, 1, call("DELETE /api/items/a"))
	assert.Nil(t, item)
	assert.ErrorContains(t, err, "failed to request: code=401")

	assetID, err = c.UploadAsset(ctx, "ppp", "aaa")
	assert.Equal(t, 1, call("POST /api/projects/ppp/assets"))
	assert.ErrorContains(t, err, "failed to request: code=401")
	assert.Equal(t, "", assetID)

	assetID, err = c.UploadAssetDirectly(ctx, "ppp", "file.txt", strings.NewReader("datadata"))
	assert.Equal(t, 2, call("POST /api/projects/ppp/assets"))
	assert.ErrorContains(t, err, "failed to request: code=401")
	assert.Equal(t, "", assetID)

	assert.ErrorContains(t, c.CommentToAsset(ctx, "c", "comment"), "failed to request: code=401")
	assert.Equal(t, 1, call("POST /api/assets/c/comments"))

	_, err = c.Asset(ctx, "a")
	assert.Equal(t, 1, call("GET /api/assets/a"))
	assert.ErrorContains(t, err, "failed to request: code=401")
}

func TestCMS_GetModels(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	ctx := context.Background()
	call := mockCMS(t, "http://cms.example.com", "TOKEN")
	c := lo.Must(New("http://cms.example.com", "TOKEN"))

	items, err := c.GetModelsPartially(ctx, "ppp", 1, 1)
	assert.Equal(t, 1, call("GET /api/projects/ppp/models"))
	assert.NoError(t, err)
	assert.Equal(t, &Models{
		Models:     testModels[0:1],
		Page:       1,
		PerPage:    1,
		TotalCount: len(testModels),
	}, items)

	items, err = c.GetModels(ctx, "ppp")
	assert.Equal(t, 6, call("GET /api/projects/ppp/models"))
	assert.NoError(t, err)
	assert.Equal(t, &Models{
		Models:     testModels,
		Page:       1,
		PerPage:    100,
		TotalCount: len(testModels),
	}, items)
}

func TestCMS_GetItems(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	ctx := context.Background()
	call := mockCMS(t, "http://cms.example.com", "TOKEN")
	c := lo.Must(New("http://cms.example.com", "TOKEN"))

	items, err := c.GetItemsPartially(ctx, "mmm", 1, 1, false)
	assert.Equal(t, 1, call("GET /api/models/mmm/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Items{
		Items:      testItems[0:1],
		Page:       1,
		PerPage:    1,
		TotalCount: len(testItems),
	}, items)

	items, err = c.GetItemsPartiallyByKey(ctx, "ppp", "mmm", 1, 1, false)
	assert.Equal(t, 1, call("GET /api/projects/ppp/models/mmm/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Items{
		Items:      testItems[0:1],
		Page:       1,
		PerPage:    1,
		TotalCount: len(testItems),
	}, items)

	items, err = c.GetItems(ctx, "mmm", false)
	assert.Equal(t, 6, call("GET /api/models/mmm/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Items{
		Items:      testItems,
		Page:       1,
		PerPage:    100,
		TotalCount: len(testItems),
	}, items)

	items, err = c.GetItemsByKey(ctx, "ppp", "mmm", false)
	assert.Equal(t, 6, call("GET /api/projects/ppp/models/mmm/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Items{
		Items:      testItems,
		Page:       1,
		PerPage:    100,
		TotalCount: len(testItems),
	}, items)
}

func TestCMS_GetItemsInParallel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	ctx := context.Background()
	call := mockCMS(t, "http://cms.example.com", "TOKEN")
	c := lo.Must(New("http://cms.example.com", "TOKEN"))

	items, err := c.GetItemsInParallel(ctx, "mmm", false, 5)
	assert.Equal(t, 5, call("GET /api/models/mmm/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Items{
		Items:      testItems,
		Page:       1,
		PerPage:    100,
		TotalCount: len(testItems),
	}, items)
}

func TestCMS_GetItemsByKeyInParallel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	ctx := context.Background()
	call := mockCMS(t, "http://cms.example.com", "TOKEN")
	c := lo.Must(New("http://cms.example.com", "TOKEN"))

	items, err := c.GetItemsByKeyInParallel(ctx, "ppp", "mmm", false, 5)
	assert.Equal(t, 5, call("GET /api/projects/ppp/models/mmm/items"))
	assert.NoError(t, err)
	assert.Equal(t, &Items{
		Items:      testItems,
		Page:       1,
		PerPage:    100,
		TotalCount: len(testItems),
	}, items)

	// empty
	items, err = c.GetItemsByKeyInParallel(ctx, "ppp", "empty", false, 5)
	assert.Equal(t, 1, call("GET /api/projects/ppp/models/empty/items"))
	assert.NoError(t, err)
	assert.Nil(t, items)
}

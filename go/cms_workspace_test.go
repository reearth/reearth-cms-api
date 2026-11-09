package cms_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	cms "github.com/reearth/reearth-cms-api/go"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCMS_ItemWithWorkspace(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	ctx := context.Background()
	const host = "http://cms.example.com"
	const itemPath = host + "/api/ws/projects/ppp/models/mmm/items/a"

	item := map[string]any{"id": "a", "modelId": "mmm"}
	httpmock.RegisterResponder("GET", itemPath, httpmock.NewJsonResponderOrPanic(http.StatusOK, item))
	httpmock.RegisterResponder("PATCH", itemPath, httpmock.NewJsonResponderOrPanic(http.StatusOK, item))
	httpmock.RegisterResponder("DELETE", itemPath, httpmock.NewJsonResponderOrPanic(http.StatusOK, item))
	httpmock.RegisterResponder("POST", itemPath+"/comments", httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]any{}))

	c := lo.Must(cms.New(host, "TOKEN")).WithWorkspace("ws").WithProject("ppp")

	t.Run("ByModel methods use the model-scoped path", func(t *testing.T) {
		httpmock.ZeroCallCounters()

		got, err := c.GetItemByModel(ctx, "mmm", "a", false)
		assert.NoError(t, err)
		assert.Equal(t, "a", got.ID)

		got, err = c.UpdateItemByModel(ctx, "mmm", "a", nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, "a", got.ID)

		assert.NoError(t, c.DeleteItemByModel(ctx, "mmm", "a"))
		assert.NoError(t, c.CommentToItemByModel(ctx, "mmm", "a", "comment"))

		info := httpmock.GetCallCountInfo()
		assert.Equal(t, 1, info["GET "+itemPath])
		assert.Equal(t, 1, info["PATCH "+itemPath])
		assert.Equal(t, 1, info["DELETE "+itemPath])
		assert.Equal(t, 1, info["POST "+itemPath+"/comments"])
	})

	t.Run("methods without model return ErrModelRequired without sending a request", func(t *testing.T) {
		httpmock.ZeroCallCounters()

		_, err := c.GetItem(ctx, "a", false)
		assert.ErrorIs(t, err, cms.ErrModelRequired)

		_, err = c.UpdateItem(ctx, "a", nil, nil)
		assert.ErrorIs(t, err, cms.ErrModelRequired)

		assert.ErrorIs(t, c.DeleteItem(ctx, "a"), cms.ErrModelRequired)
		assert.ErrorIs(t, c.CommentToItem(ctx, "a", "comment"), cms.ErrModelRequired)

		_, err = c.GetItemByModel(ctx, "", "a", false)
		assert.ErrorIs(t, err, cms.ErrModelRequired)

		assert.Equal(t, 0, httpmock.GetTotalCallCount())
	})
}

func TestCMS_ItemByModelWithoutWorkspace(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	ctx := context.Background()
	const host = "http://cms.example.com"

	httpmock.RegisterResponder("GET", host+"/api/items/a", httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]any{"id": "a"}))
	httpmock.RegisterResponder("POST", host+"/api/items/a/comments", httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]any{}))

	c := lo.Must(cms.New(host, "TOKEN"))

	// the legacy API does not need the model; it is ignored
	got, err := c.GetItemByModel(ctx, "mmm", "a", false)
	assert.NoError(t, err)
	assert.Equal(t, "a", got.ID)
	assert.NoError(t, c.CommentToItemByModel(ctx, "mmm", "a", "comment"))

	info := httpmock.GetCallCountInfo()
	assert.Equal(t, 1, info["GET "+host+"/api/items/a"])
	assert.Equal(t, 1, info["POST "+host+"/api/items/a/comments"])
}

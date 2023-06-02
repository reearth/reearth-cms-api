package cmswebhook

import (
	"encoding/json"
	"testing"

	reearthcmsapi "github.com/reearth/reearth-cms-api"
	"github.com/stretchr/testify/assert"
)

func TestPayload_UnmarshalJSON(t *testing.T) {
	p := Payload{}
	assert.NoError(t, json.Unmarshal([]byte(`{"type":"item.update","data":{"item":{"id":"i"}}}`), &p))
	assert.Equal(t, Payload{
		Type: "item.update",
		ItemData: &ItemData{
			Item: &reearthcmsapi.Item{
				ID: "i",
			},
		},
	}, p)

	p = Payload{}
	assert.NoError(t, json.Unmarshal([]byte(`{"type":"asset.decompress","data":{"id":"a"}}`), &p))
	assert.Equal(t, Payload{
		Type: "asset.decompress",
		AssetData: &AssetData{
			ID: "a",
		},
	}, p)
}

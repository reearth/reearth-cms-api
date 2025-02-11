package cms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
)

var ErrNotFound = errors.New("not found")

type Interface interface {
	GetModel(ctx context.Context, modelID string) (*Model, error)
	GetModelByKey(ctx context.Context, proejctID, modelID string) (*Model, error)
	GetModelsPartially(ctx context.Context, projectIDOrAlias string, page, perPage int) (*Models, error)
	GetModelsInParallel(ctx context.Context, modelID string, limit int) (*Models, error)
	GetModels(ctx context.Context, projectIDOrAlias string) (*Models, error)
	GetItem(ctx context.Context, itemID string, asset bool) (*Item, error)
	GetItemsPartially(ctx context.Context, modelID string, page, perPage int, asset bool) (*Items, error)
	GetItems(ctx context.Context, modelID string, asset bool) (*Items, error)
	GetItemsInParallel(ctx context.Context, modelID string, asset bool, limit int) (*Items, error)
	GetItemsPartiallyByKey(ctx context.Context, projectIDOrAlias, modelIDOrKey string, page, perPage int, asset bool) (*Items, error)
	GetItemsByKey(ctx context.Context, projectIDOrAlias, modelIDOrKey string, asset bool) (*Items, error)
	GetItemsByKeyInParallel(ctx context.Context, projectIDOrAlias, modelIDOrKey string, asset bool, limit int) (*Items, error)
	CreateItem(ctx context.Context, modelID string, fields []*Field, metadataFields []*Field) (*Item, error)
	CreateItemByKey(ctx context.Context, projectID, modelID string, fields []*Field, metadataFields []*Field) (*Item, error)
	UpdateItem(ctx context.Context, itemID string, fields []*Field, metadataFields []*Field) (*Item, error)
	DeleteItem(ctx context.Context, itemID string) error
	Asset(ctx context.Context, id string) (*Asset, error)
	UploadAsset(ctx context.Context, projectID, url string) (string, error)
	UploadAssetDirectly(ctx context.Context, projectID, name string, data io.Reader, opts ...UploadAssetOption) (string, error)
	CreateAssetUpload(ctx context.Context, projectID, name string, opts ...UploadAssetOption) (*AssetUpload, error)
	UploadToAssetUpload(ctx context.Context, upload *AssetUpload, data io.Reader) error
	CreateAssetByToken(ctx context.Context, projectID, token string) (*Asset, error)
	CommentToItem(ctx context.Context, assetID, content string) error
	CommentToAsset(ctx context.Context, assetID, content string) error
}

type CMS struct {
	base    *url.URL
	token   string
	client  *http.Client
	timeout time.Duration
}

func New(base, token string) (*CMS, error) {
	b, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %w", err)
	}

	return &CMS{
		base:   b,
		token:  token,
		client: http.DefaultClient,
	}, nil
}

func (c *CMS) WithHTTPClient(hc *http.Client) *CMS {
	return &CMS{
		base:    c.base,
		token:   c.token,
		client:  hc,
		timeout: c.timeout,
	}
}

func (c *CMS) WithTimeout(t time.Duration) *CMS {
	return &CMS{
		base:    c.base,
		token:   c.token,
		client:  c.client,
		timeout: t,
	}
}

func (c *CMS) assetParam(asset bool) map[string][]string {
	if !asset {
		return make(map[string][]string)
	}
	return map[string][]string{
		"asset": {"true"},
	}
}

func (c *CMS) GetModel(ctx context.Context, modelID string) (*Model, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "models", modelID}, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get an model: %w", err)
	}
	defer func() { _ = b.Close() }()

	model := &Model{}
	if err := json.NewDecoder(b).Decode(model); err != nil {
		return nil, fmt.Errorf("failed to parse an model: %w", err)
	}

	return model, nil
}

func (c *CMS) GetModelByKey(ctx context.Context, projectKey, modelKey string) (*Model, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "projects", projectKey, "models", modelKey}, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get an model: %w", err)
	}
	defer func() { _ = b.Close() }()

	model := &Model{}
	if err := json.NewDecoder(b).Decode(model); err != nil {
		return nil, fmt.Errorf("failed to parse an model: %w", err)
	}

	return model, nil
}

func (c *CMS) GetModelsPartially(ctx context.Context, projectIDOrAlias string, page, perPage int) (*Models, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "projects", projectIDOrAlias, "models"}, "", paginationQuery(page, perPage))
	if err != nil {
		return nil, fmt.Errorf("failed to get models: %w", err)
	}
	defer func() { _ = b.Close() }()

	models := &Models{}
	if err := json.NewDecoder(b).Decode(models); err != nil {
		return nil, fmt.Errorf("failed to parse models: %w", err)
	}

	return models, nil
}

func (c *CMS) GetModels(ctx context.Context, projectIDOrAlias string) (models *Models, _ error) {
	const perPage = 100
	for p := 1; ; p++ {
		i, err := c.GetModelsPartially(ctx, projectIDOrAlias, p, perPage)
		if err != nil {
			return nil, err
		}

		if i == nil || i.PerPage <= 0 {
			return nil, fmt.Errorf("invalid response: %#v", i)
		}

		if models == nil {
			models = i
		} else {
			models.Models = append(models.Models, i.Models...)
		}

		allPageCount := i.TotalCount / i.PerPage
		if i.Page >= allPageCount {
			break
		}
	}

	return models, nil
}

func (c *CMS) GetModelsInParallel(ctx context.Context, modelID string, limit int) (*Models, error) {
	const perPage = 100
	if limit <= 0 {
		limit = 5
	}

	res, err := parallel(limit, func(p int) (*Models, int, error) {
		r, err := c.GetModelsPartially(ctx, modelID, p+1, perPage)
		if err != nil || r == nil {
			if r == nil || r.PerPage == 0 {
				err = fmt.Errorf("invalid response: %#v", r)
			}
			return nil, 0, err
		}
		return r, int(math.Ceil(float64(r.TotalCount) / float64(perPage))), nil
	})
	if err != nil {
		return nil, err
	}

	res2 := res[0]
	res2.Models = lo.FlatMap(res, func(i *Models, _ int) []Model {
		if i == nil {
			return nil
		}
		return i.Models
	})
	return res2, nil
}

func (c *CMS) GetItem(ctx context.Context, itemID string, asset bool) (*Item, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "items", itemID}, "", c.assetParam(asset))
	if err != nil {
		return nil, fmt.Errorf("failed to get an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) GetItemsPartially(ctx context.Context, modelID string, page, perPage int, asset bool) (*Items, error) {
	q := c.assetParam(asset)
	if page >= 1 {
		q["page"] = []string{strconv.Itoa(page)}
	}
	if perPage >= 1 {
		q["perPage"] = []string{strconv.Itoa(perPage)}
	}

	b, err := c.send(ctx, http.MethodGet, []string{"api", "models", modelID, "items"}, "", q)
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}
	defer func() { _ = b.Close() }()

	items := &Items{}
	if err := json.NewDecoder(b).Decode(items); err != nil {
		return nil, fmt.Errorf("failed to parse items: %w", err)
	}

	return items, nil
}

func (c *CMS) GetItems(ctx context.Context, modelID string, asset bool) (items *Items, _ error) {
	const perPage = 100
	for p := 1; ; p++ {
		i, err := c.GetItemsPartially(ctx, modelID, p, perPage, asset)
		if err != nil {
			return nil, err
		}

		if i == nil || i.PerPage <= 0 {
			return nil, fmt.Errorf("invalid response: %#v", i)
		}

		if items == nil {
			items = i
		} else {
			items.Items = append(items.Items, i.Items...)
		}

		allPageCount := i.TotalCount / i.PerPage
		if i.Page >= allPageCount {
			break
		}
	}

	return items, nil
}

func (c *CMS) GetItemsInParallel(ctx context.Context, modelID string, asset bool, limit int) (*Items, error) {
	const perPage = 100
	if limit <= 0 {
		limit = 5
	}

	res, err := parallel(limit, func(p int) (*Items, int, error) {
		r, err := c.GetItemsPartially(ctx, modelID, p+1, perPage, asset)
		if err != nil || r == nil {
			if r == nil || r.PerPage == 0 {
				err = fmt.Errorf("invalid response: %#v", r)
			}
			return nil, 0, err
		}
		return r, int(math.Ceil(float64(r.TotalCount) / float64(perPage))), nil
	})
	if err != nil {
		return nil, err
	}

	res2 := res[0]
	res2.Items = lo.FlatMap(res, func(i *Items, _ int) []Item {
		if i == nil {
			return nil
		}
		return i.Items
	})
	return res2, nil
}

func (c *CMS) GetItemsPartiallyByKey(ctx context.Context, projectIDOrAlias, modelIDOrAlias string, page, perPage int, asset bool) (*Items, error) {
	q := c.assetParam(asset)
	if page >= 1 {
		q["page"] = []string{strconv.Itoa(page)}
	}
	if perPage >= 1 {
		q["perPage"] = []string{strconv.Itoa(perPage)}
	}

	b, err := c.send(ctx, http.MethodGet, []string{"api", "projects", projectIDOrAlias, "models", modelIDOrAlias, "items"}, "", q)
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}
	defer func() { _ = b.Close() }()

	items := &Items{}
	if err := json.NewDecoder(b).Decode(items); err != nil {
		return nil, fmt.Errorf("failed to parse items: %w", err)
	}

	return items, nil
}

func (c *CMS) GetItemsByKey(ctx context.Context, projectIDOrAlias, modelIDOrAlias string, asset bool) (*Items, error) {
	var items *Items
	const perPage = 100
	for p := 1; ; p++ {
		i, err := c.GetItemsPartiallyByKey(ctx, projectIDOrAlias, modelIDOrAlias, p, perPage, asset)
		if err != nil {
			return nil, err
		}

		if i == nil || i.PerPage <= 0 {
			return nil, fmt.Errorf("invalid response: %#v", i)
		}

		if items == nil {
			items = i
		} else {
			items.Items = append(items.Items, i.Items...)
		}

		if !i.HasNext() {
			break
		}
	}

	return items, nil
}

func (c *CMS) GetItemsByKeyInParallel(ctx context.Context, projectIDOrAlias, modelIDOrAlias string, asset bool, limit int) (*Items, error) {
	const perPage = 100
	if limit <= 0 {
		limit = 5
	}

	res, err := parallel(limit, func(p int) (*Items, int, error) {
		r, err := c.GetItemsPartiallyByKey(ctx, projectIDOrAlias, modelIDOrAlias, p+1, perPage, asset)
		if err != nil || r == nil {
			if r == nil || r.PerPage == 0 {
				err = fmt.Errorf("invalid response: %#v", r)
			}
			return nil, 0, err
		}

		return r, int(math.Ceil(float64(r.TotalCount) / float64(perPage))), nil
	})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	res2 := res[0]
	res2.Items = lo.FlatMap(res, func(i *Items, _ int) []Item {
		if i == nil {
			return nil
		}
		return i.Items
	})
	return res2, nil
}

func (c *CMS) CreateItem(ctx context.Context, modelID string, fields []*Field, metadataFields []*Field) (*Item, error) {
	rb := map[string]any{
		"fields":         fields,
		"metadataFields": metadataFields,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "models", modelID, "items"}, "", rb)
	if err != nil {
		return nil, fmt.Errorf("failed to create an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(&item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) CreateItemByKey(ctx context.Context, projectID, modelID string, fields []*Field, metadataFields []*Field) (*Item, error) {
	rb := map[string]any{
		"fields":         fields,
		"metadataFields": metadataFields,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "projects", projectID, "models", modelID, "items"}, "", rb)
	if err != nil {
		return nil, fmt.Errorf("failed to create an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(&item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) UpdateItem(ctx context.Context, itemID string, fields []*Field, metadataFields []*Field) (*Item, error) {
	rb := map[string]any{
		"fields":         fields,
		"metadataFields": metadataFields,
	}

	b, err := c.send(ctx, http.MethodPatch, []string{"api", "items", itemID}, "", rb)
	if err != nil {
		return nil, fmt.Errorf("failed to update an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(&item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) DeleteItem(ctx context.Context, itemID string) error {
	b, err := c.send(ctx, http.MethodDelete, []string{"api", "items", itemID}, "", nil)
	if err != nil {
		return fmt.Errorf("failed to delete an item: %w", err)
	}
	defer func() { _ = b.Close() }()
	return nil
}

func (c *CMS) UploadAsset(ctx context.Context, projectID, url string) (string, error) {
	rb := map[string]string{
		"url": url,
	}

	b, err2 := c.send(ctx, http.MethodPost, []string{"api", "projects", projectID, "assets"}, "", rb)
	if err2 != nil {
		return "", fmt.Errorf("failed to upload an asset: %w", err2)
	}

	defer func() { _ = b.Close() }()

	body, err2 := io.ReadAll(b)
	if err2 != nil {
		return "", fmt.Errorf("failed to read body: %w", err2)
	}

	r := &Asset{}
	if err2 := json.Unmarshal(body, &r); err2 != nil {
		return "", fmt.Errorf("failed to parse body: %w", err2)
	}

	return r.ID, nil
}

func (c *CMS) UploadAssetDirectly(ctx context.Context, projectID, name string, data io.Reader, opts ...UploadAssetOption) (string, error) {
	opt := UploadAssetOption{}.Merge(opts...)

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		var err error
		defer func() {
			_ = mw.Close()
			_ = pw.CloseWithError(err)
		}()

		if opt.ContentType != "" {
			if err2 := mw.WriteField("contentType", opt.ContentType); err2 != nil {
				err = fmt.Errorf("failed to write contentType field: %w", err2)
				return
			}
		}

		if opt.ContentEncoding != "" {
			if err2 := mw.WriteField("contentEncoding", opt.ContentEncoding); err2 != nil {
				err = fmt.Errorf("failed to write contentEncoding field: %w", err2)
				return
			}
		}

		fw, err2 := mw.CreateFormFile("file", name)
		if err2 != nil {
			err = err2
			return
		}
		_, err = io.Copy(fw, data)
	}()

	b, err2 := c.send(ctx, http.MethodPost, []string{"api", "projects", projectID, "assets"}, mw.FormDataContentType(), pr)
	if err2 != nil {
		return "", fmt.Errorf("failed to upload an asset with multipart: %w", err2)
	}

	defer func() { _ = b.Close() }()

	body, err2 := io.ReadAll(b)
	if err2 != nil {
		return "", fmt.Errorf("failed to read body: %w", err2)
	}

	type res struct {
		ID string `json:"id"`
	}

	r := &res{}
	if err2 := json.Unmarshal(body, &r); err2 != nil {
		return "", fmt.Errorf("failed to parse body: %w", err2)
	}

	return r.ID, nil
}

func (c *CMS) UploadToAssetUpload(ctx context.Context, upload *AssetUpload, data io.Reader) error {
	if upload == nil {
		return errors.New("upload is nil")
	}

	ctx2 := ctx
	if c.timeout > 0 {
		ctx3, cancel := context.WithTimeout(context.Background(), c.timeout)
		ctx2 = ctx3
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx2, http.MethodPut, upload.URL, data)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", upload.ContentType)
	if upload.ContentEncoding != "" {
		req.Header.Set("Content-Encoding", upload.ContentEncoding)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload: %s", resp.Status)
	}

	return nil
}

func (c *CMS) CreateAssetByToken(ctx context.Context, projectID, token string) (*Asset, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	rb := map[string]string{
		"token": token,
	}

	b, err2 := c.send(ctx, http.MethodPost, []string{"api", "projects", projectID, "assets"}, "", rb)
	if err2 != nil {
		return nil, fmt.Errorf("failed to upload an asset: %w", err2)
	}

	defer func() { _ = b.Close() }()

	body, err2 := io.ReadAll(b)
	if err2 != nil {
		return nil, fmt.Errorf("failed to read body: %w", err2)
	}

	r := &Asset{}
	if err2 := json.Unmarshal(body, &r); err2 != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err2)
	}

	return r, nil
}

func (c *CMS) CreateAssetUpload(ctx context.Context, projectID, name string, opts ...UploadAssetOption) (*AssetUpload, error) {
	opt := UploadAssetOption{}.Merge(opts...)

	payload := map[string]string{"name": name}
	if opt.ContentType != "" {
		payload["contentType"] = opt.ContentType
	}
	if opt.ContentEncoding != "" {
		payload["contentEncoding"] = opt.ContentEncoding
	}

	b, err2 := c.send(
		ctx,
		http.MethodPost,
		[]string{"api", "projects", projectID, "assets", "uploads"},
		"",
		payload,
	)
	if err2 != nil {
		return nil, fmt.Errorf("failed to upload an asset: %w", err2)
	}

	defer func() { _ = b.Close() }()

	body, err2 := io.ReadAll(b)
	if err2 != nil {
		return nil, fmt.Errorf("failed to read body: %w", err2)
	}

	r := &AssetUpload{}
	if err2 := json.Unmarshal(body, &r); err2 != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err2)
	}

	return r, nil
}

func (c *CMS) Asset(ctx context.Context, assetID string) (*Asset, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "assets", assetID}, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get an asset: %w", err)
	}
	defer func() { _ = b.Close() }()

	a := &Asset{}
	if err := json.NewDecoder(b).Decode(a); err != nil {
		return nil, fmt.Errorf("failed to parse an asset: %w", err)
	}

	return a, nil
}

func (c *CMS) CommentToItem(ctx context.Context, itemID, content string) error {
	rb := map[string]string{
		"content": content,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "items", itemID, "comments"}, "", rb)
	if err != nil {
		return fmt.Errorf("failed to comment to item %s: %w", itemID, err)
	}
	defer func() { _ = b.Close() }()

	return nil
}

func (c *CMS) CommentToAsset(ctx context.Context, assetID, content string) error {
	rb := map[string]string{
		"content": content,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "assets", assetID, "comments"}, "", rb)
	if err != nil {
		return fmt.Errorf("failed to comment to asset %s: %w", assetID, err)
	}
	defer func() { _ = b.Close() }()

	return nil
}

func (c *CMS) send(ctx context.Context, m string, p []string, ct string, body any) (io.ReadCloser, error) {
	ctx2 := ctx
	if c.timeout > 0 {
		ctx3, cancel := context.WithTimeout(context.Background(), c.timeout)
		ctx2 = ctx3
		defer cancel()
	}

	req, err := c.request(ctx2, m, p, ct, body)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if res.StatusCode >= 300 {
		defer func() {
			_ = res.Body.Close()
		}()

		if res.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read body: %w", err)
		}

		return nil, fmt.Errorf("failed to request: code=%d, body=%s", res.StatusCode, b)
	}

	return res.Body, nil
}

func (c *CMS) request(ctx context.Context, m string, p []string, ct string, body any) (*http.Request, error) {
	if m != "GET" && ct == "" {
		ct = "application/json"
	}

	u := c.base.JoinPath(p...)
	var b io.Reader

	if m == "POST" || m == "PUT" || m == "PATCH" {
		if ct == "application/json" && body != nil {
			bb, err := json.Marshal(body)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal JSON: %w", err)
			}
			b = bytes.NewReader(bb)
		} else if strings.HasPrefix(ct, "multipart/form-data") {
			if bb, ok := body.(io.Reader); ok {
				b = bb
			}
		}
	} else if q, ok := body.(map[string][]string); ok {
		v := url.Values(q)
		u.RawQuery = v.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, m, u.String(), b)
	if err != nil {
		return nil, fmt.Errorf("failed to init request: %w", err)
	}
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	return req, nil
}

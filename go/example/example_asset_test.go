package example

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"net/http"

	cms "github.com/reearth/reearth-cms-api/go"
)

func ExampleCMS_UploadAssetDirectly_gzip() {
	ctx := context.Background()
	c, err := cms.New(cmsURL, cmsToken)
	if err != nil {
		panic(err)
	}

	// upload gzipped asset
	body := "test"
	buf := &bytes.Buffer{}
	w := gzip.NewWriter(buf)
	_, _ = w.Write([]byte(body))
	_ = w.Close()

	assetID, err := c.UploadAssetDirectly(ctx, cmsProject, "test.txt", buf, cms.UploadAssetOption{
		ContentEncoding: "gzip",
	})
	if err != nil {
		panic(fmt.Errorf("failed to upload asset: %w", err))
	}

	// read asset
	asset, err := c.Asset(ctx, assetID)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, asset.URL, nil)
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}
	if resp.Header.Get("Content-Encoding") != "gzip" {
		panic(fmt.Errorf("unexpected content encoding: %s", resp.Header.Get("Content-Encoding")))
	}

	buf = new(bytes.Buffer)
	greader, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}

	_, err = buf.ReadFrom(greader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Asset Content: %s\n", buf.String())

	// Output:
	// Asset Content: test
}

func ExampleCMS_CreateAssetUpload_gzip() {
	ctx := context.Background()
	c, err := cms.New(cmsURL, cmsToken)
	if err != nil {
		panic(err)
	}

	// upload gzipped asset
	body := "test"
	buf := &bytes.Buffer{}
	w := gzip.NewWriter(buf)
	_, _ = w.Write([]byte(body))
	_ = w.Close()

	upload, err := c.CreateAssetUpload(ctx, cmsProject, "test.txt", cms.UploadAssetOption{
		ContentEncoding: "gzip",
	})
	if err != nil {
		panic(err)
	}

	if err := c.UploadToAssetUpload(ctx, upload, buf); err != nil {
		panic(err)
	}

	asset, err := c.CreateAssetByToken(ctx, cmsProject, upload.Token)
	if err != nil {
		panic(err)
	}

	// read asset
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, asset.URL, nil)
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}
	if resp.Header.Get("Content-Encoding") != "gzip" {
		panic(fmt.Errorf("unexpected content encoding: %s", resp.Header.Get("Content-Encoding")))
	}

	buf = new(bytes.Buffer)
	greader, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}

	_, err = buf.ReadFrom(greader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Asset Content: %s\n", buf.String())

	// Output:
	// Asset Content: test
}

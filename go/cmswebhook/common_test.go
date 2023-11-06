package cmswebhook

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMergeHandlers(t *testing.T) {
	handler1 := func(r *http.Request, p *Payload) error {
		return nil
	}

	handler2 := func(r *http.Request, p *Payload) error {
		return fmt.Errorf("error")
	}

	mergedHandler := MergeHandlers([]Handler{handler1, handler2})

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	payload := &Payload{} // Assuming Payload is a struct, replace with actual implementation

	err = mergedHandler(req, payload)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

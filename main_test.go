package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
    w := httptest.NewRecorder()
    healthHandler(w, req)
    res := w.Result()
    defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }

	if len(data) > 0 {
		t.Errorf("Expected body length 0, got %v", len(data))
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
}


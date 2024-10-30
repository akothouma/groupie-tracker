package fetch

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestFetch(t *testing.T) {
    // test for a successful request
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("test response"))
    }))

    defer mockServer.Close()

	// mockServer.URL is a random assigned port on localhost
    body, err := Fetch(mockServer.URL)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if string(body) != "test response" {
        t.Errorf("Expected 'test response', got %s", string(body))
    }

    // test for a server error
    mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError)
    }))

    defer mockServer.Close()

    body, err = Fetch(mockServer.URL)
    if err != nil {
        t.Errorf("Expected no error for non-200 status, got %v", err)
    }
    if len(body) != 0 {
        t.Errorf("Expected empty body for non-200 status, got %s", string(body))
    }

    // test for an invalid URL
    body, err = Fetch("invalid-url")
    if err == nil {
        t.Error("Expected error for invalid URL, got nil")
    }
}
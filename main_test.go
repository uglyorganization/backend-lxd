package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	go Start()

	t.Run("it should return 200 when health is ok", func(t *testing.T) {
		resp, err := http.Get("http://127.0.0.1:8080/health")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Equal(%d, %d) = false", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("it should return 404 for any other endpoint", func(t *testing.T) {
		resp, err := http.Get("http://127.0.0.1:8080/nonexistent")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Equal(%d, %d) = false", resp.StatusCode, http.StatusNotFound)
		}
	})

	t.Run("it should return error on GET", func(t *testing.T) {
		_, err := http.Get("http://127.0.0.1:8081")
		if err == nil {
			t.Fatal("Expected error, connect: connection refused")
		}

		if !strings.Contains(err.Error(), "connection refused") {
			t.Fatalf("Expected error, connect: connection refused, got: %s", err.Error())
		}
	})
}

package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(time.Second * 5)

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get success", http.MethodGet, "book/12345", http.StatusOK, nil},
		{"create success", http.MethodPost, "book", http.StatusCreated, []byte(`{"id":"67890","title":"New Book Title","author":"Jane Smith","publisher":"XYZ Publications"}`)},
		{"get unknown endpoint", http.MethodGet, "unknown", http.StatusNotFound, nil},
		{"get invalid endpoint", http.MethodGet, "book/id", http.StatusNotFound, nil},
		{"unregistered route", http.MethodPut, "book", http.StatusMethodNotAllowed, nil},
		{"delete success", http.MethodDelete, "book/67890", http.StatusNoContent, nil},
	}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:8097/"+tc.endpoint, bytes.NewBuffer(tc.body))
		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue // move to the next test
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}

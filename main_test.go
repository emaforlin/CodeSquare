package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetQRCode(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/qr", nil)
	q := req.URL.Query()
	q.Add("data", "test")

	req.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, req)
	assert.Equal(t, "image/png", w.Header().Get("Content-Type"))
}

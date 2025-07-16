package router_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	router "shortlink/server"

	"github.com/stretchr/testify/assert"
)

func TestGetLinks(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/links/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "https://testlink.com/database/index012345")
}

func TestPostLink(t *testing.T) {
	router := router.SetupRouter()

	jsonData := []byte(`{"id": "1", "url": "https://newtestwebsite.com/index12345"}`)

	req, _ := http.NewRequest("POST", "/links", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "https://newtestwebsite.com/index12345")
}

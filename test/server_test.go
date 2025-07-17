package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shortlink/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func createTempLinkPost() {
// 	router := server.SetupRouter()

// 	jsonData := []byte(`{"url": "https://newtestwebsite.com/index12345"}`)

// 	req, _ := http.NewRequest("POST", "/send", bytes.NewBuffer(jsonData))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
// }

func TestPostLink(t *testing.T) {
	router := server.SetupRouter()

	jsonData := []byte(`{"url": "https://newtestwebsite.com/index12345"}`)

	req, _ := http.NewRequest("POST", "/send", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "https://newtestwebsite.com/index12345")
}

func TestGetLinks(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/links/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "https://testlink.com/database/index012345")
}

func TestPostLinkID(t *testing.T) {
	router := server.SetupRouter()

	jsonData := []byte(`{"url": "https://newtestwebsite.com/index12345"}`)

	req, _ := http.NewRequest("POST", "/send", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var response server.Link
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 3, response.ID)
	assert.Equal(t, "https://newtestwebsite.com/index12345", response.UrlOrig)

}

// func TestPostShortLink(t *testing.T) {
// 	router := server.SetupRouter()

// 	jsonData := []byte(`{"url": "https://newtestwebsite.com/index12345"}`)

// 	req, _ := http.NewRequest("POST", "/send", bytes.NewBuffer(jsonData))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	assert.Contains(t, w.Body.String(), "https://newtestwebsite.com/index12345")
// }

// func TestGetNewLink(t *testing.T) {
// 	router := server.SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/links/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Contains(t, w.Body.String(), "https://newtestwebsite.com/index12345")
// }

// func TestLinkConversionBase62(t *testing.T) {
// 	router := server.SetupRouter()
// 	createTempLinkPost()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/links/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Contains(t, w.Body.String(), "ShortUrl:n")
// }

func TestGetShortLink(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/links/100000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "FMA3SCnI")
}

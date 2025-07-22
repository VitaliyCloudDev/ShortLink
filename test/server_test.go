package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shortlink/server"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGetShortLink(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/links/100000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "FMA3SCnI")
}

func TestGetNewShortLink(t *testing.T) {
	router := server.SetupRouter()
	//creating new link in db for testing
	jsonData := []byte(`{"url": "https://newtestwebsite.com/index12345"}`)
	snd, _ := http.NewRequest("POST", "/send", bytes.NewBuffer(jsonData))
	snd.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, snd)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/links/100000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "FMA3SCnI")
}

func TestGetTotalID(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/links/total", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//It shoud be 5 ID's in base by previos tests, its bad practice to do test like this but it better than noting
	//TODO refactor test
	assert.Equal(t, "5", strings.TrimSpace(w.Body.String()))
}

func TestGetLatestID(t *testing.T) {
	router := server.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/links/latest", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "5", strings.TrimSpace(w.Body.String()))
}

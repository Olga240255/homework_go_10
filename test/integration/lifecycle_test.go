package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rinatkh/homework_backend_1/internal/httpapi"
	"github.com/rinatkh/homework_backend_1/internal/note"
)

func TestNotesLifecycle(t *testing.T) {
	router := httpapi.NewRouter(note.NewStore())

	post := request(router, http.MethodPost, "/notes", `{"title":"HTTP","text":"draft"}`)
	assertStatus(t, post, http.StatusCreated)

	get := request(router, http.MethodGet, "/notes/1", "")
	assertStatus(t, get, http.StatusOK)

	patch := request(router, http.MethodPatch, "/notes/1", `{"text":"ready"}`)
	assertStatus(t, patch, http.StatusOK)

	deleteResponse := request(router, http.MethodDelete, "/notes/1", "")
	assertStatus(t, deleteResponse, http.StatusNoContent)
	if deleteResponse.Body.Len() != 0 {
		t.Fatalf("204 не должен содержать body: %q", deleteResponse.Body.String())
	}

	missing := request(router, http.MethodGet, "/notes/1", "")
	assertStatus(t, missing, http.StatusNotFound)
}

func request(handler http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	return recorder
}

func assertStatus(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	if response.Code != want {
		t.Fatalf("status=%d, want=%d, body=%s", response.Code, want, response.Body.String())
	}
}

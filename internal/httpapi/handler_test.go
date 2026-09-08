package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rinatkh/homework_backend_1/internal/note"
)

func TestHealth(t *testing.T) {
	response := performRequest(NewRouter(note.NewStore()), http.MethodGet, "/health", "")
	assertStatus(t, response, http.StatusOK)
}

func TestGetEmptyList(t *testing.T) {
	response := performRequest(NewRouter(note.NewStore()), http.MethodGet, "/notes", "")
	assertStatus(t, response, http.StatusOK)
	if strings.TrimSpace(response.Body.String()) != `[]` {
		t.Fatalf("пустой список должен быть [], получили %s", response.Body.String())
	}
}

func TestPostValid(t *testing.T) {
	response := performRequest(NewRouter(note.NewStore()), http.MethodPost, "/notes", `{"title":"HTTP","text":"request and response"}`)
	assertStatus(t, response, http.StatusCreated)
	if response.Header().Get("Location") != "/notes/1" {
		t.Fatalf("неожиданный Location: %q", response.Header().Get("Location"))
	}
	var item note.Note
	if err := json.NewDecoder(response.Body).Decode(&item); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if item.ID != 1 || item.Title != "HTTP" {
		t.Fatalf("неожиданная заметка: %#v", item)
	}
}

func TestPostErrors(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{name: "malformed JSON", body: `{"title":`},
		{name: "empty title", body: `{"title":"   ","text":"x"}`},
		{name: "unknown field", body: `{"titel":"опечатка"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := performRequest(NewRouter(note.NewStore()), http.MethodPost, "/notes", tt.body)
			assertStatus(t, response, http.StatusBadRequest)
		})
	}
}

func TestGetExistingAndUnknown(t *testing.T) {
	store := note.NewStore()
	created, err := store.Create("Заметка", "текст")
	if err != nil {
		t.Fatalf("подготовка теста: %v", err)
	}
	router := NewRouter(store)

	response := performRequest(router, http.MethodGet, "/notes/1", "")
	assertStatus(t, response, http.StatusOK)
	if !strings.Contains(response.Body.String(), created.Title) {
		t.Fatalf("ответ не содержит заметку: %s", response.Body.String())
	}

	response = performRequest(router, http.MethodGet, "/notes/999", "")
	assertStatus(t, response, http.StatusNotFound)
}

func TestInvalidID(t *testing.T) {
	response := performRequest(NewRouter(note.NewStore()), http.MethodGet, "/notes/abc", "")
	assertStatus(t, response, http.StatusBadRequest)
}

func TestPatch(t *testing.T) {
	store := note.NewStore()
	if _, err := store.Create("До", "старый текст"); err != nil {
		t.Fatalf("подготовка теста: %v", err)
	}
	response := performRequest(NewRouter(store), http.MethodPatch, "/notes/1", `{"text":"новый текст"}`)
	assertStatus(t, response, http.StatusOK)
	if !strings.Contains(response.Body.String(), "новый текст") {
		t.Fatalf("PATCH не изменил текст: %s", response.Body.String())
	}
}

func TestDelete(t *testing.T) {
	store := note.NewStore()
	if _, err := store.Create("Удалить", ""); err != nil {
		t.Fatalf("подготовка теста: %v", err)
	}
	router := NewRouter(store)
	response := performRequest(router, http.MethodDelete, "/notes/1", "")
	assertStatus(t, response, http.StatusNoContent)
	if response.Body.Len() != 0 {
		t.Fatalf("204 не должен содержать body: %q", response.Body.String())
	}
}

func performRequest(handler http.Handler, method, path, body string) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	if body != "" {
		request.Header.Set("Content-Type", "application/json")
	}
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	return response
}

func assertStatus(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	if response.Code != want {
		t.Fatalf("status=%d, want=%d, body=%s", response.Code, want, response.Body.String())
	}
}

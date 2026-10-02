package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rinatkh/homework_backend_1/internal/note"
)

// Эти пять тестов ученица пишет сама после реализации API. Основные
// проверочные тесты уже находятся рядом и менять их не нужно.

func TestStudentPostInvalidJSON(t *testing.T) {
	// TODO(08.1): создайте POST /notes со сломанным JSON и проверьте 400.
	//t.Skip("TODO(08.1)")
	invalidJSON := `{"title":`
	req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	NewRouter(note.NewStore()).ServeHTTP(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("ожидался статус %d, получили %d", http.StatusBadRequest, res.Code)
	}
}

func TestStudentGetUnknown(t *testing.T) {
	// TODO(08.2): создайте GET /notes/999 и проверьте 404.
	//t.Skip("TODO(08.2)")
	req := httptest.NewRequest(http.MethodGet, "/notes/999", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	NewRouter(note.NewStore()).ServeHTTP(res, req)
	if res.Code != http.StatusNotFound {
		t.Errorf("ожидался статус %d, получили %d", http.StatusNotFound, res.Code)
	}
}

func TestStudentPatch(t *testing.T) {
	// TODO(08.3): создайте заметку, измените только text и проверьте ответ.
	//t.Skip("TODO(08.3)")
	textJSON := `{"text":"текст"}`
	req := httptest.NewRequest(http.MethodPatch, "/notes/1", strings.NewReader(textJSON))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	NewRouter(note.NewStore()).ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("ожидался статус %d, получили %d", http.StatusOK, res.Code)
	}
}

func TestStudentDeleteNoBody(t *testing.T) {
	// TODO(08.4): проверьте status 204 и пустой body.
	//t.Skip("TODO(08.4)")
	store := note.NewStore()
	if _, err := store.Create("Удалить", ""); err != nil {
		t.Fatalf("подготовка теста: %v", err)
	}
	req := httptest.NewRequest(http.MethodDelete, "/notes/1", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	NewRouter(store).ServeHTTP(res, req)
	if res.Code != http.StatusNoContent {
		t.Errorf("ожидался статус %d, получили %d", http.StatusNoContent, res.Code)
	}
	if res.Body.Len() != 0 {
		t.Errorf("ожидалось пустое тело ответа, получили: %s", res.Body.String())
	}
}

func TestStudentLifecycle(t *testing.T) {
	// TODO(08.5): пройдите POST → GET → PATCH → DELETE одним тестом.
	//t.Skip("TODO(08.5)")
	TestStudentPostInvalidJSON(t)
	TestStudentGetUnknown(t)
	TestStudentPatch(t)
	TestStudentDeleteNoBody(t)
}

// Package httpapi содержит маршруты и handler-функции Notes API.
package httpapi

import (
	"net/http"

	"github.com/rinatkh/homework_backend_1/internal/note"
)

type Handler struct {
	store *note.Store
}

// NewRouter пока регистрирует только готовый справочный endpoint /health.
// Добавляйте маршруты по порядку из docs/task.md.
func NewRouter(store *note.Store) http.Handler {
	h := &Handler{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.Health)

	// TODO(03.1): зарегистрируйте GET /notes.
	// TODO(03.2): зарегистрируйте GET /notes/{id}.
	// TODO(03.3): зарегистрируйте POST /notes.
	// TODO(03.4): зарегистрируйте PATCH /notes/{id}.
	// TODO(03.5): зарегистрируйте DELETE /notes/{id}.

	return mux
}

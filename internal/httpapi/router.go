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
	//что это?
	h := &Handler{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.Health)

	// TODO(03.1): зарегистрируйте GET /notes.
	mux.HandleFunc("GET /notes", h.ListNotes)
	// TODO(03.2): зарегистрируйте GET /notes/{id}.
	mux.HandleFunc("GET /notes/{id}", h.GetNote)
	// TODO(03.3): зарегистрируйте POST /notes.
	mux.HandleFunc("POST /notes", h.CreateNote)
	// TODO(03.4): зарегистрируйте PATCH /notes/{id}.
	mux.HandleFunc("PATCH /notes/{id}", h.PatchNote)
	// TODO(03.5): зарегистрируйте DELETE /notes/{id}.
	mux.HandleFunc("DELETE /notes/{id}", h.DeleteNote)
	return mux
}

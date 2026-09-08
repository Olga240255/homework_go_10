package httpapi

import (
	"net/http"

	"github.com/rinatkh/homework_backend_1/internal/note"
)

type healthResponse struct {
	Status string `json:"status"`
}

// TODO(01.2): добавьте JSON-теги title и text для POST.
type createNoteRequest struct {
	Title string
	Text  string
}

type patchNoteRequest struct {
	// TODO(01.3): добавьте JSON-теги и объясните, почему поля — указатели.
	Title *string
	Text  *string
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
}

// ListNotes обрабатывает GET /notes.
// TODO(04.1): получите список из store.
// TODO(04.2): верните JSON-массив и status 200.
func (h *Handler) ListNotes(w http.ResponseWriter, _ *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: ListNotes")
}

// GetNote обрабатывает GET /notes/{id}.
// TODO(04.3): получите ID из path через r.PathValue("id").
// TODO(04.4): верните 400 для плохого ID и 404 для отсутствующей заметки.
// TODO(04.5): для найденной заметки верните 200 и JSON.
func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: GetNote")
}

// CreateNote обрабатывает POST /notes.
// TODO(05.4): прочитайте createNoteRequest и вызовите Store.Create.
// TODO(05.5): верните 201, Location и JSON созданной заметки.
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: CreateNote")
}

// PatchNote обрабатывает PATCH /notes/{id}.
// TODO(06.1): прочитайте ID и JSON с частичными изменениями.
// TODO(06.2): вызовите Store.Patch с указателями из запроса.
// TODO(06.3): сопоставьте ошибки со status 400/404, успех — с 200.
func (h *Handler) PatchNote(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: PatchNote")
}

// DeleteNote обрабатывает DELETE /notes/{id}.
// TODO(06.4): удалите заметку и верните 404, если её нет.
// TODO(06.5): при успехе верните 204 и не пишите body.
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: DeleteNote")
}

// parseID превращает path parameter в положительный int.
// TODO(07.4): invalid ID должен вернуть 400.
func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	writeError(w, http.StatusNotImplemented, "TODO: parseID")
	return 0, false
}

// TODO(07.5): проверьте, что DELETE 204 не добавляет даже `{}` в body.

var _ = note.Changes{}

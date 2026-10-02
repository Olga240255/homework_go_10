package httpapi

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rinatkh/homework_backend_1/internal/note"
	//"golang.org/x/crypto/openpgp/errors"
)

type healthResponse struct {
	Status string `json:"status"`
}

// TODO(01.2): добавьте JSON-теги title и text для POST.
type createNoteRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type patchNoteRequest struct {
	// TODO(01.3): добавьте JSON-теги и объясните, почему поля — указатели.
	//это изменение и пользователь может не передавать часть полей и нам нужно отловить nil
	//если не указатель то nil и "" не отличить
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
}

// ListNotes обрабатывает GET /notes.
// TODO(04.1): получите список из store.
// TODO(04.2): верните JSON-массив и status 200.
func (h *Handler) ListNotes(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, h.store.List())
}

// GetNote обрабатывает GET /notes/{id}.
// TODO(04.3): получите ID из path через r.PathValue("id").
// TODO(04.4): верните 400 для плохого ID и 404 для отсутствующей заметки.
// TODO(04.5): для найденной заметки верните 200 и JSON.
func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	//что это?
	idnum, ok := parseID(w, r)
	if ok != true {
		return
	}
	if idnum < 0 {
		writeError(w, http.StatusBadRequest, note.ErrNotFound.Error())
		return
	}
	res, err := h.store.Get(idnum)
	if errors.Is(err, note.ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, res)
}

// CreateNote обрабатывает POST /notes.
// TODO(05.4): прочитайте createNoteRequest и вызовите Store.Create.
// TODO(05.5): верните 201, Location и JSON созданной заметки.
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var input note.Note
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.store.Create(input.Title, input.Text)
	if errors.Is(err, note.ErrEmptyTitle) {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/notes/%d", res.ID))
	writeJSON(w, http.StatusCreated, res)

}

// PatchNote обрабатывает PATCH /notes/{id}.
// TODO(06.1): прочитайте ID и JSON с частичными изменениями.
// TODO(06.2): вызовите Store.Patch с указателями из запроса.
// TODO(06.3): сопоставьте ошибки со status 400/404, успех — с 200.
func (h *Handler) PatchNote(w http.ResponseWriter, r *http.Request) {
	idnum, ok := parseID(w, r)
	if ok != true {
		return
	}
	var input note.Changes
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, r.PathValue("id")+note.ErrNotFound.Error())
		return
	}
	res, err := h.store.Patch(idnum, input)
	if errors.Is(err, note.ErrNoChanges) {
		writeError(w, http.StatusNotFound, "4"+note.ErrNotFound.Error())
		return
	}
	writeJSON(w, http.StatusOK, res)
}

// DeleteNote обрабатывает DELETE /notes/{id}.
// TODO(06.4): удалите заметку и верните 404, если её нет.
// TODO(06.5): при успехе верните 204 и не пишите body.
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	idnum, ok := parseID(w, r)
	if ok != true {
		return
	}
	err := h.store.Delete(idnum)
	if errors.Is(err, note.ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
	}
	w.WriteHeader(http.StatusNoContent)
}

// parseID превращает path parameter в положительный int.
// TODO(07.4): invalid ID должен вернуть 400.
func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	rawID := r.PathValue("id")
	id, err := strconv.Atoi(rawID)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "id должен быть положительным целым числом")
		return 0, false
	}
	return id, true
}

// TODO(07.5): проверьте, что DELETE 204 не добавляет даже `{}` в body.

var _ = note.Changes{}

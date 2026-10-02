package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var errDecodeNotImplemented = errors.New("TODO: чтение JSON ещё не реализовано")

// decodeJSON преобразует JSON body в Go-структуру.
// TODO(05.1): используйте json.NewDecoder(r.Body).
// TODO(05.2): включите DisallowUnknownFields.
// TODO(05.3): отклоните второй JSON-объект после первого.
func decodeJSON(r *http.Request, dst any) error {
	maxBytes := int64(1024 * 1024)
	limitedBody := http.MaxBytesReader(nil, r.Body, maxBytes)
	input, err := io.ReadAll(limitedBody)
	if err != nil {
		return err
	}
	err = json.Unmarshal(input, dst)
	if err != nil {
		return err
	}
	return nil
}

// writeJSON задаёт Content-Type, status и только затем пишет body.
// Эта функция дана как рабочий пример правильного порядка ответа.
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

// writeError сохраняет единый формат {"error":"..."}.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// TODO(07.1): объясните, почему Header вызывается раньше WriteHeader.
//сначала заголовок передается, потом тело
// TODO(07.2): убедитесь, что все JSON-ответы имеют Content-Type.
//
// TODO(07.3): для ошибок используйте status, а не только поле error в body.

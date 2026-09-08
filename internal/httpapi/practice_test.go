package httpapi

import "testing"

// Эти пять тестов ученица пишет сама после реализации API. Основные
// проверочные тесты уже находятся рядом и менять их не нужно.

func TestStudentPostInvalidJSON(t *testing.T) {
	// TODO(08.1): создайте POST /notes со сломанным JSON и проверьте 400.
	t.Skip("TODO(08.1)")
}

func TestStudentGetUnknown(t *testing.T) {
	// TODO(08.2): создайте GET /notes/999 и проверьте 404.
	t.Skip("TODO(08.2)")
}

func TestStudentPatch(t *testing.T) {
	// TODO(08.3): создайте заметку, измените только text и проверьте ответ.
	t.Skip("TODO(08.3)")
}

func TestStudentDeleteNoBody(t *testing.T) {
	// TODO(08.4): проверьте status 204 и пустой body.
	t.Skip("TODO(08.4)")
}

func TestStudentLifecycle(t *testing.T) {
	// TODO(08.5): пройдите POST → GET → PATCH → DELETE одним тестом.
	t.Skip("TODO(08.5)")
}

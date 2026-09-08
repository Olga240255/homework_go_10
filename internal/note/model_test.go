package note

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNoteJSONTags(t *testing.T) {
	data, err := json.Marshal(Note{ID: 1, Title: "HTTP", Text: "request → response"})
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	got := string(data)
	for _, field := range []string{`"id"`, `"title"`, `"text"`} {
		if !strings.Contains(got, field) {
			t.Fatalf("ожидали поле %s в %s", field, got)
		}
	}
	if strings.Contains(got, `"ID"`) || strings.Contains(got, `"Title"`) {
		t.Fatalf("JSON должен использовать имена из тегов: %s", got)
	}
}

func TestJSONDash(t *testing.T) {
	data, err := json.Marshal(PublicExample{Title: "Можно показать", Secret: "нельзя показать"})
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	if strings.Contains(string(data), "нельзя показать") || strings.Contains(string(data), "Secret") {
		t.Fatalf("поле с json:\"-\" попало в ответ: %s", data)
	}
}

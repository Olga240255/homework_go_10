package note

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestStoreEmptyList(t *testing.T) {
	items := NewStore().List()
	if items == nil || len(items) != 0 {
		t.Fatalf("пустой store должен вернуть [], получили %#v", items)
	}
}

func TestStoreCreateAndGet(t *testing.T) {
	store := NewStore()
	created, err := store.Create("  Первая заметка  ", "текст")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if created.ID != 1 || created.Title != "Первая заметка" {
		t.Fatalf("неожиданная заметка: %#v", created)
	}

	got, err := store.Get(created.ID)
	if err != nil || got != created {
		t.Fatalf("Get: got=%#v err=%v", got, err)
	}
}

func TestStoreRejectsEmptyTitle(t *testing.T) {
	if _, err := NewStore().Create("   ", "text"); !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("ожидали ErrEmptyTitle, получили %v", err)
	}
}

func TestStorePatchAndDelete(t *testing.T) {
	store := NewStore()
	created, err := store.Create("До", "старый текст")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	newText := "новый текст"
	patched, err := store.Patch(created.ID, Changes{Text: &newText})
	if err != nil {
		t.Fatalf("Patch: %v", err)
	}
	if patched.Title != "До" || patched.Text != newText {
		t.Fatalf("PATCH изменил лишнее или не применился: %#v", patched)
	}
	if err := store.Delete(created.ID); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if _, err := store.Get(created.ID); !errors.Is(err, ErrNotFound) {
		t.Fatalf("после удаления ожидали ErrNotFound, получили %v", err)
	}
}

func TestStoreConcurrentCreate(t *testing.T) {
	store := NewStore()
	const count = 30
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			_, _ = store.Create(fmt.Sprintf("note %d", index), "")
		}(i)
	}
	wg.Wait()
	if got := len(store.List()); got != count {
		t.Fatalf("ожидали %d заметок, получили %d", count, got)
	}
}

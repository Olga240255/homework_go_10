package note

import (
	"errors"
	"slices"
	"strings"
	"sync"
)

var (
	ErrNotFound       = errors.New("заметка не найдена")
	ErrEmptyTitle     = errors.New("заголовок заметки не должен быть пустым")
	ErrNoChanges      = errors.New("не указано ни одного поля для изменения")
	ErrNotImplemented = errors.New("TODO: функция ещё не реализована")
)

// Store хранит заметки в памяти процесса.
//
// mutex уже добавлен, потому что net/http может выполнять handler-функции
// одновременно. Ваша задача — правильно использовать RLock/Lock в методах.
type Store struct {
	mu     sync.RWMutex
	notes  map[int]Note
	nextID int
}

func NewStore() *Store {
	return &Store{notes: make(map[int]Note), nextID: 1}
}

// List возвращает заметки по возрастанию ID.
// TODO(02.1): сделайте копию значений map и отсортируйте результат.
func (s *Store) List() []Note {
	if s == nil || s.notes == nil {
		return []Note{}
	}
	res := make([]Note, 0, len(s.notes))
	for _, val := range s.notes {
		res = append(res, val)
	}
	slices.SortFunc(res, func(a, b Note) int {
		if a.ID < b.ID {
			return -1
		}
		if a.ID > b.ID {
			return 1
		}
		return 0
	})
	return res
}

// Get возвращает заметку по ID или ErrNotFound.
// TODO(02.2): реализуйте чтение под RLock.
func (s *Store) Get(id int) (Note, error) {
	if s == nil {
		return Note{}, nil
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, val := range s.notes {
		if i == id {
			return val, nil
		}
	}
	return Note{}, ErrNotFound
}

// Create проверяет title, назначает ID и сохраняет заметку.
// TODO(02.3): обрежьте пробелы и выполните запись под Lock.
func (s *Store) Create(title, text string) (Note, error) {
	if s == nil {
		return Note{}, nil
	}
	t := strings.TrimSpace(title)
	if t == "" {
		return Note{}, ErrEmptyTitle
	} else {
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.notes == nil {
			s.notes = make(map[int]Note)
		}
		ID := s.nextID
		s.notes[ID] = Note{ID: ID, Title: t, Text: text}
		s.nextID++
		return s.notes[ID], nil
	}
}

// Patch изменяет только переданные поля.
// TODO(02.4): различайте nil и явно переданную пустую строку.
func (s *Store) Patch(id int, changes Changes) (Note, error) {
	if s == nil {
		return Note{}, ErrNoChanges
	}
	if changes.Title == nil && changes.Text == nil {
		return Note{}, ErrNoChanges
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	d := s.notes[id]
	if changes.Title != nil {
		d.Title = *changes.Title
	}
	if changes.Text != nil {
		d.Text = *changes.Text
	}
	s.notes[id] = d
	return s.notes[id], nil
}

// Delete удаляет заметку или возвращает ErrNotFound.
// TODO(02.5): выполните проверку и delete под Lock.
func (s *Store) Delete(id int) error {
	if s == nil {
		return ErrNotFound
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.notes[id]; !exists {
		return ErrNotFound
	}
	delete(s.notes, id)
	return nil

}

// Как не нужно: не обращайтесь к s.notes без mutex. Два HTTP-запроса могут
// работать одновременно, а конкурентная запись в обычную map небезопасна.

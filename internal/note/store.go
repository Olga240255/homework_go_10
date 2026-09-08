package note

import (
	"errors"
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
	return []Note{}
}

// Get возвращает заметку по ID или ErrNotFound.
// TODO(02.2): реализуйте чтение под RLock.
func (s *Store) Get(id int) (Note, error) {
	return Note{}, ErrNotImplemented
}

// Create проверяет title, назначает ID и сохраняет заметку.
// TODO(02.3): обрежьте пробелы и выполните запись под Lock.
func (s *Store) Create(title, text string) (Note, error) {
	return Note{}, ErrNotImplemented
}

// Patch изменяет только переданные поля.
// TODO(02.4): различайте nil и явно переданную пустую строку.
func (s *Store) Patch(id int, changes Changes) (Note, error) {
	return Note{}, ErrNotImplemented
}

// Delete удаляет заметку или возвращает ErrNotFound.
// TODO(02.5): выполните проверку и delete под Lock.
func (s *Store) Delete(id int) error {
	return ErrNotImplemented
}

// Как не нужно: не обращайтесь к s.notes без mutex. Два HTTP-запроса могут
// работать одновременно, а конкурентная запись в обычную map небезопасна.

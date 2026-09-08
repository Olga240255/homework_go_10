// Package note содержит модель заметки и хранилище домашнего Notes API.
package note

// Note — одна заметка, которую API возвращает клиенту.
//
// TODO(01.1): добавьте JSON-теги для id, title и text.
type Note struct {
	ID    int
	Title string
	Text  string
}

// Changes описывает частичное изменение заметки.
//
// TODO(01.4): объясните в комментарии, почему здесь нужны указатели.
type Changes struct {
	Title *string
	Text  *string
}

// PublicExample нужен только для упражнения с json:"-".
// TODO(01.5): добавьте теги так, чтобы Title попал в JSON под именем title,
// а Secret никогда не попал во внешний ответ.
type PublicExample struct {
	Title  string
	Secret string
}

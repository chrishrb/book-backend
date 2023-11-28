package books

import (
	"fmt"
)

type InMemoryStore struct {
	books []Book
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		books: []Book{},
	}
}

func (s *InMemoryStore) Add(book Book) (*Book, error) {
	book.ID = generateUUID()
	s.books = append(s.books, book)
	return &book, nil
}

func (s *InMemoryStore) ListByID(id string) (*Book, error) {
	for _, b := range s.books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, fmt.Errorf("%s not found", id)
}

func (s *InMemoryStore) ListAll() (*[]Book, error) {
	return &s.books, nil
}

func (s *InMemoryStore) Update(id string, book Book) (*Book, error) {
	for i, b := range s.books {
		if b.ID == id {
			s.books[i].Title = book.Title
			s.books[i].Description = book.Description
			s.books[i].NoOfPages = book.NoOfPages
			s.books[i].Year = book.Year
			return &s.books[i], nil
		}
	}
	return nil, fmt.Errorf("%s not found", id)
}

func (s *InMemoryStore) Delete(id string) error {
	for index, b := range s.books {
		if b.ID == id {
			s.books = append(s.books[:index], s.books[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%s not found", id)
}

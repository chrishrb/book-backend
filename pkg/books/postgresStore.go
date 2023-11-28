package books

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	connection *sqlx.DB
}

func NewPostgresStore(connectionString string) *PostgresStore {
	// connect to db
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return &PostgresStore{
		connection: db,
	}
}

func (s PostgresStore) Add(book Book) (*Book, error) {
	query := `INSERT INTO books (id, title, description, no_of_pages, year) 
						VALUES (:id, :title, :description, :no_of_pages, :year)`

	insertBook := Book{
		ID:          generateUUID(),
		Title:       book.Title,
		Description: book.Description,
		NoOfPages:       book.NoOfPages,
		Year: book.Year,
	}
	_, err := s.connection.NamedExec(query, &insertBook)

	if err != nil {
		return nil, err
	}

	return &insertBook, nil
}

func (s PostgresStore) ListByID(id string) (*Book, error) {
	book := Book{}
	err := s.connection.Get(&book, "SELECT * from books WHERE id=$1", id)

	if err != nil || book == (Book{}) {
		return nil, fmt.Errorf("id %s not found", id)
	}

	return &book, nil
}

func (s PostgresStore) ListAll() (*[]Book, error) {
	books := []Book{}
	err := s.connection.Select(&books, "SELECT * FROM books")

	if err != nil {
		return nil, err
	}

	return &books, nil
}

func (s PostgresStore) Update(id string, book Book) (*Book, error) {
	book.ID = id
	query := "UPDATE books SET title=:title, description=:description, no_of_pages=:no_of_pages, year=:year WHERE id=:id"
	_, err := s.connection.NamedExec(query, book)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s PostgresStore) Delete(id string) error {
	_, err := s.connection.Exec("DELETE from books WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

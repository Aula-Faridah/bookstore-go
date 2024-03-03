package repositories

import (
	"bdpit/bookstore-go/internals/models"

	_ "github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	*sqlx.DB
}

func InitBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db}
}

func (b *BookRepo) FindAll() ([]models.BookModel, error) {
	query := "SELECT * FROM books"
	result := &[]models.BookModel{}
	if err := b.Select(result, query); err != nil {
		return nil, err
	}
	return *result, nil
}

func (b *BookRepo) FindById(id int) ([]models.BookModel, error) {
	query := "SELECT * FROM books WHERE id = ?"
	result := []models.BookModel{}

	if err := b.Select(&result, query, id); err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BookRepo) SaveBook(body models.BookModel) error {
	query := "INSERT INTO books (title,description,author) VALUES(?,?,?)"
	if _, err := b.Exec(query, body.Title, body.Description, body.Author); err != nil {
		return err
	}
	return nil
}

func (b *BookRepo) UpdateBook(body models.BookModel) error {
	query := "UPDATE books SET title=?, description=?, author=? WHERE id=?"
	if _, err := b.Exec(query, body.Title, body.Description, body.Author, body.Id); err != nil {
		return err
	}
	return nil
}

func (b *BookRepo) DeleteBookById(id int) error {
	query := "DELETE FROM books WHERE id=?"

	if _, err := b.Exec(query,id); err != nil {
		return err
	}
	return nil
}

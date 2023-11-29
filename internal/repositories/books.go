package repositories

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/wildanfaz/simple-golang-echo/internal/models"
)

type BooksRepo struct {
	db *sql.DB
}

type BooksRepository interface {
	AddBook(ctx context.Context, payload *models.Book) error
	ListBooks(ctx context.Context, payload *models.ListBooksPayload) ([]models.Book, error)
	GetBook(ctx context.Context, payload *models.Book) (*models.Book, error)
	UpdateBook(ctx context.Context, payload *models.UpdateBook) error
	DeleteBook(ctx context.Context, payload *models.Book) error
}

func NewBooksRepository(db *sql.DB) BooksRepository {
	return &BooksRepo{
		db: db,
	}
}

func (r *BooksRepo) AddBook(ctx context.Context, payload *models.Book) error {
	q := `
	INSERT INTO books
		(title, description, author)
	VALUES(?, ?, ?)
	`

	if _, err := r.db.ExecContext(ctx, q,
		payload.Title,
		payload.Description,
		payload.Author,
	); err != nil {
		return err
	}

	return nil
}

func (r *BooksRepo) ListBooks(ctx context.Context, payload *models.ListBooksPayload) ([]models.Book, error) {
	var (
		book  models.Book
		books models.Books
	)

	q, values := queryListBooks(*payload)

	rows, err := r.db.QueryContext(ctx, q, values...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.Author, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BooksRepo) GetBook(ctx context.Context, payload *models.Book) (*models.Book, error) {
	var book models.Book

	q := `
	SELECT id, title, description, author, created_at, updated_at
		FROM books
	WHERE id = ?
	`

	err := r.db.QueryRowContext(ctx, q, payload.Id).Scan(&book.Id, &book.Title, &book.Description, &book.Author, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BooksRepo) UpdateBook(ctx context.Context, payload *models.UpdateBook) error {
	q, values := queryUpdateBook(*payload)

	_, err := r.db.ExecContext(ctx, q, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *BooksRepo) DeleteBook(ctx context.Context, payload *models.Book) error {
	q := `
	DELETE FROM books
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, q, payload.Id)
	if err != nil {
		return err
	}

	return nil
}

func queryListBooks(payload models.ListBooksPayload) (string, []any) {
	var values []any

	q := `
	SELECT id, title, description, author, created_at, updated_at
		FROM books
	`

	if !reflect.ValueOf(payload.Book).IsZero() {
		q += `
		WHERE `

		if payload.Title != "" {
			q += `title LIKE ? AND `
			payload.Title = "%" + payload.Title + "%"
			values = append(values, payload.Title)
		}

		if payload.Description != "" {
			q += `description LIKE ? AND `
			payload.Description = "%" + payload.Description + "%"
			values = append(values, payload.Description)
		}

		if payload.Author != "" {
			q += `author LIKE ? AND `
			payload.Author = "%" + payload.Author + "%"
			values = append(values, payload.Author)
		}

		q = strings.TrimSuffix(q, `AND `)
	}

	q += `ORDER BY id DESC `

	if payload.Limit > 0 {
		q += `LIMIT ? `
		values = append(values, payload.Limit)
	} else {
		q += `LIMIT 10 `
	}

	if payload.Page > 0 {
		q += `OFFSET ? `
		payload.Page = (payload.Page - 1) * payload.Limit
		values = append(values, payload.Page)
	}

	return q, values
}

func queryUpdateBook(payload models.UpdateBook) (string, []any) {
	var values []any

	q := `
	UPDATE books
	`

	if !reflect.ValueOf(payload).IsZero() {
		q += `
		SET
		`

		if payload.Title != "" {
			q += `title = ?, `
			values = append(values, payload.Title)
		}

		if payload.Description != "" {
			q += `description = ?, `
			values = append(values, payload.Description)
		}

		if payload.Author != "" {
			q += `author = ?, `
			values = append(values, payload.Author)
		}

		q = strings.TrimSuffix(q, `, `)

		q += `
		WHERE id = ?
		`
		values = append(values, payload.Id)
	}

	return q, values
}

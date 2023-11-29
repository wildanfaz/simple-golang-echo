package models

import "time"

type (
	Book struct {
		Id          uint      `param:"id" query:"id" json:"id"`
		Title       string    `query:"title" json:"title"`
		Description string    `query:"description" json:"description"`
		Author      string    `query:"author" json:"author"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	ListBooksPayload struct {
		Book
		Pagination
	}

	UpdateBook struct {
		Id          uint   `param:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Author      string `json:"author"`
	}

	Books []Book

	Pagination struct {
		Page  int `query:"page" json:"-"`
		Limit int `query:"limit" json:"-"`
	}
)

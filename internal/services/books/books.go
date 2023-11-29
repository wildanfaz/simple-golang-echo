package books

import (
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/simple-golang-echo/internal/repositories"
)

type Service struct {
	booksRepo repositories.BooksRepository
	log       *logrus.Logger
}

func NewService(booksRepo repositories.BooksRepository, log *logrus.Logger) *Service {
	return &Service{
		booksRepo: booksRepo,
		log:       log,
	}
}

package books

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/simple-golang-echo/internal/helpers"
	"github.com/wildanfaz/simple-golang-echo/internal/models"
)

func (s *Service) GetBook(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.Book
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("GetBook - Bind Error", err)
		return response.AsError().
			WithMessage("Invalid Payload").
			MakeJSON(c, http.StatusBadRequest)
	}

	if payload.Id == 0 {
		s.log.Errorln("GetBook - Invalid Payload")
		return response.AsError().
			WithMessage("Id must be greater than 0").
			MakeJSON(c, http.StatusBadRequest)
	}

	book, err := s.booksRepo.GetBook(c.Request().Context(), &payload)
	if err != nil && err != sql.ErrNoRows {
		s.log.Errorln("GetBook - Get Book Repository Error", err)
		return response.AsError().
			WithMessage("Internal Server Error").
			MakeJSON(c, http.StatusInternalServerError)
	}

	if err == sql.ErrNoRows {
		s.log.Errorln("GetBook - Book Not Found")
		return response.AsError().
			WithMessage("Book Not Found").
			MakeJSON(c, http.StatusNotFound)
	}

	s.log.Println("GetBook - Success")
	return response.WithMessage("Get Book Success").
		WithData(book).
		MakeJSON(c, http.StatusOK)
}

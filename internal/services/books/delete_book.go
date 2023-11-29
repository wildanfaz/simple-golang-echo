package books

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/simple-golang-echo/internal/helpers"
	"github.com/wildanfaz/simple-golang-echo/internal/models"
)

func (s *Service) DeleteBook(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.Book
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("DeleteBook - Bind Error", err)
		return response.AsError().
			WithMessage("Invalid Payload").
			MakeJSON(c, http.StatusBadRequest)
	}

	if payload.Id == 0 {
		s.log.Errorln("DeleteBook - Invalid Payload")
		return response.AsError().
			WithMessage("Id must be greater than 0").
			MakeJSON(c, http.StatusBadRequest)
	}

	_, err = s.booksRepo.GetBook(c.Request().Context(), &payload)
	if err != nil && err != sql.ErrNoRows {
		s.log.Errorln("DeleteBook - Get Book Repository Error", err)
		return response.AsError().
			WithMessage("Internal Server Error").
			MakeJSON(c, http.StatusInternalServerError)
	}

	if err == sql.ErrNoRows {
		s.log.Errorln("DeleteBook - Book Not Found")
		return response.AsError().
			WithMessage("Book Not Found").
			MakeJSON(c, http.StatusNotFound)
	}

	s.booksRepo.DeleteBook(c.Request().Context(), &payload)

	s.log.Println("DeleteBook - Success")
	return response.WithMessage("Delete Book Success").
		MakeJSON(c, http.StatusOK)
}

package books

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/simple-golang-echo/internal/helpers"
	"github.com/wildanfaz/simple-golang-echo/internal/models"
)

func (s *Service) UpdateBook(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.UpdateBook
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("UpdateBook - Bind Error", err)
		return response.AsError().
			WithMessage("Invalid Payload").
			MakeJSON(c, http.StatusBadRequest)
	}

	if payload.Id == 0 {
		s.log.Errorln("UpdateBook - Invalid Payload")
		return response.AsError().
			WithMessage("Id must be greater than 0").
			MakeJSON(c, http.StatusBadRequest)
	}

	_, err = s.booksRepo.GetBook(c.Request().Context(), &models.Book{Id: payload.Id})
	if err != nil && err != sql.ErrNoRows {
		s.log.Errorln("UpdateBook - Get Book Repository Error", err)
		return response.AsError().
			WithMessage("Internal Server Error").
			MakeJSON(c, http.StatusInternalServerError)
	}

	if err == sql.ErrNoRows {
		s.log.Errorln("UpdateBook - Book Not Found")
		return response.AsError().
			WithMessage("Book Not Found").
			MakeJSON(c, http.StatusNotFound)
	}

	err = s.booksRepo.UpdateBook(c.Request().Context(), &payload)
	if err != nil {
		s.log.Errorln("UpdateBook - Update Book Repository Error", err)
		return response.AsError().
			WithMessage("Internal Server Error").
			MakeJSON(c, http.StatusInternalServerError)
	}

	s.log.Println("UpdateBook - Success")
	return response.WithMessage("Update Book Success").
		MakeJSON(c, http.StatusOK)
}

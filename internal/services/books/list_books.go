package books

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/simple-golang-echo/internal/helpers"
	"github.com/wildanfaz/simple-golang-echo/internal/models"
)

func (s *Service) ListBooks(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.ListBooksPayload
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("ListBooks - Bind Error", err)
		return response.AsError().
			WithMessage("Invalid Payload").
			MakeJSON(c, http.StatusBadRequest)
	}

	books, err := s.booksRepo.ListBooks(c.Request().Context(), &payload)
	if err != nil {
		s.log.Errorln("ListBooks - List Book Repository Error", err)
		return response.AsError().
			WithMessage("Internal Server Error").
			MakeJSON(c, http.StatusInternalServerError)
	}

	if len(books) == 0 {
		s.log.Errorln("ListBooks - Books Not Found")
		return response.AsError().
			WithMessage("Books Not Found").
			MakeJSON(c, http.StatusNotFound)
	}

	s.log.Println("ListBooks - Success")
	return response.WithMessage("List Books Success").
		WithData(books).
		MakeJSON(c, http.StatusOK)
}

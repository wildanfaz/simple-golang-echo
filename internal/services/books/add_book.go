package books

import (
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/simple-golang-echo/internal/helpers"
	"github.com/wildanfaz/simple-golang-echo/internal/models"
)

func (s *Service) AddBook(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.Book
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("AddBook - Bind Error", err)
		return response.AsError().
			WithMessage("Invalid Payload").
			MakeJSON(c, http.StatusBadRequest)
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.Title, validation.Required),
		validation.Field(&payload.Description, validation.Required),
		validation.Field(&payload.Author, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z ]+$")).
			Error("only accept alpha or whitespace characters")),
	)
	if err != nil {
		s.log.Errorln("AddBook - Validate Struct Error", err)
		return response.AsError().
			WithMessage(err.Error()).
			MakeJSON(c, http.StatusBadRequest)
	}

	err = s.booksRepo.AddBook(c.Request().Context(), &payload)
	if err != nil {
		s.log.Errorln("AddBook - Add Book Repository Error", err)
		return response.AsError().
			WithMessage("Internal Server Error").
			MakeJSON(c, http.StatusInternalServerError)
	}

	s.log.Println("AddBook - Success")
	return response.WithMessage("Add Book Success").
		MakeJSON(c, http.StatusOK)
}

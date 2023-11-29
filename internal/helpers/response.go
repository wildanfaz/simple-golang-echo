package helpers

import "github.com/labstack/echo/v4"

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	defaultMessage := "Unknown Message"

	return &Response{
		Message: defaultMessage,
	}
}

func (r *Response) AsError() *Response {
	r.Error = true

	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data

	return r
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message

	return r
}

func (r *Response) MakeJSON(c echo.Context, code int) error {
	return c.JSON(code, r)
}

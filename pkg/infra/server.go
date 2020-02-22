package infra

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/mpppk/cli-template/pkg/infra/handler"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func registerHandlers(e *echo.Echo) {
	h := handler.New()
	e.GET("/api/sum", h.Sum)
}

func NewServer() *echo.Echo {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	registerHandlers(e)
	return e
}

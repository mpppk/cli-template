package handler

import (
	"github.com/comail/colog"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func registerHandlers(e *echo.Echo) {
	h := New()
	e.GET("/api/sum", h.Sum)
}

// NewServer create new echo server with handlers
func NewServer() *echo.Echo {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	registerHandlers(e)
	return e
}

// InitializeLog initialize log settings
func InitializeLog(verbose bool) {
	colog.Register()
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LInfo)

	if verbose {
		colog.SetMinLevel(colog.LDebug)
	}
}

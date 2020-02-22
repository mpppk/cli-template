package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mpppk/cli-template/pkg/usecase"
	"github.com/mpppk/cli-template/pkg/util"
)

type handler struct{}

type sumRequest struct {
	A    int  `query:"a" Validate:"required"`
	B    int  `query:"b" Validate:"required"`
	Norm bool `query:"norm"`
}

func New() *handler {
	return &handler{}
}

func (h *handler) Sum(c echo.Context) error {
	req := new(sumRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ToErrorResponse(err))
	}

	var result int
	if req.Norm {
		result = usecase.CalcL1Norm([]int{req.A, req.B})
	} else {
		result = usecase.CalcSum([]int{req.A, req.B})
	}
	return c.JSON(http.StatusOK, result)
}

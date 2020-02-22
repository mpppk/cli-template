package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mpppk/cli-template/pkg/usecase"
)

type Handler struct{}

type sumRequest struct {
	A    int  `query:"a" Validate:"required"`
	B    int  `query:"b" Validate:"required"`
	Norm bool `query:"norm"`
}

type sumResponse struct {
	Result int `json:"result"`
}

// New create new handlers
func New() *Handler {
	return &Handler{}
}

// Sum handle http request to calculate sum
func (h *Handler) Sum(c echo.Context) error {
	req := new(sumRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		logWithJSON("invalid request", req)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var result int
	if req.Norm {
		result = usecase.CalcL1Norm([]int{req.A, req.B})
	} else {
		result = usecase.CalcSum([]int{req.A, req.B})
	}
	return c.JSON(http.StatusOK, sumResponse{Result: result})
}

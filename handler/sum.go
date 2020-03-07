package handler

import (
	"net/http"

	"github.com/mpppk/cli-template/domain/repository"

	"github.com/labstack/echo"
	"github.com/mpppk/cli-template/usecase"
)

// Handlers represent handlers of echo server
type Handlers struct {
	sumHistoryRepository repository.SumHistory
}

type sumRequest struct {
	A    int  `query:"a" Validate:"required"`
	B    int  `query:"b" Validate:"required"`
	Norm bool `query:"norm"`
}

type sumResponse struct {
	Result int `json:"result"`
}

// New create new handlers
func New(sumHistoryRepository repository.SumHistory) *Handlers {
	return &Handlers{
		sumHistoryRepository: sumHistoryRepository,
	}
}

// Sum handle http request to calculate sum
func (h *Handlers) Sum(c echo.Context) error {
	req := new(sumRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		logWithJSON("invalid request", req)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	useCase := usecase.NewSum(h.sumHistoryRepository)

	var result int
	if req.Norm {
		result = useCase.CalcL1Norm([]int{req.A, req.B})
	} else {
		result = useCase.CalcSum([]int{req.A, req.B})
	}
	return c.JSON(http.StatusOK, sumResponse{Result: result})
}

package handler

import (
	"net/http"

	"github.com/mpppk/cli-template/domain/model"

	"github.com/labstack/echo"
	"github.com/mpppk/cli-template/usecase"
)

type sumHistoryRequest struct {
	Limit int `query:"limit" Validate:"required"`
}

type sumHistoryResponse struct {
	Result []*model.SumHistory `json:"result"`
}

// SumHistory handle http request to list sum history
func (h *Handlers) SumHistory(c echo.Context) error {
	req := new(sumHistoryRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		logWithJSON("invalid request", req)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	useCase := usecase.NewSum(h.sumHistoryRepository)

	history := useCase.ListSumHistory(req.Limit)

	return c.JSON(http.StatusOK, sumHistoryResponse{Result: history})
}

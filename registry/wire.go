//+build wireinject

package registry

//go:generate wire

import (
	"github.com/labstack/echo"
	"github.com/mpppk/cli-template/domain/model"
	"github.com/mpppk/cli-template/handler"
	"github.com/mpppk/cli-template/infra"
	"github.com/mpppk/cli-template/infra/repoimpl"
	"github.com/mpppk/cli-template/usecase"
)
import "github.com/google/wire"

func InitializeHandler(v []*model.SumHistory) *handler.Handlers {
	wire.Build(handler.New, usecase.NewSum, repoimpl.NewMemorySumHistory)
	return &handler.Handlers{}
}

func InitializeSumUseCase(v []*model.SumHistory) *usecase.Sum {
	wire.Build(repoimpl.NewMemorySumHistory, usecase.NewSum)
	return &usecase.Sum{}
}

func InitializeServer(v []*model.SumHistory) *echo.Echo {
	wire.Build(
		handler.New,
		repoimpl.NewMemorySumHistory,
		usecase.NewSum,
		infra.NewServer,
	)
	return &echo.Echo{}
}

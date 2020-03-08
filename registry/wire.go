//+build wireinject

package registry

import (
	"github.com/labstack/echo"
	"github.com/mpppk/cli-template/handler"
	"github.com/mpppk/cli-template/infra"
	"github.com/mpppk/cli-template/infra/repoimpl"
	"github.com/mpppk/cli-template/usecase"
)
import "github.com/google/wire"

func InitializeHandler() *handler.Handlers {
	wire.Build(handler.New, usecase.NewSum, repoimpl.NewMemorySumHistory)
	return &handler.Handlers{}
}

func InitializeSumUseCase() *usecase.Sum {
	wire.Build(repoimpl.NewMemorySumHistory, usecase.NewSum)
	return &usecase.Sum{}
}

func InitializeServer() *echo.Echo {
	wire.Build(
		handler.New,
		repoimpl.NewMemorySumHistory,
		usecase.NewSum,
		infra.NewServer,
	)
	return &echo.Echo{}
}

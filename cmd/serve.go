package cmd

import (
	"net/http"

	"github.com/mpppk/cli-template/pkg/util"

	"github.com/go-playground/validator/v10"

	"github.com/mpppk/cli-template/pkg/usecase"

	"github.com/labstack/echo"

	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

type sumRequest struct {
	A    int  `query:"a" Validate:"required"`
	B    int  `query:"b" Validate:"required"`
	Norm bool `query:"norm"`
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func sumHandler(c echo.Context) error {
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

func newServeCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run server",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			e := echo.New()
			e.Validator = &customValidator{validator: validator.New()}
			e.GET("/api/sum", sumHandler)
			e.Logger.Fatal(e.Start(":1323"))
			return nil
		},
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newServeCmd)
}

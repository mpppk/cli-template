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

type SumRequest struct {
	A    int  `query:"a" validate:"required"`
	B    int  `query:"b" validate:"required"`
	Norm bool `query:"norm"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func newServeCmd(fs afero.Fs) (*cobra.Command, error) {
	sumHandler := func(c echo.Context) error {
		req := new(SumRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, util.ToErrorResponse(err))
		}

		var result int
		if req.Norm {
			r, err := usecase.CalcL1Norm([]int{req.A, req.B})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			result = r
		} else {
			r, err := usecase.CalcSum([]int{req.A, req.B})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			result = r
		}
		return c.JSON(http.StatusOK, result)
	}

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run server",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}
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

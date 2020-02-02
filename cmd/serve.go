package cmd

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

func newServeCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run server",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			e := echo.New()
			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, "Hello, World!")
			})
			e.Logger.Fatal(e.Start(":1323"))
			return nil
		},
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newServeCmd)
}

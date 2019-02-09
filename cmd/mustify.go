package cmd

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"strings"

	goofyast "github.com/mpppk/goofy/ast"
	"github.com/spf13/cobra"
)

var fileName *string
var mustifyCmd = &cobra.Command{
	Use:   "mustify",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		prog, err := goofyast.NewProgram(*fileName, nil)
		if err != nil {
			panic(err)
		}
		for _, pkg := range prog.Created {
			for _, file := range pkg.Files {
				ast.FileExports(file)

				for i, decl := range file.Decls {
					funcDecl, ok := decl.(*ast.FuncDecl)
					if !ok {
						continue
					}
					newD, ok := goofyast.ConvertErrorFuncToMustFunc(prog, funcDecl)
					funcNameRunes := []rune(newD.Name.Name)
					newD.Name.Name = "Must" + strings.ToUpper(string(funcNameRunes[0])) + string(funcNameRunes[1:])
					if !ok {
						continue
					}
					file.Decls[i] = newD
				}
				if err := format.Node(os.Stdout, token.NewFileSet(), file); err != nil {
					panic(err)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(mustifyCmd)
	fileName = mustifyCmd.Flags().String("file", "", "target file path")
}

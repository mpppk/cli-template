package cmd

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"

	goofyast "github.com/mpppk/goofy/ast"
	"github.com/spf13/cobra"
)

var fileName *string
var outFilePath *string
var mustifyCmd = &cobra.Command{
	Use:   "mustify",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		prog, err := goofyast.NewProgram(*fileName)
		if err != nil {
			panic(err)
		}

		for _, pkg := range prog.Created {
			for i, file := range pkg.Files {
				currentFileName := prog.Fset.File(pkg.Files[i].Pos()).Name()
				if currentFileName != *fileName {
					continue
				}

				var newDecls []ast.Decl
				for _, decl := range file.Decls {
					if genDecl, ok := decl.(*ast.GenDecl); ok {
						if genDecl.Tok == token.IMPORT {
							newDecls = append(newDecls, genDecl)
						}
						continue
					}

					funcDecl, ok := decl.(*ast.FuncDecl)
					if !ok {
						panic("unknown decl in " + currentFileName)
					}

					if !ast.IsExported(funcDecl.Name.Name) {
						continue
					}

					newDecl, ok := goofyast.ConvertErrorFuncToMustFunc(prog, pkg, funcDecl)
					if !ok {
						continue
					}
					newDecls = append(newDecls, newDecl)
				}
				file.Decls = newDecls

				f, err := os.Create(*outFilePath)
				if err != nil {
					panic(err)
				}

				defer func() {
					if err := f.Close(); err != nil {
						panic(err)
					}
				}()
				if err := format.Node(f, token.NewFileSet(), file); err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mustifyCmd)
	fileName = mustifyCmd.Flags().String("file", "", "target file path")
	outFilePath = mustifyCmd.Flags().String("out", "must-"+*fileName, "file path to save output")
}

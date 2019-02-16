package cmd

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"path/filepath"

	goofyast "github.com/mpppk/goofy/ast"
	"github.com/spf13/cobra"
)

var filePath *string
var outFilePath *string
var mustifyCmd = &cobra.Command{
	Use:   "mustify",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		absFilePath, err := filepath.Abs(*filePath)
		if err != nil {
			panic(err)
		}

		if *outFilePath == "" {
			base := filepath.Base(*filePath)
			o := filepath.Join(filepath.Dir(*filePath), "must-"+base)
			outFilePath = &o
		}

		absOutFilePath, err := filepath.Abs(*outFilePath)
		if err != nil {
			panic(err)
		}

		prog, err := goofyast.NewProgram(*filePath)
		if err != nil {
			panic(err)
		}

		for _, pkg := range prog.Created {
			for i, file := range pkg.Files {
				currentFilePath := prog.Fset.File(pkg.Files[i].Pos()).Name()
				absCurrentFilePath, err := filepath.Abs(currentFilePath)
				if err != nil {
					panic(err)
				}

				if absFilePath != absCurrentFilePath {
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
						panic("unknown decl in " + currentFilePath)
					}

					if !ast.IsExported(funcDecl.Name.Name) {
						continue
					}

					newDecl, ok := goofyast.GenerateErrorFuncWrapper(pkg, funcDecl)
					//newDecl, ok := goofyast.ConvertErrorFuncToMustFunc(prog, pkg, funcDecl)
					if !ok {
						continue
					}
					newDecls = append(newDecls, newDecl)
				}
				file.Decls = newDecls

				f, err := os.Create(absOutFilePath)
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
	filePath = mustifyCmd.Flags().String("file", "", "target file path")
	outFilePath = mustifyCmd.Flags().String("out", "", "file path to save output")
}

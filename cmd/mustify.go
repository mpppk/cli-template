package cmd

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"

	goofyast "github.com/mpppk/goofy/ast"
	"github.com/spf13/cobra"
)

var fileName *string
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

				newDeclMap := map[string]ast.Decl{}
				var newDecls []ast.Decl
				for _, decl := range file.Decls {
					if genDecl, ok := decl.(*ast.GenDecl); ok {
						if genDecl.Tok == token.IMPORT {
							newDecls = append(newDecls, genDecl)
						}
						continue
					}

					funcDecl, ok := decl.(*ast.FuncDecl)
					if !ast.IsExported(funcDecl.Name.Name) {
						continue
					}

					fmt.Printf("func: %v, file: %v\n", funcDecl.Name.Name, currentFileName)
					newDecl, ok := goofyast.ConvertErrorFuncToMustFunc(prog, pkg, funcDecl)
					if !ok {
						continue
					}
					newDeclMap[newDecl.Name.Name] = newDecl
				}

				for _, d := range newDeclMap {
					newDecls = append(newDecls, d)
				}
				file.Decls = newDecls

				f, err := os.Create(fmt.Sprintf("must-%v", currentFileName))
				if err != nil {
					panic(err)
					// FIXME Openエラー処理
				}
				defer f.Close() // FIXME error handling
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
}

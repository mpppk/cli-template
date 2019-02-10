package cmd

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"strconv"

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

		var newDecls []ast.Decl
		for _, pkg := range prog.Created {
			for _, file := range pkg.Files {
				ast.FileExports(file)

				for _, decl := range file.Decls {
					funcDecl, ok := decl.(*ast.FuncDecl)
					if !ok {
						continue
					}

					newDecl, ok := goofyast.ConvertErrorFuncToMustFunc(prog, pkg, funcDecl)
					if !ok {
						continue
					}
					newDecls = append(newDecls, newDecl)
				}
			}
		}

		file, err := os.Create("must-hoge.go")
		if err != nil {
			// Openエラー処理
		}
		defer file.Close()
		newFile := NewFileAst("utl", NewImportSpecs([]string{"fmt"}), newDecls)
		if err := format.Node(file, token.NewFileSet(), newFile); err != nil {
			panic(err)
		}

	},
}

func NewImportSpecs(importPaths []string) (importSpecs []*ast.ImportSpec) {
	for _, importPath := range importPaths {
		importSpec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(importPath),
			},
		}
		importSpecs = append(importSpecs, importSpec)
	}
	return
}

func NewFileAst(packageName string, imports []*ast.ImportSpec, funcs []ast.Decl) *ast.File {
	var importSpecs []ast.Spec
	for _, im := range imports {
		importSpecs = append(importSpecs, im)
	}

	decls := []ast.Decl{
		&ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: importSpecs,
		},
	}

	decls = append(decls, funcs...)

	return &ast.File{
		Name:  ast.NewIdent(packageName),
		Decls: decls,
	}
}

func init() {
	rootCmd.AddCommand(mustifyCmd)
	fileName = mustifyCmd.Flags().String("file", "", "target file path")
}

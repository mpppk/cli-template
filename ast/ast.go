package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"

	"github.com/go-toolsmith/astcopy"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/loader"
)

func NewProgram(fileName string) (*loader.Program, error) {
	lo := loader.Config{
		Fset:       token.NewFileSet(),
		ParserMode: parser.ParseComments}
	dirPath := filepath.Dir(fileName)
	packages, err := parser.ParseDir(lo.Fset, dirPath, nil, parser.ParseComments)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dir: "+dirPath)
	}

	var files []*ast.File
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			files = append(files, file)
		}
	}

	lo.CreateFromFiles("main", files...)
	return lo.Load()
}

func ConvertErrorFuncToMustFunc(prog *loader.Program, currentPkg *loader.PackageInfo, orgFuncDecl *ast.FuncDecl) (*ast.FuncDecl, bool) {
	funcDecl := astcopy.FuncDecl(orgFuncDecl)
	if !IsErrorFunc(funcDecl) {
		return nil, false
	}
	results := funcDecl.Type.Results.List
	funcDecl.Type.Results.List = results[:len(results)-1]
	replaceReturnStmtByPanicIfErrorExist(prog, currentPkg, funcDecl)
	addPrefixToFunc(funcDecl, "Must")
	return funcDecl, true
}

func replaceReturnStmtByPanicIfErrorExist(orgProg *loader.Program, currentPkg *loader.PackageInfo, funcDecl *ast.FuncDecl) *ast.FuncDecl {
	newFuncDeclNode := astutil.Apply(funcDecl, func(cr *astutil.Cursor) bool {
		returnStmt, ok := cr.Node().(*ast.ReturnStmt)
		if !ok {
			return true
		}

		returnResults := returnStmt.Results
		if len(returnResults) == 0 {
			return true
		}

		lastReturnResult := returnResults[len(returnResults)-1]
		returnStmt.Results = returnResults[:len(returnResults)-1]
		if lastReturnResultIdent, ok := lastReturnResult.(*ast.Ident); ok {
			if lastReturnResultIdent.Name == "nil" {
				return true
			}

			panicIfErrExistIfStmt := generatePanicIfErrorExistStmtAst(lastReturnResultIdent.Name)
			cr.InsertBefore(panicIfErrExistIfStmt)
			return true
		}

		if lastReturnResultCallExpr, ok := lastReturnResult.(*ast.CallExpr); ok {
			typeNames, err := getCallExprReturnTypes(orgProg, currentPkg, lastReturnResultCallExpr)
			if err != nil {
				panic(err)
			}
			var lhs []string
			for range typeNames {
				lhs = append(lhs, "_")
			}

			tempErrValueName := getAvailableValueName(currentPkg.Pkg, "err", lastReturnResultCallExpr.Pos())
			lhs[len(lhs)-1] = tempErrValueName
			assignStmt := generateAssignStmt(lhs, lastReturnResultCallExpr)
			panicIfErrExistIfStmt := generatePanicIfErrorExistStmtAst(tempErrValueName)
			cr.InsertBefore(assignStmt)
			cr.InsertBefore(panicIfErrExistIfStmt)
		}

		return true
	}, nil)
	newFuncDecl := newFuncDeclNode.(*ast.FuncDecl)
	return newFuncDecl
}

func getAvailableValueName(currentPkg *types.Package, valName string, pos token.Pos) string {
	innerMost := currentPkg.Scope().Innermost(pos)
	s, _ := innerMost.LookupParent(valName, pos)
	if s == nil {
		return valName
	}

	cnt := 0
	valNameWithNumber := fmt.Sprintf("%v%v", valName, cnt)
	for {
		s, _ := innerMost.LookupParent(valNameWithNumber, pos)
		if s == nil {
			return valNameWithNumber
		}
		cnt++
		valNameWithNumber = fmt.Sprintf("%v%v", valName, cnt)
	}
}

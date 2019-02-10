package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/loader"
)

func NewProgram(fileName string, source interface{}) (*loader.Program, error) {
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

func IsErrorFunc(funcDecl *ast.FuncDecl) bool {
	lastResultIdent, ok := extractFuncLastResultIdent(funcDecl)
	if !ok {
		return false
	}
	return lastResultIdent.Name == "error"
}

func ConvertErrorFuncToMustFunc(prog *loader.Program, currentPkg *loader.PackageInfo, funcDecl *ast.FuncDecl) (*ast.FuncDecl, bool) {
	if !IsErrorFunc(funcDecl) {
		return nil, false
	}
	results := funcDecl.Type.Results.List
	funcDecl.Type.Results.List = results[:len(results)-1]
	replaceReturnStmtByPanicIfErrorExist(prog, currentPkg, funcDecl)
	addPrefixToFunc(funcDecl, "Must")
	return funcDecl, true
}

func addPrefixToFunc(funcDecl *ast.FuncDecl, prefix string) {
	funcNameRunes := []rune(funcDecl.Name.Name)
	funcDecl.Name.Name = prefix + strings.ToUpper(string(funcNameRunes[0])) + string(funcNameRunes[1:])
}

func replaceReturnStmtByPanicIfErrorExist(prog *loader.Program, currentPkg *loader.PackageInfo, funcDecl *ast.FuncDecl) *ast.FuncDecl {
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
			typeNames, err := getCallExprReturnTypes(prog, currentPkg, lastReturnResultCallExpr)
			if err != nil {
				panic(err)
			}
			var lhs []string
			for range typeNames {
				lhs = append(lhs, "_")
			}

			tempErrValueName := getAvailableValueName(currentPkg.Pkg, "err", lastReturnResultCallExpr.Pos())
			lhs[len(lhs)-1] = tempErrValueName // FIXME スコープ内に同名の変数があれば適当に変えるorDEFINEじゃなくて代入にする
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

func extractFuncLastResultIdent(funcDecl *ast.FuncDecl) (*ast.Ident, bool) {
	expr, ok := extractFuncLastResultExpr(funcDecl)
	if !ok {
		return nil, false
	}
	ident, ok := expr.(*ast.Ident)
	return ident, ok
}

func extractFuncLastResultExpr(funcDecl *ast.FuncDecl) (ast.Expr, bool) {
	if funcDecl == nil {
		panic(fmt.Sprintf("funcDecl is nil"))
	}
	if funcDecl.Type == nil {
		panic(fmt.Sprintf("funcDecl.Type is nil: %v", funcDecl.Name))
	}
	if funcDecl.Type.Results == nil {
		return nil, false
	}
	results := funcDecl.Type.Results.List
	if len(results) == 0 {
		return nil, false
	}
	return results[len(results)-1].Type, true
}

func getCallExprReturnTypes(prog *loader.Program, currentPkg *loader.PackageInfo, callExpr *ast.CallExpr) ([]string, error) {
	pkg := currentPkg
	var funcName string
	switch fun := callExpr.Fun.(type) {
	case *ast.SelectorExpr:
		packageIdent, ok := fun.X.(*ast.Ident)
		if !ok {
			return nil, errors.New("selectorExpr.X is not *ast.Ident")
		}
		pkg = prog.Package(packageIdent.Name)
		funcName = fun.Sel.Name
	case *ast.Ident:
		funcName = fun.Name
	case nil:
		panic("callExpr is nil")
	}
	fmt.Println("current pkg: " + currentPkg.Pkg.Name() + " target pkg: " + pkg.Pkg.Name())
	typeNames, ok := getFuncDeclResultTypes(pkg, funcName)
	if !ok {
		return nil, errors.New("func not found: " + funcName)
	}

	return typeNames, nil
}

func getFuncDeclResultTypes(packageInfo *loader.PackageInfo, funcName string) (types []string, ok bool) {
	funcDecl, ok := getFuncDecl(packageInfo, funcName)
	if !ok {
		fmt.Println("failed to get func decl: " + funcName)
		return nil, false
	}

	if funcDecl == nil {
		panic(fmt.Sprintf("funcDecl is nil"))
	}
	if funcDecl.Type == nil {
		panic(fmt.Sprintf("funcDecl.Type is nil: %v", funcDecl.Name))
	}
	if funcDecl.Type.Results == nil {
		panic(fmt.Sprintf("funcDecl.Type.Results is nil: %#v", funcDecl.Type))
	}
	if funcDecl.Type.Results.List == nil {
		panic("funcDecl.Type.Results.List is nil")
	}
	results := funcDecl.Type.Results.List
	for _, result := range results {
		if typeIdent, ok := result.Type.(*ast.Ident); ok {
			types = append(types, typeIdent.Name)
		}
	}
	return types, true
}

func getFuncDecl(packageInfo *loader.PackageInfo, funcName string) (*ast.FuncDecl, bool) {
	if packageInfo == nil {
		panic("packageInfo is nil")
	}
	if packageInfo.Files == nil {
		panic("packageInfo.Files is nil")
	}
	for _, file := range packageInfo.Files {
		ast.FileExports(file)
		for _, decl := range file.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				if funcDecl.Name.Name == funcName {
					return funcDecl, true
				}
			}
		}
	}
	return nil, false
}

func generateAssignStmt(lhNames []string, callExpr *ast.CallExpr) *ast.AssignStmt {
	var lhs []ast.Expr
	for _, lhName := range lhNames {
		lh := &ast.Ident{
			Name: lhName,
		}
		lhs = append(lhs, lh)
	}

	return &ast.AssignStmt{
		Lhs: lhs,
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			callExpr,
		},
	}
}

func generatePanicIfErrorExistStmtAst(errValName string) *ast.IfStmt {
	// generatePanicIfErrorExistStmtAst return ast of below code
	// if errValName != nil {
	//     panic(errValName)
	// }
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X: &ast.Ident{
				Name: errValName,
			},
			Op: token.NEQ,
			Y: &ast.Ident{
				Name: "nil",
				Obj:  nil,
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.Ident{
							Name: "panic",
						},
						Args: []ast.Expr{
							&ast.Ident{
								Name: errValName,
							},
						},
					},
				},
			},
		},
		Else: nil,
	}
}

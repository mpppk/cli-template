package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/loader"
)

func NewProgram(fileName string, source interface{}) (*loader.Program, error) {
	lo := loader.Config{ParserMode: parser.ParseComments}
	astf, err := lo.ParseFile(fileName, source)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse file from "+fileName)
	}
	lo.CreateFromFiles("main", astf)
	return lo.Load()
}

func IsErrorFunc(funcDecl *ast.FuncDecl) bool {
	lastResultIdent, ok := extractFuncLastResultIdent(funcDecl)
	if !ok {
		return false
	}
	return lastResultIdent.Name == "error"
}

func ConvertErrorFuncToMustFunc(prog *loader.Program, funcDecl *ast.FuncDecl) (*ast.FuncDecl, bool) {
	if !IsErrorFunc(funcDecl) {
		return nil, false
	}
	results := funcDecl.Type.Results.List
	funcDecl.Type.Results.List = results[:len(results)-1]
	replaceReturnStmtByPanicIfErrorExist(prog, funcDecl)
	addPrefixToFunc(funcDecl, "Must")
	return funcDecl, true
}

func addPrefixToFunc(funcDecl *ast.FuncDecl, prefix string) {
	funcNameRunes := []rune(funcDecl.Name.Name)
	funcDecl.Name.Name = prefix + strings.ToUpper(string(funcNameRunes[0])) + string(funcNameRunes[1:])
}

func replaceReturnStmtByPanicIfErrorExist(prog *loader.Program, funcDecl *ast.FuncDecl) *ast.FuncDecl {
	newFuncDeclNode := astutil.Apply(funcDecl, func(cr *astutil.Cursor) bool {
		returnStmt, ok := cr.Node().(*ast.ReturnStmt)
		if !ok {
			return true
		}

		// TODO: returnのvalueを省略した時どうなる?
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
			types, err := getCallExprReturnTypes(prog, lastReturnResultCallExpr)
			if err != nil {
				panic(err)
			}
			var lhs []string
			for range types {
				lhs = append(lhs, "_")
			}

			tempErrValueName := "e"            // FIXME
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

func extractFuncLastResultIdent(funcDecl *ast.FuncDecl) (*ast.Ident, bool) {
	expr, ok := extractFuncLastResultExpr(funcDecl)
	if !ok {
		return nil, false
	}
	ident, ok := expr.(*ast.Ident)
	return ident, ok
}

func extractFuncLastResultExpr(funcDecl *ast.FuncDecl) (ast.Expr, bool) {
	results := funcDecl.Type.Results.List
	if len(results) == 0 {
		return nil, false
	}
	return results[len(results)-1].Type, true
}

func getCallExprReturnTypes(prog *loader.Program, callExpr *ast.CallExpr) ([]string, error) {
	selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, errors.New("lastReturnResultCallExpr.Fun is not *ast.SelectorExpr")
	}

	xIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return nil, errors.New("selectorExpr.X is not *ast.Ident")
	}

	pkg := prog.Package(xIdent.Name)
	types, ok := getFuncDeclResultTypes(pkg, selectorExpr.Sel.Name)
	if !ok {
		return nil, errors.New("func not found: " + xIdent.Name)
	}

	return types, nil
}

func getFuncDeclResultTypes(packageInfo *loader.PackageInfo, funcName string) (types []string, ok bool) {
	funcDecl, ok := getFuncDecl(packageInfo, funcName)
	if !ok {
		return nil, false
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

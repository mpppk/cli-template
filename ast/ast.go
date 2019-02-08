package ast

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/loader"
)

func NewProgram(fileName string, source interface{}) (*loader.Program, error) {
	loader := loader.Config{ParserMode: parser.ParseComments}
	astf, err := loader.ParseFile(fileName, source)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse file from "+fileName)
	}
	loader.CreateFromFiles("main", astf)
	return loader.Load()
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
	// 最後の戻り値を削除
	results := funcDecl.Type.Results.List
	funcDecl.Type.Results.List = results[:len(results)-1]
	replaceReturnStmtByPanicIfErrorExist(prog, funcDecl)
	return funcDecl, true
}

func replaceReturnStmtByPanicIfErrorExist(prog *loader.Program, funcDecl *ast.FuncDecl) *ast.FuncDecl {
	// return行を書き換え
	// errorがnilなら単に削除
	// errorが存在すれば、代わりにpanic
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
			// TODO: returnの最後が関数の場合は、関数の戻り値をerrorだけ一旦受けて、
			// それがnilでなければpanicする

			selectorExpr, ok := lastReturnResultCallExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				panic("lastReturnresultCallExpr.Fun is not *ast.SelectorExpr")
			}
			_, _ = selectorExpr.X.(*ast.Ident)
			xIdent, ok := selectorExpr.X.(*ast.Ident)
			if !ok {
				panic("selectorExpr.X is not *ast.Ident")
			}

			pkg := prog.Package(xIdent.Name)
			types, ok := getFuncDeclResultTypes(pkg, selectorExpr.Sel.Name)
			if !ok {
				panic("func not found: " + xIdent.Name)
			}

			var lhs []string
			for range types {
				lhs = append(lhs, "_")
			}

			tempErrValueName := "goofyTempErrValue" // FIXME
			lhs[len(lhs)-1] = tempErrValueName      // FIXME スコープ内に同名の変数があれば適当に変えるorDEFINEじゃなくて代入にする
			assignStmt := generateAssignStmt(lhs, lastReturnResultCallExpr)
			panicIfErrExistIfStmt := generatePanicIfErrorExistStmtAst(tempErrValueName)
			cr.InsertBefore(assignStmt)
			cr.InsertBefore(panicIfErrExistIfStmt)

			// TODO: errorがnilでないので、returnStmtを削除し、代わりにpanicする
			//returnStmt.Results = returnResults[:len(returnResults)-1]
			//
			//if len(types) <= 1 {
			//	return true
			//}
			//
			//for _, newType := range types[:len(types)-1] {
			//	returnStmt.Results = append(returnStmt.Results, &ast.Ident{
			//		Name: newType,
			//	})
			//}
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

func generatePanicAst(args []ast.Expr) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{
				Name: "panic",
			},
			Args: args,
		},
	}
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

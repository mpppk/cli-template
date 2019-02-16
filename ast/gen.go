package ast

import (
	"go/ast"
	"go/token"
)

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

func generateCallExpr(funcName string, argNames []string) *ast.CallExpr {
	var args []ast.Expr
	for _, argName := range argNames {
		args = append(args, &ast.Ident{
			Name: argName,
		})
	}

	return &ast.CallExpr{
		Fun: &ast.Ident{
			Name: funcName,
		},
		Args: args,
	}
}

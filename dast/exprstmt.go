package dast

import (
	"io"
)

type ExprStmt struct {
	nodeBase
	Expr Node
}

func NewExprStmt(pos PositionHolder, expr Node) *ExprStmt {
	n := &ExprStmt{Expr: expr}
	n.init(pos)
	return n
}

func (n *ExprStmt) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	dumpNode(n.Expr, w, nest+1)
}

package dast

import (
	"io"
)

type MinusExpr struct {
	nodeBase
	Expr Node
}

func NewMinusExpr(pos PositionHolder, expr Node) *MinusExpr {
	n := &MinusExpr{Expr: expr}
	n.init(pos)
	return n
}

func (n *MinusExpr) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	dumpNode(n.Expr, w, nest+1)
}

package dast

import (
	"io"
)

type Return struct {
	nodeBase
	Expr Node
}

func NewReturn(pos PositionHolder, expr Node) *Return {
	n := &Return{Expr: expr}
	n.init(pos)
	return n
}

func (n *Return) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	dumpNode(n.Expr, w, nest+1)
}

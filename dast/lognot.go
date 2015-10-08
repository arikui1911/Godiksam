package dast

import (
	"io"
)

type LogNot struct {
	nodeBase
	Expr Node
}

func NewLogNot(pos PositionHolder, expr Node) *LogNot {
	n := &LogNot{Expr: expr}
	n.init(pos)
	return n
}

func (n *LogNot) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	dumpNode(n.Expr, w, nest+1)
}

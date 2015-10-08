package dast

import (
	"io"
)

type Throw struct {
	nodeBase
	Expr Node
}

func NewThrow(pos PositionHolder, expr Node) *Throw {
	n := &Throw{Expr: expr}
	n.init(pos)
	return n
}

func (n *Throw) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	dumpNode(n.Expr, w, nest+1)
}

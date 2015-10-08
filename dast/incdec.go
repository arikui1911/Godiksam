package dast

import (
	"fmt"
	"io"
)

type IncDec struct {
	nodeBase
	Op   Operator
	Expr Node
}

func NewIncDec(pos PositionHolder, expr Node, op Operator) *IncDec {
	n := &IncDec{Op: op, Expr: expr}
	n.init(pos)
	return n
}

func (n *IncDec) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Op)
	dumpNode(n.Expr, w, nest+1)
}

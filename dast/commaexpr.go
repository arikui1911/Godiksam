package dast

import (
	"io"
)

type CommaExpr struct {
	nodeBase
	Left  Node
	Right Node
}

func NewCommaExpr(pos PositionHolder, left Node, right Node) *CommaExpr {
	n := &CommaExpr{Left: left, Right: right}
	n.init(pos)
	return n
}

func (n *CommaExpr) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	tag("Left", w, nest+1)
	dumpNode(n.Left, w, nest+2)
	tag("Right", w, nest+1)
	dumpNode(n.Right, w, nest+2)
}

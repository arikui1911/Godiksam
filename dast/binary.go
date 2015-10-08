package dast

import (
	"fmt"
	"io"
)

type Binary struct {
	nodeBase
	Op    Operator
	Left  Node
	Right Node
}

func NewBinary(pos PositionHolder, op Operator, left Node, right Node) *Binary {
	n := &Binary{Op: op, Left: left, Right: right}
	n.init(pos)
	return n
}

func (n *Binary) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Op)
	tag("Left", w, nest+1)
	dumpNode(n.Left, w, nest+2)
	tag("Right", w, nest+1)
	dumpNode(n.Right, w, nest+2)
}

package dast

import (
	"fmt"
	"io"
)

type Assign struct {
	nodeBase
	Op    Operator
	Left  Node
	Right Node
}

func NewAssign(pos PositionHolder, op Operator, left Node, right Node) *Assign {
	n := &Assign{Op: op, Left: left, Right: right}
	n.init(pos)
	return n
}

func (n *Assign) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Op)
	tag("Left", w, nest+1)
	dumpNode(n.Left, w, nest+2)
	tag("Right", w, nest+1)
	dumpNode(n.Right, w, nest+2)
}

package dast

import (
	"fmt"
	"io"
)

type While struct {
	nodeBase
	Label string
	Cond  Node
	Body  Node
}

func NewWhile(pos PositionHolder, label string, cond Node, body Node) *While {
	n := &While{Label: label, Cond: cond, Body: body}
	n.init(pos)
	return n
}

func (n *While) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Label)
	tag("Cond", w, nest+1)
	dumpNode(n.Cond, w, nest+2)
	tag("Body", w, nest+1)
	dumpNode(n.Body, w, nest+2)
}

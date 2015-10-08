package dast

import (
	"io"
)

type If struct {
	nodeBase
	Test Node
	Body Node
	Alt  Node
}

func NewIf(pos PositionHolder, test Node, body Node, alt Node) *If {
	n := &If{Test: test, Body: body, Alt: alt}
	n.init(pos)
	return n
}

func (n *If) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	tag("Test", w, nest+1)
	dumpNode(n.Test, w, nest+2)
	tag("Body", w, nest+1)
	dumpNode(n.Body, w, nest+2)
	tag("Alt", w, nest+1)
	dumpNode(n.Alt, w, nest+2)
}

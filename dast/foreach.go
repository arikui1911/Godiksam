package dast

import (
	"fmt"
	"io"
)

type Foreach struct {
	nodeBase
	Label      string
	Var        string
	Collection Node
	Body       Node
}

func NewForeach(pos PositionHolder, label string, aVar string, collection Node, body Node) *Foreach {
	n := &Foreach{Label: label, Var: aVar, Collection: collection, Body: body}
	n.init(pos)
	return n
}

func (n *Foreach) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Label)
	tag(fmt.Sprintf("Collection: %s", n.Var), w, nest+1)
	dumpNode(n.Body, w, nest+2)
	tag("Body", w, nest+1)
	dumpNode(n.Body, w, nest+2)
}

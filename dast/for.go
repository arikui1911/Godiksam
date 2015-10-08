package dast

import (
	"fmt"
	"io"
)

type For struct {
	nodeBase
	Label string
	Init  Node
	Cond  Node
	Post  Node
	Body  Node
}

func NewFor(pos PositionHolder, label string, init Node, cond Node, post Node, body Node) *For {
	n := &For{Label: label, Init: init, Cond: cond, Post: post, Body: body}
	n.init(pos)
	return n
}

func (n *For) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Label)
	tag("Init", w, nest+1)
	dumpNode(n.Init, w, nest+2)
	tag("Cond", w, nest+1)
	dumpNode(n.Cond, w, nest+2)
	tag("Post", w, nest+1)
	dumpNode(n.Post, w, nest+2)
	tag("Body", w, nest+1)
	dumpNode(n.Body, w, nest+2)
}

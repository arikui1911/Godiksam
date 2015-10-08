package dast

import (
	"io"
)

type Funcall struct {
	nodeBase
	Func Node
	Args []Node
}

func NewFuncall(pos PositionHolder, fun Node, args []Node) *Funcall {
	n := &Funcall{Func: fun, Args: args}
	n.init(pos)
	return n
}

func (n *Funcall) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	tag("Func", w, nest+1)
	dumpNode(n.Func, w, nest+2)
	tag("Args", w, nest+1)
	dumpNodes(n.Args, w, nest+2)
}

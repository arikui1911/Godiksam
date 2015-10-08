package dast

import (
	"io"
)

type Decl struct {
	nodeBase
}

func NewDecl(pos PositionHolder, ts TypeSpec, name string, init Node) *Decl {
	n := &Decl{}
	n.init(pos)
	return n
}

func (n *Decl) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
}

package dast

import (
	"fmt"
	"io"
)

type IdentExpr struct {
	nodeBase
	Name string
}

func NewIdentExpr(pos PositionHolder, name string) *IdentExpr {
	n := &IdentExpr{Name: name}
	n.init(pos)
	return n
}

func (n *IdentExpr) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Name)
}

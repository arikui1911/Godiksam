package dast

import (
	"fmt"
	"io"
)

type IntLiteral struct {
	nodeBase
	Value int
}

func NewIntLiteral(pos PositionHolder, value int) *IntLiteral {
	n := &IntLiteral{Value: value}
	n.init(pos)
	return n
}

func (n *IntLiteral) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Value)
}

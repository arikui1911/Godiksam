package dast

import (
	"fmt"
	"io"
)

type BooleanLiteral struct {
	nodeBase
	Value bool
}

func NewBooleanLiteral(pos PositionHolder, value bool) *BooleanLiteral {
	n := &BooleanLiteral{Value: value}
	n.init(pos)
	return n
}

func (n *BooleanLiteral) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Value)
}

package dast

import (
	"fmt"
	"io"
)

type StrLiteral struct {
	nodeBase
	Value string
}

func NewStrLiteral(pos PositionHolder, value string) *StrLiteral {
	n := &StrLiteral{Value: value}
	n.init(pos)
	return n
}

func (n *StrLiteral) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Value)
}

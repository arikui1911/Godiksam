package dast

import (
	"fmt"
	"io"
)

type DoubleLiteral struct {
	nodeBase
	Value float64
}

func NewDoubleLiteral(pos PositionHolder, value float64) *DoubleLiteral {
	n := &DoubleLiteral{Value: value}
	n.init(pos)
	return n
}

func (n *DoubleLiteral) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Value)
}

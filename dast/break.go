package dast

import (
	"fmt"
	"io"
)

type Break struct {
	nodeBase
	Label string
}

func NewBreak(pos PositionHolder, label string) *Break {
	n := &Break{Label: label}
	n.init(pos)
	return n
}

func (n *Break) dump(w io.Writer, nest int) {
	header(n, w, nest, false)
	fmt.Fprintf(w, "%v\n", n.Label)
}

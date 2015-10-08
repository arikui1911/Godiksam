package dast

import (
	//	"fmt"
	"io"
)

type Param struct {
	nodeBase
}

func NewParam(pos PositionHolder, ts TypeSpec, name string) *Param {
	n := &Param{}
	n.init(pos)
	return n
}

func (n *Param) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
}

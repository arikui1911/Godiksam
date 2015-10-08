package dast

import (
	"io"
)

type Block struct {
	nodeBase
	children []Node
}

func NewBlock(pos PositionHolder) *Block {
	b := &Block{children: make([]Node, 0)}
	b.init(pos)
	return b
}

func (b *Block) Add(child Node) {
	b.children = append(b.children, child)
}

func (b *Block) Append(children []Node) {
	b.children = append(b.children, children...)
}

func (b *Block) dump(w io.Writer, nest int) {
	header(b, w, nest, true)
	dumpNodes(b.children, w, nest+1)
}

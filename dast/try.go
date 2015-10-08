package dast

import (
	"fmt"
	"io"
)

type Try struct {
	nodeBase
	Try     Node
	Catcher string
	Catch   Node
	Finally Node
}

func NewTry(pos PositionHolder, try Node, catcher string, catch Node, finally Node) *Try {
	n := &Try{Try: try, Catcher: catcher, Catch: catch, Finally: finally}
	n.init(pos)
	return n
}

func (n *Try) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	tag("Try", w, nest+1)
	dumpNode(n.Try, w, nest+2)
	tag(fmt.Sprintf("Catch: %s", n.Catcher), w, nest+1)
	dumpNode(n.Catch, w, nest+2)
	tag("Finally", w, nest+1)
	dumpNode(n.Finally, w, nest+2)
}

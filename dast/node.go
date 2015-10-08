package dast

import (
	"fmt"
	"io"
)

type Node interface {
	Line() int
	Column() int
	dump(io.Writer, int)
}

type nodeBase struct {
	line int
	col  int
}

type PositionHolder interface {
	Line() int
	Column() int
}

type TypeSpec interface{}

func (n *nodeBase) init(pos PositionHolder) {
	n.line = pos.Line()
	n.col = pos.Column()
}

func (n *nodeBase) Line() int   { return n.line }
func (n *nodeBase) Column() int { return n.col }

func Dump(tree Node, output io.Writer) {
	dumpNode(tree, output, 0)
}

func indent(w io.Writer, nest int) {
	if nest < 1 {
		return
	}
	for i := 0; i < nest; i++ {
		fmt.Fprint(w, "    ")
	}
}

func header(n Node, w io.Writer, nest int, requireNewline bool) {
	indent(w, nest)
	fmt.Fprintf(w, "%d:%d:%T: ", n.Line(), n.Column(), n)
	if requireNewline {
		fmt.Fprint(w, "\n")
	}
}

func tag(name string, w io.Writer, nest int) {
	indent(w, nest)
	fmt.Fprintf(w, "@%s:\n", name)
}

func dumpNode(n Node, w io.Writer, nest int) {
	if n == nil {
		indent(w, nest)
		fmt.Fprintln(w, "(nil)")
		return
	}
	n.dump(w, nest)
}

func dumpNodes(nodes []Node, w io.Writer, nest int) {
	if nodes == nil {
		indent(w, nest)
		fmt.Fprintln(w, "(nil)")
		return
	}
	for _, n := range nodes {
		dumpNode(n, w, nest)
	}
}

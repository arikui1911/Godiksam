package dkc

import (
	"github.com/arikui1911/Godiksam/dast"
)

type declaration struct {
	name    string
	aType   *TypeSpec
	init    dast.Node
	varIdx  int
	isLocal bool
}

func newDeclaration(ts *TypeSpec, name string, init dast.Node) *declaration {
	return &declaration{
		name:   name,
		aType:  ts,
		init:   init,
		varIdx: -1,
	}
}

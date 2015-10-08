package t

import (
	"fmt"
	"github.com/arikui1911/Godiksam/dast"
	"github.com/arikui1911/Godiksam/dkc"
	"os"
	"strings"
)

func ExampleParser() {
	src := `
if (123) {
  456;
}
`

	p := dkc.NewParser(strings.NewReader(src), "(test)", 1)
	err := p.Parse()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	dast.Dump(p.Tree, os.Stdout)

	fmt.Println("Hello, world!")

	// Output:
	// hogeeeeeee
}

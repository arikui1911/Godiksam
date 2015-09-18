package t

import (
	"fmt"
	"godiksam/dast"
	"godiksam/dkc"
	"os"
	"strings"
)

func ExampleParser() {
	src := `
123;
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

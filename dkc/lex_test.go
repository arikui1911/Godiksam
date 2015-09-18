package dkc

import (
	"fmt"
	"strings"
)

func ExampleParser() {
	src := `
// hogeee
hoge

/* foo
bar
baz
*/

true

123 45.6
0xff
0 0.8e+2

"foo\nbar
baz\t!!!"

(666)++

"hogeee"

0x

`
	p := NewParser(strings.NewReader(src), "(test)", 1)
	lval := new(yySymType)

	for p.Lex(lval) != 0 {
		fmt.Printf("%v\n", lval.token)
	}

	fmt.Println("Hello!")

	// Output:
	// fooooooo!!!
}

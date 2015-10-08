package dast

type Operator int

const (
	ADD Operator = iota
	SUB
	MUL
	DIV
	MOD

	EQ
	NE
	GT
	LT
	GE
	LE
	LOG_OR
	LOG_AND

	INC
	DEC

	NORMAL_ASSIGN
	ADD_ASSIGN
	SUB_ASSIGN
	MUL_ASSIGN
	DIV_ASSIGN
	MOD_ASSIGN
)

var operatorStrings = []string{
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"EQ",
	"NE",
	"GT",
	"LT",
	"GE",
	"LE",
	"LOG_OR",
	"LOG_AND",
	"INC",
	"DEC",
	"NORMAL_ASSIGN",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"MOD_ASSIGN",
}

func (op Operator) String() string {
	return operatorStrings[op-ADD]
}

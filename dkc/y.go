//line diksam.go.y:1
package dkc

import __yyfmt__ "fmt"

//line diksam.go.y:3
import (
	"fmt"
	"godiksam/dast"
)

func tokName(c int) string {
	switch c {
	case 0:
		return "EOF"
	case '\n':
		return "\\n"
	}
	i := c - kIF
	if i < 0 || i > len(yyToknames)-1 {
		return fmt.Sprintf("%c", c)
	}
	return yyToknames[i]
}

func addTopLevelStmt(l yyLexer, stmt dast.Node) {
	p := l.(*Parser)
	if p.Tree == nil {
		p.Tree = dast.NewBlock(stmt.Line(), stmt.Column())
	}
	p.Tree.(*dast.Block).Add(stmt)
}

func defun(l yyLexer, pos *Token, typeSpec baseType, name *Token, params []dast.Node, body dast.Node) {
}

func openBlock(l yyLexer, pos *Token) dast.Node {
	return nil
}

func closeBlock(l yyLexer, block dast.Node, stmts []dast.Node) dast.Node {
	return nil
}

//line diksam.go.y:45
type yySymType struct {
	yys      int
	token    *Token
	node     dast.Node
	nodes    []dast.Node
	typeSpec baseType
}

const kIF = 57346
const kELSE = 57347
const kELSIF = 57348
const kSWITCH = 57349
const kCASE = 57350
const kDEFAULT = 57351
const kWHILE = 57352
const kDO = 57353
const kFOR = 57354
const kFOREACH = 57355
const kRETURN = 57356
const kBREAK = 57357
const kCONTINUE = 57358
const kNULL = 57359
const kTRUE = 57360
const kFALSE = 57361
const kTRY = 57362
const kCATCH = 57363
const kFINALLY = 57364
const kTHROW = 57365
const kTHROWS = 57366
const kBOOLEAN = 57367
const kVOID = 57368
const kINT = 57369
const kDOUBLE = 57370
const kSTRING = 57371
const kNATIVE_POINTER = 57372
const kNEW = 57373
const kREQUIRE = 57374
const kRENAME = 57375
const kCLASS = 57376
const kINTERFACE = 57377
const kPUBLIC = 57378
const kPRIVATE = 57379
const kVIRTUAL = 57380
const kOVERRIDE = 57381
const kABSTRACT = 57382
const kTHIS = 57383
const kSUPER = 57384
const kCONSTRUCTOR = 57385
const kINSTANCEOF = 57386
const kDELEGATE = 57387
const kENUM = 57388
const kFINAL = 57389
const kCONST = 57390
const tEQ = 57391
const tNE = 57392
const tGE = 57393
const tLE = 57394
const tADD_A = 57395
const tSUB_A = 57396
const tMUL_A = 57397
const tDIV_A = 57398
const tMOD_A = 57399
const tINC = 57400
const tDEC = 57401
const tDCAST_BEG = 57402
const tDCAST_END = 57403
const tLOG_AND = 57404
const tLOG_OR = 57405
const tINT_LITERAL = 57406
const tDOUBLE_LITERAL = 57407
const tSTRING_LITERAL = 57408
const tREGEXP_LITERAL = 57409
const tIDENT = 57410

var yyToknames = []string{
	"kIF",
	"kELSE",
	"kELSIF",
	"kSWITCH",
	"kCASE",
	"kDEFAULT",
	"kWHILE",
	"kDO",
	"kFOR",
	"kFOREACH",
	"kRETURN",
	"kBREAK",
	"kCONTINUE",
	"kNULL",
	"kTRUE",
	"kFALSE",
	"kTRY",
	"kCATCH",
	"kFINALLY",
	"kTHROW",
	"kTHROWS",
	"kBOOLEAN",
	"kVOID",
	"kINT",
	"kDOUBLE",
	"kSTRING",
	"kNATIVE_POINTER",
	"kNEW",
	"kREQUIRE",
	"kRENAME",
	"kCLASS",
	"kINTERFACE",
	"kPUBLIC",
	"kPRIVATE",
	"kVIRTUAL",
	"kOVERRIDE",
	"kABSTRACT",
	"kTHIS",
	"kSUPER",
	"kCONSTRUCTOR",
	"kINSTANCEOF",
	"kDELEGATE",
	"kENUM",
	"kFINAL",
	"kCONST",
	"tEQ",
	"tNE",
	"tGE",
	"tLE",
	"tADD_A",
	"tSUB_A",
	"tMUL_A",
	"tDIV_A",
	"tMOD_A",
	"tINC",
	"tDEC",
	"tDCAST_BEG",
	"tDCAST_END",
	"tLOG_AND",
	"tLOG_OR",
	"tINT_LITERAL",
	"tDOUBLE_LITERAL",
	"tSTRING_LITERAL",
	"tREGEXP_LITERAL",
	"tIDENT",
	" {",
	" ;",
	" ,",
	" =",
	" <",
	" >",
	" +",
	" -",
	" *",
	" /",
	" %",
	" !",
	" (",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line diksam.go.y:386

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 37,
}

const yyNprod = 88
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 288

var yyAct = []int{

	51, 153, 18, 6, 5, 45, 4, 147, 55, 7,
	171, 40, 162, 46, 131, 40, 40, 54, 89, 9,
	10, 11, 170, 29, 30, 12, 164, 159, 13, 69,
	14, 40, 15, 16, 17, 139, 137, 40, 88, 21,
	87, 40, 145, 85, 41, 86, 138, 136, 127, 58,
	59, 60, 61, 62, 64, 65, 40, 98, 33, 97,
	101, 102, 103, 104, 105, 106, 109, 113, 57, 26,
	27, 28, 83, 19, 84, 80, 80, 63, 74, 75,
	14, 35, 15, 16, 17, 36, 25, 125, 126, 163,
	32, 128, 46, 146, 129, 132, 80, 64, 65, 143,
	40, 98, 92, 134, 72, 73, 80, 80, 80, 53,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 83,
	63, 84, 76, 77, 78, 140, 70, 71, 144, 91,
	82, 99, 40, 118, 119, 148, 90, 124, 160, 150,
	149, 39, 40, 142, 152, 135, 156, 53, 151, 31,
	46, 158, 157, 53, 141, 161, 130, 49, 38, 24,
	165, 114, 115, 116, 117, 168, 166, 46, 7, 167,
	172, 173, 174, 29, 30, 175, 56, 66, 9, 10,
	11, 22, 29, 30, 12, 67, 68, 13, 7, 14,
	169, 15, 16, 17, 93, 94, 48, 3, 9, 10,
	11, 1, 29, 30, 12, 8, 34, 13, 50, 14,
	108, 15, 16, 17, 154, 155, 123, 111, 112, 26,
	27, 28, 96, 47, 29, 30, 110, 23, 26, 27,
	28, 35, 19, 20, 52, 36, 25, 107, 100, 2,
	35, 37, 79, 81, 36, 25, 0, 133, 26, 27,
	28, 14, 19, 15, 16, 17, 42, 0, 43, 44,
	35, 0, 0, 0, 36, 25, 0, 95, 0, 0,
	26, 27, 28, 0, 47, 0, 0, 0, 0, 0,
	0, 0, 35, 120, 121, 122, 36, 25,
}
var yyPact = []int{

	5, 5, -1000, -1000, -1000, 90, 71, -37, 246, 206,
	89, 89, 40, 206, -1000, -1000, -1000, -1000, -1000, -76,
	113, -4, 115, -1000, 136, 206, -1000, -1000, -1000, -1000,
	-1000, 53, 3, 45, -1000, 206, 206, -1000, 49, -1000,
	206, 206, -41, -43, -63, 66, -30, -1000, 59, -1000,
	32, 173, 184, -1000, 61, -1000, 206, 206, 206, 206,
	206, 206, 206, 155, -1000, -1000, 206, 206, 206, -15,
	206, 206, 206, 206, 206, 206, 206, 206, 206, -1000,
	39, -1000, 55, -1000, 206, -1000, -34, 206, 206, 88,
	-1000, -1000, -1000, -67, 40, -1000, 164, -1000, 77, -1000,
	115, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -35, -1000,
	136, 53, 53, -1000, 3, 3, 3, 3, 45, 45,
	-1000, -1000, -1000, -36, 84, 75, 29, 40, -40, 23,
	-77, 206, -1000, -1000, -1000, 2, -1000, 206, 78, 226,
	-1000, -1000, -1000, -1000, 209, 40, 206, 206, -55, -1000,
	-1000, -1000, 70, -1000, 40, -69, -1000, 19, -56, 40,
	-1000, -1000, 206, 206, 40, 168, -60, -72, -1000, 40,
	40, 40, -1000, 209, -1000, -1000,
}
var yyPgo = []int{

	0, 6, 234, 0, 1, 5, 3, 2, 233, 181,
	159, 149, 90, 58, 206, 39, 227, 222, 216, 210,
	205, 196, 4, 201, 239, 197,
}
var yyR1 = []int{

	0, 23, 23, 24, 24, 22, 22, 22, 22, 25,
	25, 25, 25, 18, 18, 2, 3, 3, 17, 17,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 4, 4, 4, 20, 20, 5,
	5, 21, 21, 6, 6, 7, 7, 7, 7, 7,
	7, 7, 8, 8, 9, 9, 10, 10, 10, 11,
	11, 11, 11, 11, 12, 12, 12, 13, 13, 13,
	13, 14, 14, 14, 15, 15, 15, 15, 15, 16,
	16, 16, 16, 16, 16, 16, 19, 19,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 6,
	5, 6, 5, 2, 4, 1, 2, 3, 1, 2,
	2, 3, 5, 6, 6, 10, 8, 3, 3, 3,
	7, 4, 9, 3, 0, 2, 6, 0, 2, 0,
	1, 0, 1, 1, 3, 1, 3, 3, 3, 3,
	3, 3, 1, 3, 1, 3, 1, 3, 3, 1,
	3, 3, 3, 3, 1, 3, 3, 1, 3, 3,
	3, 1, 2, 2, 1, 3, 4, 2, 2, 3,
	1, 1, 1, 1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -23, -24, -25, -1, -22, -6, 4, -20, 14,
	15, 16, 20, 23, 25, 27, 28, 29, -7, 68,
	-8, -15, -9, -16, -10, 81, 64, 65, 66, 18,
	19, -11, -12, -13, -14, 76, 80, -24, 68, 70,
	71, 81, 10, 12, 13, -5, -6, 68, -21, 68,
	-21, -3, -2, 69, -6, 84, 63, 72, 53, 54,
	55, 56, 57, 81, 58, 59, 62, 49, 50, -6,
	73, 74, 51, 52, 75, 76, 77, 78, 79, -14,
	-15, -14, 81, 70, 72, -7, -6, 81, 81, 81,
	70, 70, 70, 21, 22, 83, -17, -1, -22, 70,
	-9, -7, -7, -7, -7, -7, -7, 82, -19, -7,
	-10, -11, -11, 82, -12, -12, -12, -12, -13, -13,
	-14, -14, -14, -18, 82, -22, -6, 82, -6, -5,
	68, 81, -3, 83, -1, 68, 82, 71, 82, 71,
	-3, 70, 68, 70, -3, 82, 70, 84, -6, -7,
	-3, 70, -22, -4, 5, 6, -3, -5, -6, 82,
	68, -3, 81, 70, 82, -3, -6, -5, -3, 22,
	82, 82, -3, -3, -3, -4,
}
var yyDef = []int{

	37, -2, 1, 3, 4, 0, 0, 0, 0, 39,
	41, 41, 0, 0, 5, 6, 7, 8, 43, 80,
	45, 71, 52, 74, 54, 0, 81, 82, 83, 84,
	85, 56, 59, 64, 67, 0, 0, 2, 0, 20,
	0, 0, 0, 0, 0, 0, 40, 80, 0, 42,
	0, 0, 37, 15, 0, 38, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 78, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	71, 73, 0, 21, 0, 44, 0, 0, 39, 0,
	27, 28, 29, 0, 0, 16, 37, 18, 0, 33,
	53, 46, 47, 48, 49, 50, 51, 75, 0, 86,
	55, 57, 58, 79, 60, 61, 62, 63, 65, 66,
	68, 69, 70, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 31, 17, 19, 0, 76, 0, 0, 0,
	10, 12, 13, 22, 34, 0, 39, 0, 0, 87,
	9, 11, 0, 23, 0, 0, 24, 0, 0, 0,
	14, 35, 0, 39, 0, 30, 0, 0, 26, 0,
	0, 0, 32, 34, 25, 36,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 80, 3, 3, 3, 79, 3, 3,
	81, 82, 77, 75, 71, 76, 3, 78, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 84, 70,
	73, 72, 74, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 69, 3, 83,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 4:
		//line diksam.go.y:77
		{
			addTopLevelStmt(yylex, yyS[yypt-0].node)
		}
	case 5:
		//line diksam.go.y:81
		{
			yyVAL.typeSpec = ttBOOLEAN
		}
	case 6:
		//line diksam.go.y:82
		{
			yyVAL.typeSpec = ttINT
		}
	case 7:
		//line diksam.go.y:83
		{
			yyVAL.typeSpec = ttDOUBLE
		}
	case 8:
		//line diksam.go.y:84
		{
			yyVAL.typeSpec = ttSTRING
		}
	case 9:
		//line diksam.go.y:87
		{
			defun(yylex, yyS[yypt-4].token, yyS[yypt-5].typeSpec, yyS[yypt-4].token, yyS[yypt-2].nodes, yyS[yypt-0].node)
		}
	case 10:
		//line diksam.go.y:91
		{
			defun(yylex, yyS[yypt-3].token, yyS[yypt-4].typeSpec, yyS[yypt-3].token, nil, yyS[yypt-0].node)
		}
	case 11:
		//line diksam.go.y:95
		{
			defun(yylex, yyS[yypt-4].token, yyS[yypt-5].typeSpec, yyS[yypt-4].token, yyS[yypt-2].nodes, nil)
		}
	case 12:
		//line diksam.go.y:99
		{
			defun(yylex, yyS[yypt-3].token, yyS[yypt-4].typeSpec, yyS[yypt-3].token, nil, nil)
		}
	case 13:
		//line diksam.go.y:104
		{
			yyVAL.nodes = append(make([]dast.Node, 0), dast.NewParam(yyS[yypt-0].token, yyS[yypt-1].typeSpec, yyS[yypt-0].token))
		}
	case 14:
		//line diksam.go.y:108
		{
			yyVAL.nodes = append(yyS[yypt-3].nodes, dast.NewParam(yyS[yypt-0].token, yyS[yypt-1].typeSpec, yyS[yypt-0].token))
		}
	case 15:
		//line diksam.go.y:113
		{
			yyVAL.node = openBlock(yylex, yyS[yypt-0].token)
		}
	case 16:
		//line diksam.go.y:118
		{
			yyVAL.node = closeBlock(yylex, yyS[yypt-1].node, nil)
		}
	case 17:
		//line diksam.go.y:122
		{
			yyVAL.node = closeBlock(yylex, yyS[yypt-2].node, yyS[yypt-1].nodes)
		}
	case 18:
		//line diksam.go.y:127
		{
			yyVAL.nodes = append(make([]dast.Node, 0), yyS[yypt-0].node)
		}
	case 19:
		//line diksam.go.y:131
		{
			yyVAL.nodes = append(yyS[yypt-1].nodes, yyS[yypt-0].node)
		}
	case 20:
		//line diksam.go.y:136
		{
			yyVAL.node = dast.NewExprStmt(yyS[yypt-0].token, yyS[yypt-1].node)
		}
	case 21:
		//line diksam.go.y:140
		{
			yyVAL.node = dast.NewDecl(yyS[yypt-1].token, yyS[yypt-2].typeSpec, yyS[yypt-1].token, nil)
		}
	case 22:
		//line diksam.go.y:144
		{
			yyVAL.node = dast.NewDecl(yyS[yypt-3].token, yyS[yypt-4].typeSpec, yyS[yypt-3].token, yyS[yypt-1].node)
		}
	case 23:
		//line diksam.go.y:148
		{
			yyVAL.node = dast.NewIf(yyS[yypt-5].token, yyS[yypt-3].node, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 24:
		//line diksam.go.y:152
		{
			yyVAL.node = dast.NewWhile(yyS[yypt-4].token, yyS[yypt-5].token, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 25:
		//line diksam.go.y:156
		{
			yyVAL.node = dast.NewFor(yyS[yypt-8].token, yyS[yypt-9].token, yyS[yypt-6].node, yyS[yypt-4].node, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 26:
		//line diksam.go.y:160
		{
			yyVAL.node = dast.NewForeach(yyS[yypt-6].token, yyS[yypt-7].token, yyS[yypt-4].token, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 27:
		//line diksam.go.y:164
		{
			yyVAL.node = dast.NewReturn(yyS[yypt-2].token, yyS[yypt-1].node)
		}
	case 28:
		//line diksam.go.y:168
		{
			yyVAL.node = dast.NewBreak(yyS[yypt-2].token, yyS[yypt-1].token)
		}
	case 29:
		//line diksam.go.y:172
		{
			yyVAL.node = dast.NewContinue(yyS[yypt-2].token, yyS[yypt-1].token)
		}
	case 30:
		//line diksam.go.y:176
		{
			yyVAL.node = dast.NewTry(yyS[yypt-6].token, yyS[yypt-5].node, yyS[yypt-2].node, yyS[yypt-0].node, nil)
		}
	case 31:
		//line diksam.go.y:180
		{
			yyVAL.node = dast.NewTry(yyS[yypt-3].token, yyS[yypt-2].node, nil, nil, yyS[yypt-0].node)
		}
	case 32:
		//line diksam.go.y:184
		{
			yyVAL.node = dast.NewTry(yyS[yypt-8].token, yyS[yypt-7].node, yyS[yypt-4].node, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 33:
		//line diksam.go.y:188
		{
			yyVAL.node = dast.NewThrow(yyS[yypt-2].token, yyS[yypt-1].node)
		}
	case 34:
		//line diksam.go.y:193
		{
			yyVAL.node = nil
		}
	case 35:
		//line diksam.go.y:197
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 36:
		//line diksam.go.y:201
		{
			yyVAL.node = dast.NewIf(yyS[yypt-5].token, yyS[yypt-3].node, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 37:
		//line diksam.go.y:206
		{
			yyVAL.token = nil
		}
	case 38:
		yyVAL.token = yyS[yypt-0].token
	case 39:
		//line diksam.go.y:212
		{
			yyVAL.node = nil
		}
	case 40:
		yyVAL.node = yyS[yypt-0].node
	case 41:
		//line diksam.go.y:218
		{
			yyVAL.token = nil
		}
	case 42:
		yyVAL.token = yyS[yypt-0].token
	case 43:
		yyVAL.node = yyS[yypt-0].node
	case 44:
		//line diksam.go.y:225
		{
			yyVAL.node = dast.NewCommaExpr(yyS[yypt-1].token, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 45:
		yyVAL.node = yyS[yypt-0].node
	case 46:
		//line diksam.go.y:231
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.NORMAL_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 47:
		//line diksam.go.y:235
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.ADD_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 48:
		//line diksam.go.y:239
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.SUB_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 49:
		//line diksam.go.y:243
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.MUL_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 50:
		//line diksam.go.y:247
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.DIV_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 51:
		//line diksam.go.y:251
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.MOD_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 52:
		yyVAL.node = yyS[yypt-0].node
	case 53:
		//line diksam.go.y:257
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LOG_OR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 54:
		yyVAL.node = yyS[yypt-0].node
	case 55:
		//line diksam.go.y:263
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LOG_AND, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 56:
		yyVAL.node = yyS[yypt-0].node
	case 57:
		//line diksam.go.y:269
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.EQ, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 58:
		//line diksam.go.y:273
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.NE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 59:
		yyVAL.node = yyS[yypt-0].node
	case 60:
		//line diksam.go.y:279
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.GT, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 61:
		//line diksam.go.y:283
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LT, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 62:
		//line diksam.go.y:287
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.GE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 63:
		//line diksam.go.y:291
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 64:
		yyVAL.node = yyS[yypt-0].node
	case 65:
		//line diksam.go.y:297
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.ADD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 66:
		//line diksam.go.y:301
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.SUB, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 67:
		yyVAL.node = yyS[yypt-0].node
	case 68:
		//line diksam.go.y:307
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.MUL, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 69:
		//line diksam.go.y:311
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.DIV, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 70:
		//line diksam.go.y:315
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.MOD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 71:
		yyVAL.node = yyS[yypt-0].node
	case 72:
		//line diksam.go.y:321
		{
			yyVAL.node = dast.NewMinusExpr(yyS[yypt-1].token, yyS[yypt-0].node)
		}
	case 73:
		//line diksam.go.y:325
		{
			yyVAL.node = dast.NewLogNot(yyS[yypt-1].token, yyS[yypt-0].node)
		}
	case 74:
		yyVAL.node = yyS[yypt-0].node
	case 75:
		//line diksam.go.y:331
		{
			yyVAL.node = dast.NewFuncall(yyS[yypt-1].token, yyS[yypt-2].node, nil)
		}
	case 76:
		//line diksam.go.y:335
		{
			yyVAL.node = dast.NewFuncall(yyS[yypt-2].token, yyS[yypt-3].node, yyS[yypt-1].nodes)
		}
	case 77:
		//line diksam.go.y:339
		{
			yyVAL.node = dast.NewIncDec(yyS[yypt-0].token, yyS[yypt-1].node, +1)
		}
	case 78:
		//line diksam.go.y:343
		{
			yyVAL.node = dast.NewIncDec(yyS[yypt-0].token, yyS[yypt-1].node, -1)
		}
	case 79:
		//line diksam.go.y:348
		{
			yyVAL.node = yyS[yypt-1].node
		}
	case 80:
		//line diksam.go.y:352
		{
			yyVAL.node = dast.NewIdentExpr(yyS[yypt-0].token, yyS[yypt-0].token.Value.(string))
		}
	case 81:
		//line diksam.go.y:356
		{
			yyVAL.node = dast.NewIntExpr(yyS[yypt-0].token, yyS[yypt-0].token.Value.(int))
		}
	case 82:
		//line diksam.go.y:360
		{
			yyVAL.node = dast.NewDoubleExpr(yyS[yypt-0].token, yyS[yypt-0].token.Value.(float64))
		}
	case 83:
		//line diksam.go.y:364
		{
			yyVAL.node = dast.NewStrExpr(yyS[yypt-0].token, yyS[yypt-0].token.Value.(string))
		}
	case 84:
		//line diksam.go.y:368
		{
			yyVAL.node = dast.NewBooleanExpr(yyS[yypt-0].token, true)
		}
	case 85:
		//line diksam.go.y:372
		{
			yyVAL.node = dast.NewBooleanExpr(yyS[yypt-0].token, false)
		}
	case 86:
		//line diksam.go.y:377
		{
			yyVAL.nodes = append(make([]dast.Node, 0), yyS[yypt-0].node)
		}
	case 87:
		//line diksam.go.y:381
		{
			yyVAL.nodes = append(yyS[yypt-2].nodes, yyS[yypt-0].node)
		}
	}
	goto yystack /* stack new state and value */
}

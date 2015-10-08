//line diksam.go.y:1
package dkc

import __yyfmt__ "fmt"

//line diksam.go.y:3
import (
	"fmt"
	"github.com/arikui1911/Godiksam/dast"
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
		p.Tree = dast.NewBlock(stmt)
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
	ident    string
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

//line diksam.go.y:394

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

const yyLast = 281

var yyAct = []int{

	51, 153, 18, 6, 5, 45, 4, 147, 55, 171,
	159, 64, 65, 46, 40, 40, 40, 54, 58, 59,
	60, 61, 62, 64, 65, 170, 164, 145, 14, 69,
	15, 16, 17, 139, 63, 162, 131, 57, 137, 89,
	88, 40, 40, 85, 138, 86, 63, 87, 33, 136,
	29, 30, 127, 113, 76, 77, 78, 98, 7, 97,
	101, 102, 103, 104, 105, 106, 109, 41, 9, 10,
	11, 163, 29, 30, 12, 74, 75, 13, 83, 14,
	84, 15, 16, 17, 83, 124, 84, 125, 126, 82,
	32, 128, 46, 40, 129, 132, 26, 27, 28, 146,
	47, 98, 53, 134, 72, 73, 143, 40, 35, 99,
	40, 92, 36, 25, 107, 39, 40, 34, 26, 27,
	28, 91, 19, 118, 119, 140, 70, 71, 144, 90,
	35, 53, 151, 160, 36, 25, 148, 133, 142, 150,
	149, 53, 141, 135, 152, 130, 156, 31, 49, 38,
	46, 158, 157, 79, 81, 161, 56, 24, 66, 169,
	165, 114, 115, 116, 117, 168, 166, 46, 7, 167,
	172, 173, 174, 29, 30, 175, 22, 3, 9, 10,
	11, 1, 29, 30, 12, 67, 68, 13, 7, 14,
	8, 15, 16, 17, 120, 121, 122, 108, 9, 10,
	11, 21, 29, 30, 12, 154, 155, 13, 123, 14,
	48, 15, 16, 17, 96, 111, 112, 93, 94, 26,
	27, 28, 50, 47, 110, 2, 23, 37, 26, 27,
	28, 35, 19, 100, 20, 36, 25, 80, 80, 14,
	35, 15, 16, 17, 36, 25, 52, 95, 26, 27,
	28, 42, 19, 43, 44, 0, 0, 0, 80, 0,
	35, 0, 0, 0, 36, 25, 0, 0, 80, 80,
	80, 0, 80, 80, 80, 80, 80, 80, 80, 80,
	80,
}
var yyPact = []int{

	184, 184, -1000, -1000, -1000, 81, 45, -14, 241, 155,
	80, 80, 33, 155, -1000, -1000, -1000, -1000, -1000, -76,
	93, -35, 96, -1000, 136, 155, -1000, -1000, -1000, -1000,
	-1000, 53, 0, -23, -1000, 155, 155, -1000, 8, -1000,
	155, 155, -34, -41, -42, 59, 22, -1000, 51, -1000,
	41, 196, 164, -1000, 39, -1000, 155, 155, 155, 155,
	155, 155, 155, 32, -1000, -1000, 155, 155, 155, -29,
	155, 155, 155, 155, 155, 155, 155, 155, 155, -1000,
	-47, -1000, 3, -1000, 155, -1000, -30, 155, 155, 77,
	-1000, -1000, -1000, -45, 33, -1000, 54, -1000, 75, -1000,
	96, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -33, -1000,
	136, 53, 53, -1000, 0, 0, 0, 0, -23, -23,
	-1000, -1000, -1000, -38, 72, 70, 36, 33, -55, 29,
	-77, 68, -1000, -1000, -1000, 14, -1000, 155, 62, 214,
	-1000, -1000, -1000, -1000, 200, 33, 155, 155, -72, -1000,
	-1000, -1000, 65, -1000, 33, -46, -1000, 1, -56, 33,
	-1000, -1000, 155, 155, 33, 137, -57, -73, -1000, 33,
	33, 33, -1000, 200, -1000, -1000,
}
var yyPgo = []int{

	0, 6, 246, 0, 1, 5, 3, 2, 234, 176,
	157, 147, 90, 48, 117, 201, 226, 214, 208, 197,
	190, 210, 4, 181, 225, 177,
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
	-3, 70, 68, 70, -3, 82, 70, 84, 68, -7,
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
		//line diksam.go.y:79
		{
			addTopLevelStmt(yylex, yyS[yypt-0].node)
		}
	case 5:
		//line diksam.go.y:83
		{
			yyVAL.typeSpec = ttBOOLEAN
		}
	case 6:
		//line diksam.go.y:84
		{
			yyVAL.typeSpec = ttINT
		}
	case 7:
		//line diksam.go.y:85
		{
			yyVAL.typeSpec = ttDOUBLE
		}
	case 8:
		//line diksam.go.y:86
		{
			yyVAL.typeSpec = ttSTRING
		}
	case 9:
		//line diksam.go.y:89
		{
			defun(yylex, yyS[yypt-4].token, yyS[yypt-5].typeSpec, yyS[yypt-4].token, yyS[yypt-2].nodes, yyS[yypt-0].node)
		}
	case 10:
		//line diksam.go.y:93
		{
			defun(yylex, yyS[yypt-3].token, yyS[yypt-4].typeSpec, yyS[yypt-3].token, nil, yyS[yypt-0].node)
		}
	case 11:
		//line diksam.go.y:97
		{
			defun(yylex, yyS[yypt-4].token, yyS[yypt-5].typeSpec, yyS[yypt-4].token, yyS[yypt-2].nodes, nil)
		}
	case 12:
		//line diksam.go.y:101
		{
			defun(yylex, yyS[yypt-3].token, yyS[yypt-4].typeSpec, yyS[yypt-3].token, nil, nil)
		}
	case 13:
		//line diksam.go.y:106
		{
			yyVAL.nodes = append(make([]dast.Node, 0), dast.NewParam(yyS[yypt-0].token, newTypeSpec(yyS[yypt-1].typeSpec), yyS[yypt-0].token.Value.(string)))
		}
	case 14:
		//line diksam.go.y:110
		{
			yyVAL.nodes = append(yyS[yypt-3].nodes, dast.NewParam(yyS[yypt-0].token, newTypeSpec(yyS[yypt-1].typeSpec), yyS[yypt-0].token.Value.(string)))
		}
	case 15:
		//line diksam.go.y:115
		{
			yyVAL.node = openBlock(yylex, yyS[yypt-0].token)
		}
	case 16:
		//line diksam.go.y:120
		{
			yyVAL.node = closeBlock(yylex, yyS[yypt-1].node, nil)
		}
	case 17:
		//line diksam.go.y:124
		{
			yyVAL.node = closeBlock(yylex, yyS[yypt-2].node, yyS[yypt-1].nodes)
		}
	case 18:
		//line diksam.go.y:129
		{
			yyVAL.nodes = append(make([]dast.Node, 0), yyS[yypt-0].node)
		}
	case 19:
		//line diksam.go.y:133
		{
			yyVAL.nodes = append(yyS[yypt-1].nodes, yyS[yypt-0].node)
		}
	case 20:
		//line diksam.go.y:138
		{
			yyVAL.node = dast.NewExprStmt(yyS[yypt-0].token, yyS[yypt-1].node)
		}
	case 21:
		//line diksam.go.y:142
		{
			yyVAL.node = dast.NewDecl(yyS[yypt-1].token, newTypeSpec(yyS[yypt-2].typeSpec), yyS[yypt-1].token.Value.(string), nil)
		}
	case 22:
		//line diksam.go.y:146
		{
			yyVAL.node = dast.NewDecl(yyS[yypt-3].token, newTypeSpec(yyS[yypt-4].typeSpec), yyS[yypt-3].token.Value.(string), yyS[yypt-1].node)
		}
	case 23:
		//line diksam.go.y:150
		{
			yyVAL.node = dast.NewIf(yyS[yypt-5].token, yyS[yypt-3].node, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 24:
		//line diksam.go.y:154
		{
			yyVAL.node = dast.NewWhile(yyS[yypt-4].token, yyS[yypt-5].ident, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 25:
		//line diksam.go.y:158
		{
			yyVAL.node = dast.NewFor(yyS[yypt-8].token, yyS[yypt-9].ident, yyS[yypt-6].node, yyS[yypt-4].node, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 26:
		//line diksam.go.y:162
		{
			yyVAL.node = dast.NewForeach(yyS[yypt-6].token, yyS[yypt-7].ident, yyS[yypt-4].token.Value.(string), yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 27:
		//line diksam.go.y:166
		{
			yyVAL.node = dast.NewReturn(yyS[yypt-2].token, yyS[yypt-1].node)
		}
	case 28:
		//line diksam.go.y:170
		{
			yyVAL.node = dast.NewBreak(yyS[yypt-2].token, yyS[yypt-1].ident)
		}
	case 29:
		//line diksam.go.y:174
		{
			yyVAL.node = dast.NewContinue(yyS[yypt-2].token, yyS[yypt-1].ident)
		}
	case 30:
		//line diksam.go.y:178
		{
			yyVAL.node = dast.NewTry(yyS[yypt-6].token, yyS[yypt-5].node, yyS[yypt-2].token.Value.(string), yyS[yypt-0].node, nil)
		}
	case 31:
		//line diksam.go.y:182
		{
			yyVAL.node = dast.NewTry(yyS[yypt-3].token, yyS[yypt-2].node, "", nil, yyS[yypt-0].node)
		}
	case 32:
		//line diksam.go.y:186
		{
			yyVAL.node = dast.NewTry(yyS[yypt-8].token, yyS[yypt-7].node, yyS[yypt-4].token.Value.(string), yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 33:
		//line diksam.go.y:190
		{
			yyVAL.node = dast.NewThrow(yyS[yypt-2].token, yyS[yypt-1].node)
		}
	case 34:
		//line diksam.go.y:195
		{
			yyVAL.node = nil
		}
	case 35:
		//line diksam.go.y:199
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 36:
		//line diksam.go.y:203
		{
			yyVAL.node = dast.NewIf(yyS[yypt-5].token, yyS[yypt-3].node, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 37:
		//line diksam.go.y:208
		{
			yyVAL.ident = ""
		}
	case 38:
		//line diksam.go.y:212
		{
			yyVAL.ident = yyS[yypt-1].token.Value.(string)
		}
	case 39:
		//line diksam.go.y:217
		{
			yyVAL.node = nil
		}
	case 40:
		yyVAL.node = yyS[yypt-0].node
	case 41:
		//line diksam.go.y:223
		{
			yyVAL.ident = ""
		}
	case 42:
		//line diksam.go.y:227
		{
			yyVAL.ident = yyS[yypt-0].token.Value.(string)
		}
	case 43:
		yyVAL.node = yyS[yypt-0].node
	case 44:
		//line diksam.go.y:233
		{
			yyVAL.node = dast.NewCommaExpr(yyS[yypt-1].token, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 45:
		yyVAL.node = yyS[yypt-0].node
	case 46:
		//line diksam.go.y:239
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.NORMAL_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 47:
		//line diksam.go.y:243
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.ADD_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 48:
		//line diksam.go.y:247
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.SUB_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 49:
		//line diksam.go.y:251
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.MUL_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 50:
		//line diksam.go.y:255
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.DIV_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 51:
		//line diksam.go.y:259
		{
			yyVAL.node = dast.NewAssign(yyS[yypt-1].token, dast.MOD_ASSIGN, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 52:
		yyVAL.node = yyS[yypt-0].node
	case 53:
		//line diksam.go.y:265
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LOG_OR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 54:
		yyVAL.node = yyS[yypt-0].node
	case 55:
		//line diksam.go.y:271
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LOG_AND, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 56:
		yyVAL.node = yyS[yypt-0].node
	case 57:
		//line diksam.go.y:277
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.EQ, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 58:
		//line diksam.go.y:281
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.NE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 59:
		yyVAL.node = yyS[yypt-0].node
	case 60:
		//line diksam.go.y:287
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.GT, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 61:
		//line diksam.go.y:291
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LT, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 62:
		//line diksam.go.y:295
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.GE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 63:
		//line diksam.go.y:299
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.LE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 64:
		yyVAL.node = yyS[yypt-0].node
	case 65:
		//line diksam.go.y:305
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.ADD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 66:
		//line diksam.go.y:309
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.SUB, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 67:
		yyVAL.node = yyS[yypt-0].node
	case 68:
		//line diksam.go.y:315
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.MUL, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 69:
		//line diksam.go.y:319
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.DIV, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 70:
		//line diksam.go.y:323
		{
			yyVAL.node = dast.NewBinary(yyS[yypt-1].token, dast.MOD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 71:
		yyVAL.node = yyS[yypt-0].node
	case 72:
		//line diksam.go.y:329
		{
			yyVAL.node = dast.NewMinusExpr(yyS[yypt-1].token, yyS[yypt-0].node)
		}
	case 73:
		//line diksam.go.y:333
		{
			yyVAL.node = dast.NewLogNot(yyS[yypt-1].token, yyS[yypt-0].node)
		}
	case 74:
		yyVAL.node = yyS[yypt-0].node
	case 75:
		//line diksam.go.y:339
		{
			yyVAL.node = dast.NewFuncall(yyS[yypt-1].token, yyS[yypt-2].node, nil)
		}
	case 76:
		//line diksam.go.y:343
		{
			yyVAL.node = dast.NewFuncall(yyS[yypt-2].token, yyS[yypt-3].node, yyS[yypt-1].nodes)
		}
	case 77:
		//line diksam.go.y:347
		{
			yyVAL.node = dast.NewIncDec(yyS[yypt-0].token, yyS[yypt-1].node, dast.INC)
		}
	case 78:
		//line diksam.go.y:351
		{
			yyVAL.node = dast.NewIncDec(yyS[yypt-0].token, yyS[yypt-1].node, dast.DEC)
		}
	case 79:
		//line diksam.go.y:356
		{
			yyVAL.node = yyS[yypt-1].node
		}
	case 80:
		//line diksam.go.y:360
		{
			yyVAL.node = dast.NewIdentExpr(yyS[yypt-0].token, yyS[yypt-0].token.Value.(string))
		}
	case 81:
		//line diksam.go.y:364
		{
			yyVAL.node = dast.NewIntLiteral(yyS[yypt-0].token, yyS[yypt-0].token.Value.(int))
		}
	case 82:
		//line diksam.go.y:368
		{
			yyVAL.node = dast.NewDoubleLiteral(yyS[yypt-0].token, yyS[yypt-0].token.Value.(float64))
		}
	case 83:
		//line diksam.go.y:372
		{
			yyVAL.node = dast.NewStrLiteral(yyS[yypt-0].token, yyS[yypt-0].token.Value.(string))
		}
	case 84:
		//line diksam.go.y:376
		{
			yyVAL.node = dast.NewBooleanLiteral(yyS[yypt-0].token, true)
		}
	case 85:
		//line diksam.go.y:380
		{
			yyVAL.node = dast.NewBooleanLiteral(yyS[yypt-0].token, false)
		}
	case 86:
		//line diksam.go.y:385
		{
			yyVAL.nodes = append(make([]dast.Node, 0), yyS[yypt-0].node)
		}
	case 87:
		//line diksam.go.y:389
		{
			yyVAL.nodes = append(yyS[yypt-2].nodes, yyS[yypt-0].node)
		}
	}
	goto yystack /* stack new state and value */
}

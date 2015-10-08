package dkc

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/arikui1911/Godiksam/dast"
)

type Token struct {
	Symbol int
	Value  interface{}
	line   int
	column int
}

func (t *Token) Line() int   { return t.line }
func (t *Token) Column() int { return t.column }

func (t *Token) String() string {
	return fmt.Sprintf("&%T{%d:%d:%s:%v}", t, t.Line(), t.Column(), tokName(t.Symbol), t.Value)
}

type Parser struct {
	Tree        dast.Node
	src         *bufio.Reader
	ungetcheds  *list.List
	SrcFileName string
	line        int
	column      int
	lastToken   *Token
	lastErr     error
	ErrorCount  int
	OnError     func(*Parser, error)
	blocks      *list.List
}

func NewParser(src io.Reader, srcFileName string, initialLineNumber int) *Parser {
	return &Parser{
		src:         bufio.NewReader(src),
		ungetcheds:  list.New(),
		SrcFileName: srcFileName,
		line:        initialLineNumber,
		column:      -1,
		blocks:      list.New(),
	}
}

func (p *Parser) Parse() error {
	p.lastErr = nil
	yyParse(p)
	return p.lastErr
}

func (p *Parser) getch() (rune, error) {
	if p.ungetcheds.Len() > 0 {
		c := p.ungetcheds.Remove(p.ungetcheds.Front()).(rune)
		return c, nil
	}
	c, _, err := p.src.ReadRune()
	return c, err
}

// c == '\n' の時は p.column が狂う
func (p *Parser) ungetch(c rune) {
	p.column--
	if c == '\n' {
		p.line--
	}
	p.ungetcheds.PushFront(c)
}

func (p *Parser) PushBlock(b *dast.Block) {
	p.blocks.PushFront(b)
}

func (p *Parser) PopBlock() *dast.Block {
	if p.blocks.Len() <= 0 {
		return nil
	}
	return p.blocks.Remove(p.blocks.Front()).(*dast.Block)
}

func (p *Parser) Line() int {
	if p.lastToken == nil {
		return p.line
	}
	return p.lastToken.Line()
}

func (p *Parser) Column() int {
	if p.lastToken == nil {
		return p.column
	}
	return p.lastToken.Column()
}

func (p *Parser) Error(msg string) {
	p.reportError(fmt.Errorf("%s:%d:%d: %s", p.SrcFileName, p.Line(), p.Column(), msg))
}

func (p *Parser) reportError(e error) {
	p.ErrorCount++
	if p.OnError != nil {
		p.OnError(p, e)
		return
	}
	fmt.Fprintf(os.Stderr, "%s\n", e)
	p.lastErr = e
}

func (p *Parser) Lex(lval *yySymType) int {
	t := &Token{
		Symbol: 0,
		line:   p.line,
		column: p.column,
	}
	err := p.lexLoop(t)
	if err != nil && err != io.EOF {
		p.reportError(err)
		return 0
	}
	p.lastToken = t
	lval.token = t
	return t.Symbol
}

func (p *Parser) lexLoop(t *Token) error {
	buf := make([]rune, 0)
	state := kSTATE_INITIAL
	var str *stringLiteral
	var c rune
	var err error

	for {
		c, err = p.getch()
		if err != nil {
			break
		}
		p.column++
		if c == '\n' {
			p.line++
			p.column = -1
		}
		switch state {
		case kSTATE_INITIAL:
			t.line = p.line
			t.column = p.column
			switch {
			case unicode.IsSpace(c):
				// do nothing
			case c == '0':
				state = kSTATE_ZERO
			case unicode.IsDigit(c):
				buf = append(buf, c)
				state = kSTATE_INTEGER
			case c == '"':
				str = &stringLiteral{
					begLine:   p.line,
					begCol:    p.column,
					isCharLit: false,
				}
				state = kSTATE_STRING
			case c == '\'':
				str = &stringLiteral{
					begLine:   p.line,
					begCol:    p.column,
					isCharLit: true,
				}
				state = kSTATE_STRING
			case c == '/':
				state = kSTATE_SLASH
			case unicode.IsLetter(c):
				buf = append(buf, c)
				state = kSTATE_IDENT
			default:
				buf = append(buf, c)
				state = kSTATE_OPERATOR
			}
		case kSTATE_ZERO:
			switch {
			case c == '.':
				buf = append(buf, '0')
				state = kSTATE_FLOAT_DOT
			case c == 'x' || c == 'X':
				state = kSTATE_HEX
			default:
				p.ungetch(c)
				t.Value = 0
				t.Symbol = tINT_LITERAL
				return nil
			}
		case kSTATE_HEX:
			if isXDigit(c) {
				buf = append(buf, c)
				state = kSTATE_HEX_DIGITS
			} else {
				return fmt.Errorf("%s:%d:%d: missing number literal digits", p.SrcFileName, p.Line(), p.Column())
			}
		case kSTATE_HEX_DIGITS:
			if !isXDigit(c) {
				p.ungetch(c)
				t.Value = parseInt(string(buf), 16)
				t.Symbol = tINT_LITERAL
				return nil
			}
			buf = append(buf, c)
		case kSTATE_INTEGER:
			switch {
			case c == '.':
				state = kSTATE_FLOAT_DOT
			case unicode.IsDigit(c):
				buf = append(buf, c)
			default:
				p.ungetch(c)
				t.Value = parseInt(string(buf), 10)
				t.Symbol = tINT_LITERAL
				return nil
			}
		case kSTATE_FLOAT_DOT:
			if !unicode.IsDigit(c) {
				p.ungetch(c)
				p.ungetch('.')
				t.Value = parseInt(string(buf), 10)
				t.Symbol = tINT_LITERAL
				return nil
			}
			buf = append(buf, '.')
			buf = append(buf, c)
			state = kSTATE_FLOAT
		case kSTATE_FLOAT:
			switch {
			case unicode.IsDigit(c):
				buf = append(buf, c)
			case c == 'e' || c == 'E':
				buf = append(buf, c)
				state = kSTATE_FLOAT_E
			default:
				p.ungetch(c)
				t.Value = parseFloat(string(buf))
				t.Symbol = tDOUBLE_LITERAL
				return nil
			}
		case kSTATE_FLOAT_E:
			switch {
			case c == '+' || c == '-':
				buf = append(buf, c)
			case unicode.IsDigit(c):
				buf = append(buf, c)
				state = kSTATE_FLOAT_E_DIGITS
			default:
				return fmt.Errorf("missing number literal digits")
			}
		case kSTATE_FLOAT_E_DIGITS:
			if !unicode.IsDigit(c) {
				p.ungetch(c)
				t.Value = parseFloat(string(buf))
				t.Symbol = tDOUBLE_LITERAL
				return nil
			}
			buf = append(buf, c)
		case kSTATE_STRING:
			switch c {
			case '\\':
				state = kSTATE_STRING_ESC
			case '"':
				if !str.isCharLit {
					t.Value = string(buf)
					t.Symbol = tSTRING_LITERAL
					return nil
				}
				buf = append(buf, c)
			case '\'':
				if str.isCharLit {
					t.Value = string(buf)
					t.Symbol = tINT_LITERAL
					return nil
				}
				buf = append(buf, c)
			default:
				buf = append(buf, c)
			}
		case kSTATE_STRING_ESC:
			ch, ok := escapes[c]
			if ok {
				buf = append(buf, ch)
			} else {
				buf = append(buf, '\\')
				buf = append(buf, c)
			}
			state = kSTATE_STRING
		case kSTATE_IDENT:
			if unicode.IsLetter(c) {
				buf = append(buf, c)
			} else {
				p.ungetch(c)
				sym, ok := keywords[string(buf)]
				if ok {
					t.Symbol = sym
				} else {
					t.Value = string(buf)
					t.Symbol = tIDENT
				}
				return nil
			}
		case kSTATE_OPERATOR:
			buf = append(buf, c)
			if !isPrefixOfOperators(string(buf), operators) {
				p.ungetch(c)
				buf = buf[0 : len(buf)-1] // chop
				switch {
				case len(buf) <= 0:
					panic("must not happen")
				case len(buf) == 1:
					t.Symbol = int(buf[0])
					return nil
				default:
					t.Symbol = operators[string(buf)]
					return nil
				}
			}
		case kSTATE_SLASH:
			switch c {
			case '/':
				state = kSTATE_COMMENT_LINE
			case '*':
				state = kSTATE_COMMENT_BLOCK
			default:
				p.ungetch(c)
				buf = append(buf, '/')
				state = kSTATE_OPERATOR
			}
		case kSTATE_COMMENT_LINE:
			if c == '\n' {
				state = kSTATE_INITIAL
			}
		case kSTATE_COMMENT_BLOCK:
			if c == '*' {
				state = kSTATE_COMMENT_STAR
			}
		case kSTATE_COMMENT_STAR:
			if c == '/' {
				state = kSTATE_INITIAL
			} else {
				p.ungetch(c)
				state = kSTATE_COMMENT_BLOCK
			}
		}
	}

	if err == io.EOF && (state == kSTATE_STRING || state == kSTATE_STRING_ESC) {
		return fmt.Errorf("%s:%d:%d: missing string literal termination", p.SrcFileName, str.begLine, str.begCol)
	}

	return err
}

var xdigitLetters = map[rune]bool{
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'A': true,
	'B': true,
	'C': true,
	'D': true,
	'E': true,
	'F': true,
}

func isXDigit(c rune) bool {
	if unicode.IsDigit(c) {
		return true
	}
	_, ok := xdigitLetters[c]
	return ok
}

func parseInt(src string, base int) int {
	v, err := strconv.ParseInt(src, base, 0)
	if err != nil {
		panic(err)
	}
	return int(v)
}

func parseFloat(src string) float64 {
	v, err := strconv.ParseFloat(src, 0)
	if err != nil {
		panic(err)
	}
	return v
}

func isPrefixOfOperators(str string, table map[string]int) bool {
	for k, _ := range table {
		if strings.HasPrefix(k, str) {
			return true
		}
	}
	return false
}

var keywords = map[string]int{
	"if":             kIF,
	"else":           kELSE,
	"elsif":          kELSIF,
	"switch":         kSWITCH,
	"case":           kCASE,
	"default":        kDEFAULT,
	"while":          kWHILE,
	"do":             kDO,
	"for":            kFOR,
	"foreach":        kFOREACH,
	"return":         kRETURN,
	"break":          kBREAK,
	"continue":       kCONTINUE,
	"null":           kNULL,
	"true":           kTRUE,
	"false":          kFALSE,
	"try":            kTRY,
	"catch":          kCATCH,
	"finally":        kFINALLY,
	"throw":          kTHROW,
	"throws":         kTHROWS,
	"boolean":        kBOOLEAN,
	"void":           kVOID,
	"int":            kINT,
	"double":         kDOUBLE,
	"string":         kSTRING,
	"native_pointer": kNATIVE_POINTER,
	"new":            kNEW,
	"require":        kREQUIRE,
	"rename":         kRENAME,
	"class":          kCLASS,
	"interface":      kINTERFACE,
	"public":         kPUBLIC,
	"private":        kPRIVATE,
	"virtual":        kVIRTUAL,
	"override":       kOVERRIDE,
	"abstract":       kABSTRACT,
	"this":           kTHIS,
	"super":          kSUPER,
	"constructor":    kCONSTRUCTOR,
	"instanceof":     kINSTANCEOF,
	"delegate":       kDELEGATE,
	"enum":           kENUM,
	"final":          kFINAL,
	"const":          kCONST,
}

var operators = map[string]int{
	"&&": tLOG_AND,
	"||": tLOG_OR,
	"==": tEQ,
	"!=": tNE,
	">=": tGE,
	"<=": tLE,
	"+=": tADD_A,
	"-=": tSUB_A,
	"*=": tMUL_A,
	"/=": tDIV_A,
	"%=": tMOD_A,
	"++": tINC,
	"--": tDEC,
	"::": tDCAST_BEG,
	":>": tDCAST_END,
}

var escapes = map[rune]rune{
	'\\': '\\',
	'"':  '"',
	'\'': '\'',
	'n':  '\n',
	't':  '\t',
}

type stringLiteral struct {
	begLine   int
	begCol    int
	isCharLit bool
}

type state_t int

const (
	kSTATE_INITIAL state_t = iota
	kSTATE_SLASH
	kSTATE_COMMENT_LINE
	kSTATE_COMMENT_BLOCK
	kSTATE_COMMENT_STAR
	kSTATE_ZERO
	kSTATE_HEX
	kSTATE_HEX_DIGITS
	kSTATE_INTEGER
	kSTATE_FLOAT_DOT
	kSTATE_FLOAT
	kSTATE_FLOAT_E
	kSTATE_FLOAT_E_DIGITS
	kSTATE_STRING
	kSTATE_STRING_ESC
	kSTATE_IDENT
	kSTATE_OPERATOR
)

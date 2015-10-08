%{
package dkc

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


%}

%union {
	token    *Token
	ident     string
	node      dast.Node
	nodes     []dast.Node
	typeSpec  baseType
}

%token<token> kIF kELSE kELSIF kSWITCH kCASE kDEFAULT kWHILE kDO kFOR kFOREACH
			  kRETURN kBREAK kCONTINUE kNULL kTRUE kFALSE kTRY kCATCH kFINALLY
			  kTHROW kTHROWS kBOOLEAN kVOID kINT kDOUBLE kSTRING kNATIVE_POINTER
              kNEW kREQUIRE kRENAME kCLASS kINTERFACE kPUBLIC kPRIVATE kVIRTUAL
              kOVERRIDE kABSTRACT kTHIS kSUPER kCONSTRUCTOR kINSTANCEOF
              kDELEGATE kENUM kFINAL kCONST tEQ tNE tGE tLE tADD_A tSUB_A
              tMUL_A tDIV_A tMOD_A tINC tDEC tDCAST_BEG tDCAST_END tLOG_AND tLOG_OR
              tINT_LITERAL tDOUBLE_LITERAL tSTRING_LITERAL tREGEXP_LITERAL tIDENT

%type<node>  stmt block_beg block if_tail opt_expr expr assign log_or log_and
             equality relational additive multive unary postfix primary
%type<nodes> stmts params args
%type<token> '{' ';' ',' '=' '<' '>' '+' '-' '*' '/' '%' '!' '('
%type<ident> opt_label opt_ident
%type<typeSpec> type_spec




%%

trans_unit : def_or_stmt
           | trans_unit def_or_stmt

def_or_stmt : def_func
            | stmt
    {
        addTopLevelStmt(yylex, $1)
	}

type_spec : kBOOLEAN { $$ = ttBOOLEAN }
          | kINT     { $$ = ttINT }
          | kDOUBLE  { $$ = ttDOUBLE }
          | kSTRING  { $$ = ttSTRING }

def_func : type_spec tIDENT '(' params ')' block
    {
        defun(yylex, $2, $1, $2, $4, $6)
    }
 		 | type_spec tIDENT '(' ')' block
    {
        defun(yylex, $2, $1, $2, nil, $5)
    }
         | type_spec tIDENT '(' params ')' ';'
    {
        defun(yylex, $2, $1, $2, $4, nil)
    }
         | type_spec tIDENT '(' ')' ';'
 	{
        defun(yylex, $2, $1, $2, nil, nil)
    }

params : type_spec tIDENT
    {
        $$ = append(make([]dast.Node, 0), dast.NewParam($2, newTypeSpec($1), $2.Value.(string)))
    }
       | params ',' type_spec tIDENT
    {
        $$ = append($1, dast.NewParam($4, newTypeSpec($3), $4.Value.(string)))
    }

block_beg : '{'
    {
		$$ = openBlock(yylex, $1)
    }

block : block_beg '}'
    {
        $$ = closeBlock(yylex, $1, nil)
    }
      | block_beg stmts '}'
    {
        $$ = closeBlock(yylex, $1, $2)
    }

stmts : stmt
    {
        $$ = append(make([]dast.Node, 0), $1)
    }
      | stmts stmt
    {
        $$ = append($1, $2)
    }

stmt : expr ';'
    {
        $$ = dast.NewExprStmt($2, $1)
    }
     | type_spec tIDENT ';'
    {
        $$ = dast.NewDecl($2, newTypeSpec($1), $2.Value.(string), nil)
    }
     | type_spec tIDENT '=' expr ';'
    {
        $$ = dast.NewDecl($2, newTypeSpec($1), $2.Value.(string), $4)
    }
     | kIF '(' expr ')' block if_tail
    {
        $$ = dast.NewIf($1, $3, $5, $6)
    }
     | opt_label kWHILE '(' expr ')' block
    {
        $$ = dast.NewWhile($2, $1, $4, $6)
    }
     | opt_label kFOR '(' opt_expr ';' opt_expr ';' opt_expr ')' block
    {
        $$ = dast.NewFor($2, $1, $4, $6, $8, $10)
    }
     | opt_label kFOREACH '(' tIDENT ':' expr ')' block
    {
        $$ = dast.NewForeach($2, $1, $4.Value.(string), $6, $8)
    }
     | kRETURN opt_expr ';'
    {
        $$ = dast.NewReturn($1, $2)
    }
     | kBREAK opt_ident ';'
    {
        $$ = dast.NewBreak($1, $2)
    }
     | kCONTINUE opt_ident ';'
    {
        $$ = dast.NewContinue($1, $2)
    }
     | kTRY block kCATCH '(' tIDENT ')' block
    {
        $$ = dast.NewTry($1, $2, $5.Value.(string), $7, nil)
    }
     | kTRY block kFINALLY block
    {
        $$ = dast.NewTry($1, $2, "", nil, $4)
    }
     | kTRY block kCATCH '(' tIDENT ')' block kFINALLY block
    {
        $$ = dast.NewTry($1, $2, $5.Value.(string), $7, $9)
    }
     | kTHROW expr ';'
    {
        $$ = dast.NewThrow($1, $2)
    }

if_tail :
    {
        $$ = nil
    }
        | kELSE block
    {
        $$ = $2
    }
        | kELSIF '(' expr ')' block if_tail
    {
        $$ = dast.NewIf($1, $3, $5, $6)
    }

opt_label :
    {
        $$ = ""
    }
          | tIDENT ':'
	{
		$$ = $1.Value.(string)
	}

opt_expr :
    {
        $$ = nil
    }
         | expr

opt_ident :
    {
        $$ = ""
    }
          | tIDENT
	{
		$$ = $1.Value.(string)
	}

expr : assign
     | expr ',' assign
    {
        $$ = dast.NewCommaExpr($2, $1, $3)
    }

assign : log_or
       | postfix '=' assign
    {
        $$ = dast.NewAssign($2, dast.NORMAL_ASSIGN, $1, $3)
    }
       | postfix tADD_A assign
    {
        $$ = dast.NewAssign($2, dast.ADD_ASSIGN, $1, $3)
    }
       | postfix tSUB_A assign
    {
        $$ = dast.NewAssign($2, dast.SUB_ASSIGN, $1, $3)
    }
       | postfix tMUL_A assign
    {
        $$ = dast.NewAssign($2, dast.MUL_ASSIGN, $1, $3)
    }
       | postfix tDIV_A assign
    {
        $$ = dast.NewAssign($2, dast.DIV_ASSIGN, $1, $3)
    }
       | postfix tMOD_A assign
    {
        $$ = dast.NewAssign($2, dast.MOD_ASSIGN, $1, $3)
    }

log_or : log_and
       | log_or tLOG_OR log_and
    {
        $$ = dast.NewBinary($2, dast.LOG_OR, $1, $3)
    }

log_and : equality
        | log_and tLOG_AND equality
    {
        $$ = dast.NewBinary($2, dast.LOG_AND, $1, $3)
    }

equality : relational
         | equality tEQ relational
    {
        $$ = dast.NewBinary($2, dast.EQ, $1, $3)
    }
         | equality tNE relational
    {
        $$ = dast.NewBinary($2, dast.NE, $1, $3)
    }

relational : additive
           | relational '<' additive
    {
        $$ = dast.NewBinary($2, dast.GT, $1, $3)
    }
           | relational '>' additive
    {
        $$ = dast.NewBinary($2, dast.LT, $1, $3)
    }
           | relational tGE additive
    {
        $$ = dast.NewBinary($2, dast.GE, $1, $3)
    }
           | relational tLE additive
    {
        $$ = dast.NewBinary($2, dast.LE, $1, $3)
    }

additive : multive
         | additive '+' multive
    {
        $$ = dast.NewBinary($2, dast.ADD, $1, $3)
    }
         | additive '-' multive
    {
        $$ = dast.NewBinary($2, dast.SUB, $1, $3)
    }

multive : unary
        | multive '*' unary
    {
        $$ = dast.NewBinary($2, dast.MUL, $1, $3)
    }
        | multive '/' unary
    {
        $$ = dast.NewBinary($2, dast.DIV, $1, $3)
    }
        | multive '%' unary
    {
        $$ = dast.NewBinary($2, dast.MOD, $1, $3)
    }

unary : postfix
      | '-' unary
    {
        $$ = dast.NewMinusExpr($1, $2)
    }
      | '!' unary
    {
        $$ = dast.NewLogNot($1, $2)
    }

postfix : primary
        | postfix '(' ')'
    {
        $$ = dast.NewFuncall($2, $1, nil)
    }
        | postfix '(' args ')'
    {
        $$ = dast.NewFuncall($2, $1, $3)
    }
        | postfix tINC
    {
        $$ = dast.NewIncDec($2, $1, dast.INC)
    }
        | postfix tDEC
    {
        $$ = dast.NewIncDec($2, $1, dast.DEC)
    }

primary : '(' expr ')'
    {
        $$ = $2
    }
        | tIDENT
    {
        $$ = dast.NewIdentExpr($1, $1.Value.(string))
    }
        | tINT_LITERAL
    {
        $$ = dast.NewIntLiteral($1, $1.Value.(int))
    }
        | tDOUBLE_LITERAL
    {
        $$ = dast.NewDoubleLiteral($1, $1.Value.(float64))
    }
        | tSTRING_LITERAL
    {
        $$ = dast.NewStrLiteral($1, $1.Value.(string))
    }
        | kTRUE
    {
        $$ = dast.NewBooleanLiteral($1, true)
    }
        | kFALSE
    {
        $$ = dast.NewBooleanLiteral($1, false)
    }

args : assign
    {
        $$ = append(make([]dast.Node, 0), $1)
    }
     | args ',' assign
    {
        $$ = append($1, $3)
    }


%%

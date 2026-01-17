// Code generated from Blunderbuss.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Blunderbuss
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BlunderbussParser struct {
	*antlr.BaseParser
}

var BlunderbussParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func blunderbussParserInit() {
	staticData := &BlunderbussParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'+'", "'-'", "'*'", "'&'", "'%'", "'/'", "'||'", "'&&'", "'!'",
		"'ne'", "'eq'", "'le'", "'ge'", "'lt'", "'gt'", "", "'}'", "'{'", "']'",
		"'['", "'('", "')'", "','", "'='", "'extern'", "'elseif'", "'if'", "'else'",
		"'for'", "'break'", "'next'", "'return'", "'cache'", "'safe'", "'func'",
		"'''", "'\"'", "';'", "", "'any'", "'ptr'", "'int'", "'str'",
	}
	staticData.SymbolicNames = []string{
		"", "STRING", "PLUS", "MINUS", "MULT", "AMPS", "MOD", "DIV", "OR", "AND",
		"EXCL", "NOT_EQUAL", "EQUAL", "LE", "GE", "LT", "GT", "NUM", "RBRACE",
		"LBRACE", "RBRACKET", "LBRACKET", "LPAREN", "RPAREN", "COMMA", "ASSIGN",
		"EXTERN", "ELSEIF", "IF", "ELSE", "FOR", "BREAK", "NEXT", "RETURN",
		"CACHE", "SAFE", "FUNC", "SQUOTE", "DQUOTE", "SEMI", "TYPE", "ANY",
		"PTR", "INT", "STR", "WS", "LINE_COMMENT", "BLOCK_COMMENT", "ID",
	}
	staticData.RuleNames = []string{
		"program", "extern", "func", "args", "param", "call_args", "func_call",
		"block", "expr", "stmt", "if_stmt", "for_stmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 189, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 4, 0, 27, 8, 0, 11, 0, 12, 0, 28, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 2, 3, 2, 37, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 49, 8, 3, 10, 3, 12, 3, 52, 9, 3, 3,
		3, 54, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5,
		65, 8, 5, 10, 5, 12, 5, 68, 9, 5, 3, 5, 70, 8, 5, 1, 5, 1, 5, 1, 6, 3,
		6, 75, 8, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7, 82, 8, 7, 10, 7, 12, 7,
		85, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 102, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 118, 8,
		8, 10, 8, 12, 8, 121, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 3, 9, 158, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 171, 8, 10, 10, 10, 12,
		10, 174, 9, 10, 1, 10, 1, 10, 3, 10, 178, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 0, 1, 16, 12, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 0, 3, 2, 0, 4, 4, 6, 7, 1, 0, 2, 3, 2, 0,
		8, 9, 11, 16, 206, 0, 26, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 36, 1, 0, 0,
		0, 6, 44, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 60, 1, 0, 0, 0, 12, 74, 1,
		0, 0, 0, 14, 79, 1, 0, 0, 0, 16, 101, 1, 0, 0, 0, 18, 157, 1, 0, 0, 0,
		20, 159, 1, 0, 0, 0, 22, 179, 1, 0, 0, 0, 24, 27, 3, 4, 2, 0, 25, 27, 3,
		2, 1, 0, 26, 24, 1, 0, 0, 0, 26, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28,
		26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 5, 0, 0,
		1, 31, 1, 1, 0, 0, 0, 32, 33, 5, 26, 0, 0, 33, 34, 3, 4, 2, 0, 34, 3, 1,
		0, 0, 0, 35, 37, 5, 34, 0, 0, 36, 35, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37,
		38, 1, 0, 0, 0, 38, 39, 5, 36, 0, 0, 39, 40, 5, 48, 0, 0, 40, 41, 3, 6,
		3, 0, 41, 42, 5, 40, 0, 0, 42, 43, 3, 14, 7, 0, 43, 5, 1, 0, 0, 0, 44,
		53, 5, 22, 0, 0, 45, 50, 3, 8, 4, 0, 46, 47, 5, 24, 0, 0, 47, 49, 3, 8,
		4, 0, 48, 46, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51,
		1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 45, 1, 0, 0, 0,
		53, 54, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 5, 23, 0, 0, 56, 7, 1,
		0, 0, 0, 57, 58, 5, 40, 0, 0, 58, 59, 5, 48, 0, 0, 59, 9, 1, 0, 0, 0, 60,
		69, 5, 22, 0, 0, 61, 66, 3, 16, 8, 0, 62, 63, 5, 24, 0, 0, 63, 65, 3, 16,
		8, 0, 64, 62, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 61, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 5, 23, 0, 0, 72, 11, 1,
		0, 0, 0, 73, 75, 5, 35, 0, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75,
		76, 1, 0, 0, 0, 76, 77, 5, 48, 0, 0, 77, 78, 3, 10, 5, 0, 78, 13, 1, 0,
		0, 0, 79, 83, 5, 19, 0, 0, 80, 82, 3, 18, 9, 0, 81, 80, 1, 0, 0, 0, 82,
		85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0,
		0, 85, 83, 1, 0, 0, 0, 86, 87, 5, 18, 0, 0, 87, 15, 1, 0, 0, 0, 88, 89,
		6, 8, -1, 0, 89, 90, 5, 10, 0, 0, 90, 102, 3, 16, 8, 11, 91, 92, 5, 5,
		0, 0, 92, 102, 3, 16, 8, 10, 93, 94, 5, 22, 0, 0, 94, 95, 3, 16, 8, 0,
		95, 96, 5, 23, 0, 0, 96, 102, 1, 0, 0, 0, 97, 102, 5, 17, 0, 0, 98, 102,
		5, 1, 0, 0, 99, 102, 5, 48, 0, 0, 100, 102, 3, 12, 6, 0, 101, 88, 1, 0,
		0, 0, 101, 91, 1, 0, 0, 0, 101, 93, 1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 101,
		98, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 100, 1, 0, 0, 0, 102, 119, 1,
		0, 0, 0, 103, 104, 10, 9, 0, 0, 104, 105, 7, 0, 0, 0, 105, 118, 3, 16,
		8, 10, 106, 107, 10, 8, 0, 0, 107, 108, 7, 1, 0, 0, 108, 118, 3, 16, 8,
		9, 109, 110, 10, 7, 0, 0, 110, 111, 7, 2, 0, 0, 111, 118, 3, 16, 8, 8,
		112, 113, 10, 6, 0, 0, 113, 114, 5, 21, 0, 0, 114, 115, 3, 16, 8, 0, 115,
		116, 5, 20, 0, 0, 116, 118, 1, 0, 0, 0, 117, 103, 1, 0, 0, 0, 117, 106,
		1, 0, 0, 0, 117, 109, 1, 0, 0, 0, 117, 112, 1, 0, 0, 0, 118, 121, 1, 0,
		0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 17, 1, 0, 0, 0,
		121, 119, 1, 0, 0, 0, 122, 123, 5, 40, 0, 0, 123, 124, 5, 48, 0, 0, 124,
		125, 5, 25, 0, 0, 125, 126, 3, 16, 8, 0, 126, 127, 5, 39, 0, 0, 127, 158,
		1, 0, 0, 0, 128, 129, 5, 40, 0, 0, 129, 130, 5, 48, 0, 0, 130, 158, 5,
		39, 0, 0, 131, 132, 5, 48, 0, 0, 132, 133, 5, 25, 0, 0, 133, 134, 3, 16,
		8, 0, 134, 135, 5, 39, 0, 0, 135, 158, 1, 0, 0, 0, 136, 137, 5, 48, 0,
		0, 137, 138, 5, 21, 0, 0, 138, 139, 3, 16, 8, 0, 139, 140, 5, 20, 0, 0,
		140, 141, 5, 25, 0, 0, 141, 142, 3, 16, 8, 0, 142, 143, 5, 39, 0, 0, 143,
		158, 1, 0, 0, 0, 144, 145, 5, 33, 0, 0, 145, 146, 3, 16, 8, 0, 146, 147,
		5, 39, 0, 0, 147, 158, 1, 0, 0, 0, 148, 149, 3, 12, 6, 0, 149, 150, 5,
		39, 0, 0, 150, 158, 1, 0, 0, 0, 151, 152, 5, 32, 0, 0, 152, 158, 5, 39,
		0, 0, 153, 154, 5, 31, 0, 0, 154, 158, 5, 39, 0, 0, 155, 158, 3, 20, 10,
		0, 156, 158, 3, 22, 11, 0, 157, 122, 1, 0, 0, 0, 157, 128, 1, 0, 0, 0,
		157, 131, 1, 0, 0, 0, 157, 136, 1, 0, 0, 0, 157, 144, 1, 0, 0, 0, 157,
		148, 1, 0, 0, 0, 157, 151, 1, 0, 0, 0, 157, 153, 1, 0, 0, 0, 157, 155,
		1, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158, 19, 1, 0, 0, 0, 159, 160, 5, 28,
		0, 0, 160, 161, 5, 22, 0, 0, 161, 162, 3, 16, 8, 0, 162, 163, 5, 23, 0,
		0, 163, 172, 3, 14, 7, 0, 164, 165, 5, 27, 0, 0, 165, 166, 5, 22, 0, 0,
		166, 167, 3, 16, 8, 0, 167, 168, 5, 23, 0, 0, 168, 169, 3, 14, 7, 0, 169,
		171, 1, 0, 0, 0, 170, 164, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170,
		1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 177, 1, 0, 0, 0, 174, 172, 1, 0,
		0, 0, 175, 176, 5, 29, 0, 0, 176, 178, 3, 14, 7, 0, 177, 175, 1, 0, 0,
		0, 177, 178, 1, 0, 0, 0, 178, 21, 1, 0, 0, 0, 179, 180, 5, 30, 0, 0, 180,
		181, 5, 22, 0, 0, 181, 182, 3, 18, 9, 0, 182, 183, 3, 16, 8, 0, 183, 184,
		5, 39, 0, 0, 184, 185, 3, 18, 9, 0, 185, 186, 5, 23, 0, 0, 186, 187, 3,
		14, 7, 0, 187, 23, 1, 0, 0, 0, 15, 26, 28, 36, 50, 53, 66, 69, 74, 83,
		101, 117, 119, 157, 172, 177,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BlunderbussParserInit initializes any static state used to implement BlunderbussParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBlunderbussParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BlunderbussParserInit() {
	staticData := &BlunderbussParserStaticData
	staticData.once.Do(blunderbussParserInit)
}

// NewBlunderbussParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBlunderbussParser(input antlr.TokenStream) *BlunderbussParser {
	BlunderbussParserInit()
	this := new(BlunderbussParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BlunderbussParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Blunderbuss.g4"

	return this
}

// BlunderbussParser tokens.
const (
	BlunderbussParserEOF           = antlr.TokenEOF
	BlunderbussParserSTRING        = 1
	BlunderbussParserPLUS          = 2
	BlunderbussParserMINUS         = 3
	BlunderbussParserMULT          = 4
	BlunderbussParserAMPS          = 5
	BlunderbussParserMOD           = 6
	BlunderbussParserDIV           = 7
	BlunderbussParserOR            = 8
	BlunderbussParserAND           = 9
	BlunderbussParserEXCL          = 10
	BlunderbussParserNOT_EQUAL     = 11
	BlunderbussParserEQUAL         = 12
	BlunderbussParserLE            = 13
	BlunderbussParserGE            = 14
	BlunderbussParserLT            = 15
	BlunderbussParserGT            = 16
	BlunderbussParserNUM           = 17
	BlunderbussParserRBRACE        = 18
	BlunderbussParserLBRACE        = 19
	BlunderbussParserRBRACKET      = 20
	BlunderbussParserLBRACKET      = 21
	BlunderbussParserLPAREN        = 22
	BlunderbussParserRPAREN        = 23
	BlunderbussParserCOMMA         = 24
	BlunderbussParserASSIGN        = 25
	BlunderbussParserEXTERN        = 26
	BlunderbussParserELSEIF        = 27
	BlunderbussParserIF            = 28
	BlunderbussParserELSE          = 29
	BlunderbussParserFOR           = 30
	BlunderbussParserBREAK         = 31
	BlunderbussParserNEXT          = 32
	BlunderbussParserRETURN        = 33
	BlunderbussParserCACHE         = 34
	BlunderbussParserSAFE          = 35
	BlunderbussParserFUNC          = 36
	BlunderbussParserSQUOTE        = 37
	BlunderbussParserDQUOTE        = 38
	BlunderbussParserSEMI          = 39
	BlunderbussParserTYPE          = 40
	BlunderbussParserANY           = 41
	BlunderbussParserPTR           = 42
	BlunderbussParserINT           = 43
	BlunderbussParserSTR           = 44
	BlunderbussParserWS            = 45
	BlunderbussParserLINE_COMMENT  = 46
	BlunderbussParserBLOCK_COMMENT = 47
	BlunderbussParserID            = 48
)

// BlunderbussParser rules.
const (
	BlunderbussParserRULE_program   = 0
	BlunderbussParserRULE_extern    = 1
	BlunderbussParserRULE_func      = 2
	BlunderbussParserRULE_args      = 3
	BlunderbussParserRULE_param     = 4
	BlunderbussParserRULE_call_args = 5
	BlunderbussParserRULE_func_call = 6
	BlunderbussParserRULE_block     = 7
	BlunderbussParserRULE_expr      = 8
	BlunderbussParserRULE_stmt      = 9
	BlunderbussParserRULE_if_stmt   = 10
	BlunderbussParserRULE_for_stmt  = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFunc_() []IFuncContext
	Func_(i int) IFuncContext
	AllExtern() []IExternContext
	Extern(i int) IExternContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserEOF, 0)
}

func (s *ProgramContext) AllFunc_() []IFuncContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncContext); ok {
			len++
		}
	}

	tst := make([]IFuncContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncContext); ok {
			tst[i] = t.(IFuncContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Func_(i int) IFuncContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *ProgramContext) AllExtern() []IExternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExternContext); ok {
			len++
		}
	}

	tst := make([]IExternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExternContext); ok {
			tst[i] = t.(IExternContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Extern(i int) IExternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BlunderbussParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&85966454784) != 0) {
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case BlunderbussParserCACHE, BlunderbussParserFUNC:
			{
				p.SetState(24)
				p.Func_()
			}

		case BlunderbussParserEXTERN:
			{
				p.SetState(25)
				p.Extern()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(BlunderbussParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternContext is an interface to support dynamic dispatch.
type IExternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTERN() antlr.TerminalNode
	Func_() IFuncContext

	// IsExternContext differentiates from other interfaces.
	IsExternContext()
}

type ExternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternContext() *ExternContext {
	var p = new(ExternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_extern
	return p
}

func InitEmptyExternContext(p *ExternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_extern
}

func (*ExternContext) IsExternContext() {}

func NewExternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternContext {
	var p = new(ExternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_extern

	return p
}

func (s *ExternContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternContext) EXTERN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserEXTERN, 0)
}

func (s *ExternContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *ExternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterExtern(s)
	}
}

func (s *ExternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitExtern(s)
	}
}

func (s *ExternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitExtern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Extern() (localctx IExternContext) {
	localctx = NewExternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BlunderbussParserRULE_extern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(BlunderbussParserEXTERN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Func_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	Args() IArgsContext
	TYPE() antlr.TerminalNode
	Block() IBlockContext
	CACHE() antlr.TerminalNode

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) FUNC() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserFUNC, 0)
}

func (s *FuncContext) ID() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserID, 0)
}

func (s *FuncContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *FuncContext) TYPE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserTYPE, 0)
}

func (s *FuncContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncContext) CACHE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserCACHE, 0)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BlunderbussParserRULE_func)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserCACHE {
		{
			p.SetState(35)
			p.Match(BlunderbussParserCACHE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(38)
		p.Match(BlunderbussParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(BlunderbussParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Args()
	}
	{
		p.SetState(41)
		p.Match(BlunderbussParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLPAREN, 0)
}

func (s *ArgsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRPAREN, 0)
}

func (s *ArgsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserCOMMA, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BlunderbussParserRULE_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserTYPE {
		{
			p.SetState(45)
			p.Param()
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BlunderbussParserCOMMA {
			{
				p.SetState(46)
				p.Match(BlunderbussParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(47)
				p.Param()
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(55)
		p.Match(BlunderbussParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) TYPE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserTYPE, 0)
}

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserID, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BlunderbussParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(BlunderbussParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(BlunderbussParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICall_argsContext is an interface to support dynamic dispatch.
type ICall_argsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCall_argsContext differentiates from other interfaces.
	IsCall_argsContext()
}

type Call_argsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_argsContext() *Call_argsContext {
	var p = new(Call_argsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_call_args
	return p
}

func InitEmptyCall_argsContext(p *Call_argsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_call_args
}

func (*Call_argsContext) IsCall_argsContext() {}

func NewCall_argsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_argsContext {
	var p = new(Call_argsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_call_args

	return p
}

func (s *Call_argsContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_argsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLPAREN, 0)
}

func (s *Call_argsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRPAREN, 0)
}

func (s *Call_argsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Call_argsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Call_argsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserCOMMA)
}

func (s *Call_argsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserCOMMA, i)
}

func (s *Call_argsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_argsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_argsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterCall_args(s)
	}
}

func (s *Call_argsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitCall_args(s)
	}
}

func (s *Call_argsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitCall_args(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Call_args() (localctx ICall_argsContext) {
	localctx = NewCall_argsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BlunderbussParserRULE_call_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281509340775458) != 0 {
		{
			p.SetState(61)
			p.expr(0)
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BlunderbussParserCOMMA {
			{
				p.SetState(62)
				p.Match(BlunderbussParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(63)
				p.expr(0)
			}

			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(71)
		p.Match(BlunderbussParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Call_args() ICall_argsContext
	SAFE() antlr.TerminalNode

	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) ID() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserID, 0)
}

func (s *Func_callContext) Call_args() ICall_argsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_argsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_argsContext)
}

func (s *Func_callContext) SAFE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserSAFE, 0)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterFunc_call(s)
	}
}

func (s *Func_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitFunc_call(s)
	}
}

func (s *Func_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitFunc_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BlunderbussParserRULE_func_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserSAFE {
		{
			p.SetState(73)
			p.Match(BlunderbussParserSAFE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(76)
		p.Match(BlunderbussParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Call_args()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRBRACE, 0)
}

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BlunderbussParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(BlunderbussParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&282625222639616) != 0 {
		{
			p.SetState(80)
			p.Stmt()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(BlunderbussParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	EXCL() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AMPS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	NUM() antlr.TerminalNode
	STRING() antlr.TerminalNode
	ID() antlr.TerminalNode
	Func_call() IFunc_callContext
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) EXCL() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserEXCL, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AMPS() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserAMPS, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRPAREN, 0)
}

func (s *ExprContext) NUM() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserNUM, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserSTRING, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserID, 0)
}

func (s *ExprContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserMOD, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserMINUS, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserOR, 0)
}

func (s *ExprContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserEQUAL, 0)
}

func (s *ExprContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserNOT_EQUAL, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLE, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserGE, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLT, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserGT, 0)
}

func (s *ExprContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLBRACKET, 0)
}

func (s *ExprContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRBRACKET, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *BlunderbussParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, BlunderbussParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(89)
			p.Match(BlunderbussParserEXCL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.expr(11)
		}

	case 2:
		{
			p.SetState(91)
			p.Match(BlunderbussParserAMPS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.expr(10)
		}

	case 3:
		{
			p.SetState(93)
			p.Match(BlunderbussParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.expr(0)
		}
		{
			p.SetState(95)
			p.Match(BlunderbussParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(97)
			p.Match(BlunderbussParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(98)
			p.Match(BlunderbussParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(99)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(100)
			p.Func_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BlunderbussParserRULE_expr)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(104)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&208) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(105)
					p.expr(10)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BlunderbussParserRULE_expr)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(107)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == BlunderbussParserPLUS || _la == BlunderbussParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(108)
					p.expr(9)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BlunderbussParserRULE_expr)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(110)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&129792) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(111)
					p.expr(8)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BlunderbussParserRULE_expr)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(113)
					p.Match(BlunderbussParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(114)
					p.expr(0)
				}
				{
					p.SetState(115)
					p.Match(BlunderbussParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	SEMI() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Func_call() IFunc_callContext
	NEXT() antlr.TerminalNode
	BREAK() antlr.TerminalNode
	If_stmt() IIf_stmtContext
	For_stmt() IFor_stmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) TYPE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserTYPE, 0)
}

func (s *StmtContext) ID() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserID, 0)
}

func (s *StmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserASSIGN, 0)
}

func (s *StmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserSEMI, 0)
}

func (s *StmtContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLBRACKET, 0)
}

func (s *StmtContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRBRACKET, 0)
}

func (s *StmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRETURN, 0)
}

func (s *StmtContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *StmtContext) NEXT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserNEXT, 0)
}

func (s *StmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserBREAK, 0)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BlunderbussParserRULE_stmt)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(BlunderbussParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(BlunderbussParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.expr(0)
		}
		{
			p.SetState(126)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(BlunderbussParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(BlunderbussParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expr(0)
		}
		{
			p.SetState(134)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(BlunderbussParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.expr(0)
		}
		{
			p.SetState(139)
			p.Match(BlunderbussParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(BlunderbussParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.expr(0)
		}
		{
			p.SetState(142)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(144)
			p.Match(BlunderbussParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.expr(0)
		}
		{
			p.SetState(146)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(148)
			p.Func_call()
		}
		{
			p.SetState(149)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(151)
			p.Match(BlunderbussParserNEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(153)
			p.Match(BlunderbussParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(155)
			p.If_stmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(156)
			p.For_stmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllELSEIF() []antlr.TerminalNode
	ELSEIF(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) IF() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserIF, 0)
}

func (s *If_stmtContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserLPAREN)
}

func (s *If_stmtContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLPAREN, i)
}

func (s *If_stmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserRPAREN)
}

func (s *If_stmtContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRPAREN, i)
}

func (s *If_stmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *If_stmtContext) AllELSEIF() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserELSEIF)
}

func (s *If_stmtContext) ELSEIF(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserELSEIF, i)
}

func (s *If_stmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserELSE, 0)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (s *If_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitIf_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BlunderbussParserRULE_if_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(BlunderbussParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.expr(0)
	}
	{
		p.SetState(162)
		p.Match(BlunderbussParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Block()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BlunderbussParserELSEIF {
		{
			p.SetState(164)
			p.Match(BlunderbussParserELSEIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(BlunderbussParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.expr(0)
		}
		{
			p.SetState(167)
			p.Match(BlunderbussParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Block()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserELSE {
		{
			p.SetState(175)
			p.Match(BlunderbussParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Expr() IExprContext
	SEMI() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserFOR, 0)
}

func (s *For_stmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLPAREN, 0)
}

func (s *For_stmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *For_stmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *For_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *For_stmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserSEMI, 0)
}

func (s *For_stmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRPAREN, 0)
}

func (s *For_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterFor_stmt(s)
	}
}

func (s *For_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitFor_stmt(s)
	}
}

func (s *For_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitFor_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BlunderbussParserRULE_for_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(BlunderbussParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Stmt()
	}
	{
		p.SetState(182)
		p.expr(0)
	}
	{
		p.SetState(183)
		p.Match(BlunderbussParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Stmt()
	}
	{
		p.SetState(185)
		p.Match(BlunderbussParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *BlunderbussParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BlunderbussParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

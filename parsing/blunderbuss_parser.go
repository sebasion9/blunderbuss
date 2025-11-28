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
		"", "", "'+'", "'-'", "'*'", "'/'", "'||'", "'&&'", "'<<'", "'>>'",
		"'!'", "'=='", "'=<'", "'=>'", "'<'", "'>'", "", "'}'", "'{'", "'('",
		"')'", "','", "'='", "'elseif'", "'if'", "'else'", "'for'", "'break'",
		"'next'", "'return'", "'cache'", "'safe'", "'effect'", "'lazy'", "'func'",
		"'''", "'\"'", "';'", "", "'int'", "'byte'", "'str'",
	}
	staticData.SymbolicNames = []string{
		"", "STRING", "PLUS", "MINUS", "MULT", "DIV", "OR", "AND", "LSHIFT",
		"RSHIFT", "EXCL", "EQUAL", "LE", "GE", "LT", "GT", "NUM", "RBRACE",
		"LBRACE", "LPAREN", "RPAREN", "COMMA", "ASSIGN", "ELSEIF", "IF", "ELSE",
		"FOR", "BREAK", "NEXT", "RETURN", "CACHE", "SAFE", "EFFECT", "LAZY",
		"FUNC", "SQUOTE", "DQUOTE", "SEMI", "TYPE", "INT", "BYTE", "STR", "WS",
		"LINE_COMMENT", "BLOCK_COMMENT", "ID",
	}
	staticData.RuleNames = []string{
		"program", "func", "args", "param", "call_args", "func_call", "block",
		"expr", "stmt", "effect_block", "if_stmt", "for_stmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 194, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 4, 0, 26, 8, 0, 11, 0, 12, 0, 27, 1, 1, 3, 1, 31,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2,
		43, 8, 2, 10, 2, 12, 2, 46, 9, 2, 3, 2, 48, 8, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 59, 8, 4, 10, 4, 12, 4, 62, 9, 4,
		3, 4, 64, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 74,
		8, 6, 10, 6, 12, 6, 77, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 92, 8, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 103, 8, 7, 10, 7, 12, 7, 106, 9,
		7, 1, 8, 3, 8, 109, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 130, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 141, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 146, 8, 9, 10, 9, 12, 9, 149, 9,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 5, 10, 164, 8, 10, 10, 10, 12, 10, 167, 9, 10, 1,
		10, 1, 10, 3, 10, 171, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		3, 11, 179, 8, 11, 1, 11, 1, 11, 3, 11, 183, 8, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 189, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 0, 1, 14, 12, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 4, 2, 0, 30, 30, 33, 33, 2,
		0, 6, 7, 11, 15, 1, 0, 4, 5, 2, 0, 2, 3, 8, 9, 213, 0, 25, 1, 0, 0, 0,
		2, 30, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 51, 1, 0, 0, 0, 8, 54, 1, 0, 0,
		0, 10, 67, 1, 0, 0, 0, 12, 70, 1, 0, 0, 0, 14, 91, 1, 0, 0, 0, 16, 140,
		1, 0, 0, 0, 18, 142, 1, 0, 0, 0, 20, 152, 1, 0, 0, 0, 22, 172, 1, 0, 0,
		0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25,
		1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 1, 1, 0, 0, 0, 29, 31, 7, 0, 0, 0,
		30, 29, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 5,
		34, 0, 0, 33, 34, 5, 45, 0, 0, 34, 35, 3, 4, 2, 0, 35, 36, 5, 38, 0, 0,
		36, 37, 3, 12, 6, 0, 37, 3, 1, 0, 0, 0, 38, 47, 5, 19, 0, 0, 39, 44, 3,
		6, 3, 0, 40, 41, 5, 21, 0, 0, 41, 43, 3, 6, 3, 0, 42, 40, 1, 0, 0, 0, 43,
		46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 48, 1, 0, 0,
		0, 46, 44, 1, 0, 0, 0, 47, 39, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49,
		1, 0, 0, 0, 49, 50, 5, 20, 0, 0, 50, 5, 1, 0, 0, 0, 51, 52, 5, 38, 0, 0,
		52, 53, 5, 45, 0, 0, 53, 7, 1, 0, 0, 0, 54, 63, 5, 19, 0, 0, 55, 60, 3,
		14, 7, 0, 56, 57, 5, 21, 0, 0, 57, 59, 3, 14, 7, 0, 58, 56, 1, 0, 0, 0,
		59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 64, 1,
		0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 55, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 66, 5, 20, 0, 0, 66, 9, 1, 0, 0, 0, 67, 68, 5, 45,
		0, 0, 68, 69, 3, 8, 4, 0, 69, 11, 1, 0, 0, 0, 70, 75, 5, 18, 0, 0, 71,
		74, 3, 18, 9, 0, 72, 74, 3, 16, 8, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0,
		0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78,
		1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 79, 5, 17, 0, 0, 79, 13, 1, 0, 0, 0,
		80, 81, 6, 7, -1, 0, 81, 82, 5, 10, 0, 0, 82, 92, 3, 14, 7, 9, 83, 84,
		5, 19, 0, 0, 84, 85, 3, 14, 7, 0, 85, 86, 5, 20, 0, 0, 86, 92, 1, 0, 0,
		0, 87, 92, 5, 16, 0, 0, 88, 92, 5, 1, 0, 0, 89, 92, 5, 45, 0, 0, 90, 92,
		3, 10, 5, 0, 91, 80, 1, 0, 0, 0, 91, 83, 1, 0, 0, 0, 91, 87, 1, 0, 0, 0,
		91, 88, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 90, 1, 0, 0, 0, 92, 104, 1,
		0, 0, 0, 93, 94, 10, 8, 0, 0, 94, 95, 7, 1, 0, 0, 95, 103, 3, 14, 7, 9,
		96, 97, 10, 7, 0, 0, 97, 98, 7, 2, 0, 0, 98, 103, 3, 14, 7, 8, 99, 100,
		10, 6, 0, 0, 100, 101, 7, 3, 0, 0, 101, 103, 3, 14, 7, 7, 102, 93, 1, 0,
		0, 0, 102, 96, 1, 0, 0, 0, 102, 99, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104,
		102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 15, 1, 0, 0, 0, 106, 104, 1,
		0, 0, 0, 107, 109, 5, 33, 0, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0,
		0, 109, 110, 1, 0, 0, 0, 110, 111, 5, 38, 0, 0, 111, 112, 5, 45, 0, 0,
		112, 113, 5, 22, 0, 0, 113, 114, 3, 14, 7, 0, 114, 115, 5, 37, 0, 0, 115,
		141, 1, 0, 0, 0, 116, 117, 5, 38, 0, 0, 117, 118, 5, 45, 0, 0, 118, 141,
		5, 37, 0, 0, 119, 120, 5, 45, 0, 0, 120, 121, 5, 22, 0, 0, 121, 122, 3,
		14, 7, 0, 122, 123, 5, 37, 0, 0, 123, 141, 1, 0, 0, 0, 124, 125, 5, 29,
		0, 0, 125, 126, 3, 14, 7, 0, 126, 127, 5, 37, 0, 0, 127, 141, 1, 0, 0,
		0, 128, 130, 5, 31, 0, 0, 129, 128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130,
		131, 1, 0, 0, 0, 131, 132, 3, 10, 5, 0, 132, 133, 5, 37, 0, 0, 133, 141,
		1, 0, 0, 0, 134, 135, 5, 28, 0, 0, 135, 141, 5, 37, 0, 0, 136, 137, 5,
		27, 0, 0, 137, 141, 5, 37, 0, 0, 138, 141, 3, 20, 10, 0, 139, 141, 3, 22,
		11, 0, 140, 108, 1, 0, 0, 0, 140, 116, 1, 0, 0, 0, 140, 119, 1, 0, 0, 0,
		140, 124, 1, 0, 0, 0, 140, 129, 1, 0, 0, 0, 140, 134, 1, 0, 0, 0, 140,
		136, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 139, 1, 0, 0, 0, 141, 17, 1,
		0, 0, 0, 142, 143, 5, 32, 0, 0, 143, 147, 5, 18, 0, 0, 144, 146, 3, 16,
		8, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150,
		151, 5, 17, 0, 0, 151, 19, 1, 0, 0, 0, 152, 153, 5, 24, 0, 0, 153, 154,
		5, 19, 0, 0, 154, 155, 3, 14, 7, 0, 155, 156, 5, 20, 0, 0, 156, 165, 3,
		12, 6, 0, 157, 158, 5, 23, 0, 0, 158, 159, 5, 19, 0, 0, 159, 160, 3, 14,
		7, 0, 160, 161, 5, 20, 0, 0, 161, 162, 3, 12, 6, 0, 162, 164, 1, 0, 0,
		0, 163, 157, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 170, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 169,
		5, 25, 0, 0, 169, 171, 3, 12, 6, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1,
		0, 0, 0, 171, 21, 1, 0, 0, 0, 172, 173, 5, 26, 0, 0, 173, 178, 5, 19, 0,
		0, 174, 175, 5, 38, 0, 0, 175, 176, 5, 45, 0, 0, 176, 177, 5, 22, 0, 0,
		177, 179, 3, 14, 7, 0, 178, 174, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 182, 5, 37, 0, 0, 181, 183, 3, 14, 7, 0, 182, 181,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 188, 5, 37,
		0, 0, 185, 186, 5, 45, 0, 0, 186, 187, 5, 22, 0, 0, 187, 189, 3, 14, 7,
		0, 188, 185, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		191, 5, 20, 0, 0, 191, 192, 3, 12, 6, 0, 192, 23, 1, 0, 0, 0, 20, 27, 30,
		44, 47, 60, 63, 73, 75, 91, 102, 104, 108, 129, 140, 147, 165, 170, 178,
		182, 188,
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
	BlunderbussParserDIV           = 5
	BlunderbussParserOR            = 6
	BlunderbussParserAND           = 7
	BlunderbussParserLSHIFT        = 8
	BlunderbussParserRSHIFT        = 9
	BlunderbussParserEXCL          = 10
	BlunderbussParserEQUAL         = 11
	BlunderbussParserLE            = 12
	BlunderbussParserGE            = 13
	BlunderbussParserLT            = 14
	BlunderbussParserGT            = 15
	BlunderbussParserNUM           = 16
	BlunderbussParserRBRACE        = 17
	BlunderbussParserLBRACE        = 18
	BlunderbussParserLPAREN        = 19
	BlunderbussParserRPAREN        = 20
	BlunderbussParserCOMMA         = 21
	BlunderbussParserASSIGN        = 22
	BlunderbussParserELSEIF        = 23
	BlunderbussParserIF            = 24
	BlunderbussParserELSE          = 25
	BlunderbussParserFOR           = 26
	BlunderbussParserBREAK         = 27
	BlunderbussParserNEXT          = 28
	BlunderbussParserRETURN        = 29
	BlunderbussParserCACHE         = 30
	BlunderbussParserSAFE          = 31
	BlunderbussParserEFFECT        = 32
	BlunderbussParserLAZY          = 33
	BlunderbussParserFUNC          = 34
	BlunderbussParserSQUOTE        = 35
	BlunderbussParserDQUOTE        = 36
	BlunderbussParserSEMI          = 37
	BlunderbussParserTYPE          = 38
	BlunderbussParserINT           = 39
	BlunderbussParserBYTE          = 40
	BlunderbussParserSTR           = 41
	BlunderbussParserWS            = 42
	BlunderbussParserLINE_COMMENT  = 43
	BlunderbussParserBLOCK_COMMENT = 44
	BlunderbussParserID            = 45
)

// BlunderbussParser rules.
const (
	BlunderbussParserRULE_program      = 0
	BlunderbussParserRULE_func         = 1
	BlunderbussParserRULE_args         = 2
	BlunderbussParserRULE_param        = 3
	BlunderbussParserRULE_call_args    = 4
	BlunderbussParserRULE_func_call    = 5
	BlunderbussParserRULE_block        = 6
	BlunderbussParserRULE_expr         = 7
	BlunderbussParserRULE_stmt         = 8
	BlunderbussParserRULE_effect_block = 9
	BlunderbussParserRULE_if_stmt      = 10
	BlunderbussParserRULE_for_stmt     = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFunc_() []IFuncContext
	Func_(i int) IFuncContext

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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26843545600) != 0) {
		{
			p.SetState(24)
			p.Func_()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	LAZY() antlr.TerminalNode

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

func (s *FuncContext) LAZY() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLAZY, 0)
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
	p.EnterRule(localctx, 2, BlunderbussParserRULE_func)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserCACHE || _la == BlunderbussParserLAZY {
		{
			p.SetState(29)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BlunderbussParserCACHE || _la == BlunderbussParserLAZY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(32)
		p.Match(BlunderbussParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(BlunderbussParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Args()
	}
	{
		p.SetState(35)
		p.Match(BlunderbussParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
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
	p.EnterRule(localctx, 4, BlunderbussParserRULE_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserTYPE {
		{
			p.SetState(39)
			p.Param()
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BlunderbussParserCOMMA {
			{
				p.SetState(40)
				p.Match(BlunderbussParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(41)
				p.Param()
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(49)
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
	p.EnterRule(localctx, 6, BlunderbussParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(BlunderbussParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
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
	p.EnterRule(localctx, 8, BlunderbussParserRULE_call_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35184372679682) != 0 {
		{
			p.SetState(55)
			p.expr(0)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BlunderbussParserCOMMA {
			{
				p.SetState(56)
				p.Match(BlunderbussParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(57)
				p.expr(0)
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(65)
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
	p.EnterRule(localctx, 10, BlunderbussParserRULE_func_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(BlunderbussParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
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
	AllEffect_block() []IEffect_blockContext
	Effect_block(i int) IEffect_blockContext
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

func (s *BlockContext) AllEffect_block() []IEffect_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEffect_blockContext); ok {
			len++
		}
	}

	tst := make([]IEffect_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEffect_blockContext); ok {
			tst[i] = t.(IEffect_blockContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Effect_block(i int) IEffect_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffect_blockContext); ok {
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

	return t.(IEffect_blockContext)
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
	p.EnterRule(localctx, 12, BlunderbussParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(BlunderbussParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35475305791488) != 0 {
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case BlunderbussParserEFFECT:
			{
				p.SetState(71)
				p.Effect_block()
			}

		case BlunderbussParserIF, BlunderbussParserFOR, BlunderbussParserBREAK, BlunderbussParserNEXT, BlunderbussParserRETURN, BlunderbussParserSAFE, BlunderbussParserLAZY, BlunderbussParserTYPE, BlunderbussParserID:
			{
				p.SetState(72)
				p.Stmt()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
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
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	NUM() antlr.TerminalNode
	STRING() antlr.TerminalNode
	ID() antlr.TerminalNode
	Func_call() IFunc_callContext
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode

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

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserOR, 0)
}

func (s *ExprContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserEQUAL, 0)
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

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserDIV, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserMINUS, 0)
}

func (s *ExprContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLSHIFT, 0)
}

func (s *ExprContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRSHIFT, 0)
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, BlunderbussParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(81)
			p.Match(BlunderbussParserEXCL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.expr(9)
		}

	case 2:
		{
			p.SetState(83)
			p.Match(BlunderbussParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.expr(0)
		}
		{
			p.SetState(85)
			p.Match(BlunderbussParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(87)
			p.Match(BlunderbussParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(88)
			p.Match(BlunderbussParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(89)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(90)
			p.Func_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BlunderbussParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(94)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63680) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(95)
					p.expr(9)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BlunderbussParserRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(97)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == BlunderbussParserMULT || _la == BlunderbussParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(98)
					p.expr(8)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BlunderbussParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(100)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&780) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(101)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	Expr() IExprContext
	SEMI() antlr.TerminalNode
	LAZY() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Func_call() IFunc_callContext
	SAFE() antlr.TerminalNode
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

func (s *StmtContext) Expr() IExprContext {
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

func (s *StmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserSEMI, 0)
}

func (s *StmtContext) LAZY() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLAZY, 0)
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

func (s *StmtContext) SAFE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserSAFE, 0)
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
	p.EnterRule(localctx, 16, BlunderbussParserRULE_stmt)
	var _la int

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BlunderbussParserLAZY {
			{
				p.SetState(107)
				p.Match(BlunderbussParserLAZY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(110)
			p.Match(BlunderbussParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(BlunderbussParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.expr(0)
		}
		{
			p.SetState(114)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(BlunderbussParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(BlunderbussParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.expr(0)
		}
		{
			p.SetState(122)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.Match(BlunderbussParserRETURN)
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

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BlunderbussParserSAFE {
			{
				p.SetState(128)
				p.Match(BlunderbussParserSAFE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(131)
			p.Func_call()
		}
		{
			p.SetState(132)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(134)
			p.Match(BlunderbussParserNEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(136)
			p.Match(BlunderbussParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(BlunderbussParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(138)
			p.If_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(139)
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

// IEffect_blockContext is an interface to support dynamic dispatch.
type IEffect_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EFFECT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsEffect_blockContext differentiates from other interfaces.
	IsEffect_blockContext()
}

type Effect_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffect_blockContext() *Effect_blockContext {
	var p = new(Effect_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_effect_block
	return p
}

func InitEmptyEffect_blockContext(p *Effect_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BlunderbussParserRULE_effect_block
}

func (*Effect_blockContext) IsEffect_blockContext() {}

func NewEffect_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Effect_blockContext {
	var p = new(Effect_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BlunderbussParserRULE_effect_block

	return p
}

func (s *Effect_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Effect_blockContext) EFFECT() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserEFFECT, 0)
}

func (s *Effect_blockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserLBRACE, 0)
}

func (s *Effect_blockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserRBRACE, 0)
}

func (s *Effect_blockContext) AllStmt() []IStmtContext {
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

func (s *Effect_blockContext) Stmt(i int) IStmtContext {
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

func (s *Effect_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Effect_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Effect_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.EnterEffect_block(s)
	}
}

func (s *Effect_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BlunderbussListener); ok {
		listenerT.ExitEffect_block(s)
	}
}

func (s *Effect_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BlunderbussVisitor:
		return t.VisitEffect_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BlunderbussParser) Effect_block() (localctx IEffect_blockContext) {
	localctx = NewEffect_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BlunderbussParserRULE_effect_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(BlunderbussParserEFFECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(BlunderbussParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35471010824192) != 0 {
		{
			p.SetState(144)
			p.Stmt()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
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
		p.SetState(152)
		p.Match(BlunderbussParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.expr(0)
	}
	{
		p.SetState(155)
		p.Match(BlunderbussParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Block()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BlunderbussParserELSEIF {
		{
			p.SetState(157)
			p.Match(BlunderbussParserELSEIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(BlunderbussParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.expr(0)
		}
		{
			p.SetState(160)
			p.Match(BlunderbussParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Block()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserELSE {
		{
			p.SetState(168)
			p.Match(BlunderbussParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
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
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	TYPE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllASSIGN() []antlr.TerminalNode
	ASSIGN(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

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

func (s *For_stmtContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserSEMI)
}

func (s *For_stmtContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserSEMI, i)
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

func (s *For_stmtContext) TYPE() antlr.TerminalNode {
	return s.GetToken(BlunderbussParserTYPE, 0)
}

func (s *For_stmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserID)
}

func (s *For_stmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserID, i)
}

func (s *For_stmtContext) AllASSIGN() []antlr.TerminalNode {
	return s.GetTokens(BlunderbussParserASSIGN)
}

func (s *For_stmtContext) ASSIGN(i int) antlr.TerminalNode {
	return s.GetToken(BlunderbussParserASSIGN, i)
}

func (s *For_stmtContext) AllExpr() []IExprContext {
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

func (s *For_stmtContext) Expr(i int) IExprContext {
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(BlunderbussParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(BlunderbussParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserTYPE {
		{
			p.SetState(174)
			p.Match(BlunderbussParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(BlunderbussParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.expr(0)
		}

	}
	{
		p.SetState(180)
		p.Match(BlunderbussParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35184372679682) != 0 {
		{
			p.SetState(181)
			p.expr(0)
		}

	}
	{
		p.SetState(184)
		p.Match(BlunderbussParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BlunderbussParserID {
		{
			p.SetState(185)
			p.Match(BlunderbussParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(BlunderbussParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.expr(0)
		}

	}
	{
		p.SetState(190)
		p.Match(BlunderbussParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
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
	case 7:
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
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

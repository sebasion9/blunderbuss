package semantics

import (
	"fmt"
	"strings"
)

type ScopeTreeType int
const (
	FNCALL ScopeTreeType = iota
	FOR
	IF
	FUNC
	OTHER
)

type ScopeTree struct {
	name string
	children []*ScopeTree
	parent *ScopeTree
	vars map[string]ScopeItf
	scopeType ScopeTreeType
	currOff int
}

func (st *ScopeTree) IncrOff(incr int) {
	st.currOff += incr
}
func (st *ScopeTree) GetOff() int {
	return st.currOff
}

func (st *ScopeTree) GetChildren() []*ScopeTree {
	return st.children
}

func (st *ScopeTree) GetName() string {
	return st.name
}

func (st *ScopeTree) AddChildren(children *ScopeTree) {
	if st.children == nil {
		st.children = make([]*ScopeTree, 0)
	}
	st.children = append(st.children, children)
}

func (st *ScopeTree) GetParent() *ScopeTree {
	return st.parent
}

func (st *ScopeTree) GetVars() *map[string]ScopeItf {
	return &st.vars
}

func (st *ScopeTree) GetScopeByName(name string) *ScopeTree {
	for _, c := range st.GetChildren() {
		if c.name == name {
			return c
		}
		foundScope := c.GetScopeByName(name)
		if foundScope != nil {
			return foundScope
		}
	}
	return nil
}

func GetRootScope(scope *ScopeTree) *ScopeTree {
	parent := scope.GetParent()
	if parent == nil {
		return scope
	}
	return GetRootScope(parent)
}

func GetAllVarsDownLen(scope *ScopeTree) int {
	sum := 0
	if scope.scopeType != FNCALL && scope.scopeType != OTHER {
		sum = len(*scope.GetVars())
	}
	for _, c := range scope.GetChildren() {
		sum += GetAllVarsDownLen(c)
	}
	return sum
}



func SearchIdUp(scope *ScopeTree, id string) *ScopeTree {
	if scope == nil {
		return nil
	}
	vars := *scope.GetVars()
	if _, found := vars[id]; found {
		return scope
	}
	return SearchIdUp(scope.GetParent(), id)
}

func GetWhatFunc(scope *ScopeTree) *ScopeTree {
	if scope.scopeType == FUNC {
		return scope
	}
	return GetWhatFunc(scope.GetParent())
}

func GetWhatFor(scope *ScopeTree) *ScopeTree {
	if scope.scopeType == FOR {
		return scope
	}
	return GetWhatFor(scope.GetParent())
}


func (st *ScopeTree) DebugPrint(ident int) {
	i := strings.Repeat(" ", ident)
	fmt.Printf("%sSCOPE: %s\n,", i, st.name)
	fmt.Printf("%sVARS:", i)
	for i, v := range st.vars {
		fmt.Printf("%sNAME: %s", i, v.Raw())
	}
	for _, c := range st.GetChildren() {
		c.DebugPrint(ident + 1)
	}
}


func NewScopeTree(name string, parent *ScopeTree, scopeType ScopeTreeType) *ScopeTree {
	scope :=  &ScopeTree{
		name,
		nil,
		parent,
		map[string]ScopeItf{},
		scopeType,
		8,
	}
	if parent != nil {
		parent.AddChildren(scope)
	}
	return scope
}


type CompilerContext struct {
	scopes *ScopeTree
	currentScope *ScopeTree
}

//TODO: remember to set current scope when creating new
func NewCompilerContext() CompilerContext {
	scopes := NewScopeTree("program", nil, OTHER)
	return CompilerContext{
		scopes: scopes,
		currentScope: scopes,
	}
}


func (cc *CompilerContext) GetCurrScope() *ScopeTree {//map[string]ScopeItf {
	return cc.currentScope
}



type ScopeItf interface {
	Raw() any
	Type() Type_
	Offset() int
}

type ScopeFunc struct {
	id string
	args []ScopeFuncArg
	type_ Type_
}

func NewScopeFunc(id string, args []ScopeFuncArg, type_ Type_) *ScopeFunc{
	return &ScopeFunc{id, args, type_}
}

func(sf *ScopeFunc) Raw() any {
	return sf.id
}

func(sf *ScopeFunc) Type() Type_ {
	return sf.type_
}

func(sf *ScopeFunc) Offset() int {
	return 0
}

type ScopeVar struct {
	expr any
	type_ Type_
	offset int
}

func NewScopeVar(expr any, type_ Type_, offset int) *ScopeVar {
	return &ScopeVar{expr, type_, offset}
}


func(sf *ScopeVar) Raw() any {
	return sf.expr
}

func(sf *ScopeVar) Type() Type_ {
	return sf.type_
}

func(sf *ScopeVar) Offset() int {
	return sf.offset
}

type ScopeFuncArg struct {
	expr any
	type_ Type_
	idx int
}

func NewScopeFuncArg(expr any, type_ Type_, idx int) *ScopeFuncArg{
	return &ScopeFuncArg{expr, type_, idx}
}

func(sf *ScopeFuncArg) Raw() any {
	return sf.expr
}

func(sf *ScopeFuncArg) Type() Type_ {
	return sf.type_
}

func(sf *ScopeFuncArg) Offset() int {
	return 0
}


type Register struct {
	name string
	type_ Type_
}

func (r *Register) Write(type_ Type_) {
	r.type_ = type_

}

func InitRegisters(scope *map[string]ScopeItf) {
	(*scope)["rax"] = &Register{"rax", VOID_}
	(*scope)["rbx"] = &Register{"rbx", VOID_}
	(*scope)["rcx"] = &Register{"rcx", VOID_}
	(*scope)["rdx"] = &Register{"rdx", VOID_}
	(*scope)["rsi"] = &Register{"rsi", VOID_}
	(*scope)["rdi"] = &Register{"rdi", VOID_}
	(*scope)["rsp"] = &Register{"rsp", VOID_}
	(*scope)["rbp"] = &Register{"rbp", VOID_}
	(*scope)["r8"] = &Register{"r8", VOID_}
	(*scope)["r9"] = &Register{"r9", VOID_}
}

func(r *Register) Raw() any {
	return r.type_
}

func(r *Register) Type() Type_ {
	return r.type_
}

func(r *Register) Offset() int {
	return 0
}


var scopeId = 0
func BlockScopeName(name string) string {
	scopeName := fmt.Sprintf("BLOCK__SCOPE__%s__%d", name, scopeId)
	scopeId++
	return scopeName
}

func EndFnLabel(name string) string {
	return fmt.Sprintf("END__%s", name)
}

func StartBlockLabel(name string) string {
	return fmt.Sprintf("START__%s", name)
}

func EndStmtLabel(name string) string {
	return fmt.Sprintf("STMT__%s", name)
}

func EndBlockLabel(name string) string {
	return fmt.Sprintf("END__%s", name)
}


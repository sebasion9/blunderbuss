package semantics

import "fmt"


type CompilerContext struct {
	scopes map[string]map[string]ScopeItf
	currentScopeIdx string
}

func NewCompilerContext() CompilerContext {
	scopes := map[string]map[string]ScopeItf{}
	scopes["program"] = make(map[string]ScopeItf)
	return CompilerContext{
		scopes: scopes,
		currentScopeIdx: "program",
	}
}


func (cc *CompilerContext) GetCurrScope() map[string]ScopeItf{
	return cc.scopes[cc.currentScopeIdx]
}

func (cc *CompilerContext) GetScopeByName(name string) map[string]ScopeItf {
	return cc.scopes[name]
}

func (cc *CompilerContext) NewScope(scopeName string) map[string]ScopeItf {
	cc.scopes[scopeName] = make(map[string]ScopeItf)
	cc.currentScopeIdx = scopeName
	return cc.scopes[scopeName]
}

// if nil, then bad
// derefer until str, num or nil
func (cc *CompilerContext) DereferId(ID string) any {

	return ""
}

type ScopeItf interface {
	Raw() any
	Type() Type_
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

type ScopeVar struct {
	expr any
	type_ Type_
}

func NewScopeVar(expr any, type_ Type_) *ScopeVar {
	return &ScopeVar{expr, type_}
}

func(sf *ScopeVar) Raw() any {
	return sf.expr
}

func(sf *ScopeVar) Type() Type_ {
	return sf.type_
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

func PrefixScope(scope *map[string]ScopeItf, fnName string) {
	for k, v := range *scope {
		newKey := fmt.Sprintf("%s__%s", fnName, k)
		delete(*scope, k)
		(*scope)[newKey] = v
	}
}

// enter block, create fresh scope
// save variable here etc.
// look for variables in descending, so outer scopes
// on exit clear the scope
// scope should be smth like funcName+rand or idx (recursive fn calls)

// could add here func stack frames also
// separate struct for stack

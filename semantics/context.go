package semantics


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
}

type ScopeFunc struct {
	id string
	type_ Type_
}

func NewScopeFunc(id string, type_ Type_) *ScopeFunc{
	return &ScopeFunc{id, type_}
}

func(sf *ScopeFunc) Raw() any {
	return sf.id
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

// enter block, create fresh scope
// save variable here etc.
// look for variables in descending, so outer scopes
// on exit clear the scope
// scope should be smth like funcName+rand or idx (recursive fn calls)

// could add here func stack frames also
// separate struct for stack

package semantics


type CompilerContext struct {
	// scope
	// variables
	// deref ID's
	//   curentScope ID EXPR
	scopes []map[string]ScopeItf
	currentScopeIdx int
}

func NewCompilerContext() CompilerContext {
	scopes := []map[string]ScopeItf{}
	scopes = append(scopes, make(map[string]ScopeItf))
	return CompilerContext{
		scopes: scopes,
		currentScopeIdx: 0,
	}
}


func (cc *CompilerContext) GetCurrScope() *map[string]ScopeItf{
	return &cc.scopes[cc.currentScopeIdx]
}


func (cc *CompilerContext) NewScope() *map[string]ScopeItf{
	cc.currentScopeIdx++
	cc.scopes = append(cc.scopes, make(map[string]ScopeItf))
	return cc.GetCurrScope()
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
}

func NewScopeFunc(id string) *ScopeFunc{
	return &ScopeFunc{id}
}

func(sf *ScopeFunc) Raw() any {
	return sf.id
}

type ScopeVar struct {
	expr any
}

func NewScopeVar(expr any) *ScopeVar {
	return &ScopeVar{expr}
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

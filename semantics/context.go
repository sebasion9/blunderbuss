package semantics;


type CompilerContext struct {
	// scope
	// variables
	// deref ID's
	//   curentScope ID EXPR
	scopes []map[string]any
	currentScope int
}
func NewCompilerContext() CompilerContext {
	return CompilerContext {
		scopes: []map[string]any{},
		currentScope: 0,
	}
}

func (cc *CompilerContext) NewScope() {
	cc.currentScope++
}

// if nil, then bad
// derefer until str, num or nil
func (cc *CompilerContext) DereferId(ID string) any {

	return ""
}

// enter block, create fresh scope
// save variable here etc.
// look for variables in descending, so outer scopes
// on exit clear the scope
// scope should be smth like funcName+rand or idx (recursive fn calls)


// could add here func stack frames also
// separate struct for stack

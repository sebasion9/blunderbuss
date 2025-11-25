package semantics;

const (
	INT = "int"
	STR = "str"
	// DOUBLE = "double"
	// BOOL = "bool"
)

type Type_ int
const (
	INT_ Type_ = iota
	STR_
	VOID_
)
func TypeEnumFromStr(type_ string) Type_ {
	if type_ == INT {
		return INT_
	}
	if type_ == STR {
		return STR_
	}
	return VOID_
}

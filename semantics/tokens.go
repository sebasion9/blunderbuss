package semantics;

const (
	INT = "int"
	STR = "str"
	VOID = "void"
	ANY = "any"
	// DOUBLE = "double"
	// BOOL = "bool"
)

type Type_ int
const (
	INT_ Type_ = iota
	STR_
	VOID_
	ANY_
)
func TypeEnumFromStr(type_ string) Type_ {
	if type_ == INT {
		return INT_
	}
	if type_ == STR {
		return STR_
	}
	if type_ == VOID {
		return VOID_
	}
	return ANY_
}
func StrFromTypeEnum(type_  Type_) string {
	if type_ == INT_ {
		return INT
	}

	if type_ == STR_ {
		return STR
	}
	if type_ == VOID_ {
		return VOID
	}
	return ANY
}

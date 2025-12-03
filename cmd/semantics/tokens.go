package semantics;

const (
	INT = "int"
	STR = "str"
	ANY = "any"
	PTR = "ptr"
)

type Type_ int
const (
	INT_ Type_ = iota
	STR_
	ANY_
	PTR_
)
func TypeEnumFromStr(type_ string) Type_ {
	if type_ == INT {
		return INT_
	}
	if type_ == STR {
		return STR_
	}
	if type_ == PTR {
		return PTR_
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

	if type_ == PTR_ {
		return PTR
	}

	return ANY
}

func IntFromTypeEnum(type_ Type_) int {
	if type_ == INT_ {
		return 1
	}

	if type_ == STR_ {
		return 2
	}

	if type_ == PTR_ {
		return 3
	}
	return -1
}

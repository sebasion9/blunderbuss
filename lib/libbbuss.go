package lib

// extern necessary libc methods, based on tree at the beginning

type LibFuncDecl struct {
	Name string
	Type_ string
	Args []struct {
		Name string
		Type_ string
	}

}

func GetDeclaredFuncs() []LibFuncDecl {
	return []LibFuncDecl {{
			Name: "printf",
			Type_: "void",
			Args: []struct{Name string; Type_ string}{
				{Name: "fmt", Type_: "str"},
				{Name: "msg", Type_: "str"},
			},
		},
	}
}

func Print(arg any) {
	//TODO: this should pass arguments to asm func
	//TODO: this should be able to print integers too

}



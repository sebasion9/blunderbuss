package lib

// extern necessary libc methods, based on tree at the beginning

func GetDeclaredFuncs() [][2]string {
	return [][2]string {
		{"print", "void"},
	}
}

func Print(arg any) {
	//TODO: this should pass arguments to asm func

}



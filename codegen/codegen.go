package codegen
import "fmt"

type Codegen struct {

};

// func (c *Codegen) mov(dst string, src string) string {
// 	return ""
// }
func (c *Codegen) GenFuncInit(code *[]string, funcName string) {
	*code = append(*code,
		fmt.Sprintf("global %s", funcName),
		fmt.Sprintf("%s:", funcName),
	)
}

func (c *Codegen) GenFuncExit(code *[]string) {
	*code = append(*code, "ret")
}

func (c *Codegen) GenMainInit(code *[]string) {
	*code = append(*code, 
		"global _start",
		"_start:",
	)
}

func (c *Codegen) GenMainExit(code *[]string, returnCode int) {
	*code = append(*code,
		"mov rax, 60; exit syscall",
		fmt.Sprintf("mov rdi, %d", returnCode),
		"syscall",
	)
}



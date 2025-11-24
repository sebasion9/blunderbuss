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
		// "global _start",
		"global main",
		"extern puts",
		"main:",
	)
}

func (c *Codegen) GenMainExit(code *[]string, returnCode int) {
	// *code = append(*code,
	// 	"mov rax, 60; exit syscall",
	// 	fmt.Sprintf("mov rdi, %d", returnCode),
	// 	"syscall",
	// )
	*code = append(*code,
		fmt.Sprintf("mov rax, %d", returnCode),
		"ret",
	)

}

func (c *Codegen) GenStrPrimitive(code *[]string, name string, val string) {
	*code = append(*code,
		fmt.Sprintf("%s: db %s, 0", name, val),
	)

}

func (c *Codegen) GenIntPrimitive(code *[]string, name string, val int) {
	*code = append(*code,
		fmt.Sprintf("%s: dd %d", name, val),
	)
}

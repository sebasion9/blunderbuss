package codegen
import "fmt"

type Codegen struct {
	global []string
	text []string
	data []string

};

func (c *Codegen) GetGlobal() []string {
	return c.global
}

func (c *Codegen) GetText() []string {
	return c.text
}

func (c *Codegen) GetData() []string {
	return c.data
}

func (c *Codegen) GenInit() {
	c.global = append(c.global, "section .text")
	c.global = append(c.global, "extern puts")
	c.data = append(c.data, "section .data")
}

func (c *Codegen) ConcatAsm() []string{
	asm := append(c.data, c.global...)
	asm = append(asm, c.text...)
	return asm
}

// func (c *Codegen) mov(dst string, src string) string {
// 	return ""
// }
func (c *Codegen) GenFuncInit(funcName string) {
	//TODO:
	c.global = append(c.global, fmt.Sprintf("global %s", funcName))
	c.text= append(c.text,
		fmt.Sprintf("%s:", funcName),
		"push rbp",
		"mov rbp, rsp",
	)
}

func (c *Codegen) GenFuncExit(retVal any) {
	c.text = append(c.text,
		"pop rbp",
		"ret",
	)
}

// func (c *Codegen) GenMainInit(code *[]string) {
// 	*code = append(*code, 
// 		"global main",
// 		"extern puts",
// 		"main:",
// 	)
// }
//
// func (c *Codegen) GenMainExit(code *[]string, returnCode int) {
// 	*code = append(*code,
// 		fmt.Sprintf("mov rax, %d", returnCode),
// 		"ret",
// 	)
//
// }

func (c *Codegen) GenStrPrimitive(name string, val string) {
	c.data = append(c.data,
		fmt.Sprintf("%s: db %s, 0", name, val),
	)

}

func (c *Codegen) GenIntPrimitive(name string, val int) {
	c.data = append(c.data,
		fmt.Sprintf("%s: dd %d", name, val),
	)
}

func (c *Codegen) GenPrint(val string) {
	//c.GenFuncInit("print")
	c.text = append(c.text,
		fmt.Sprintf("mov rdi, %s", val),
		"call puts",
	)
	//c.GenFuncExit(nil)
}

var registers = [6]string{"rdi","rsi","rdx","rcx","r8","r9"}
func (c *Codegen) getArg(idx int) string {
	if idx < 6 {
		return registers[idx]
	}
	return fmt.Sprintf("[rbp + %d]", 8 + (idx - 6) * 8)
}


var generatedIdCount = 0
func (c *Codegen) CreateId() string {
	generatedIdCount++
	return fmt.Sprintf("__VAR__%d", generatedIdCount)
}



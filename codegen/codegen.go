package codegen

import (
	"fmt"
	"strings"
)

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
	//TODO: extern
	c.global = append(c.global, "extern printf")
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

func (c *Codegen) GenStrPrimitive(name string, val string) {
	c.data = append(c.data,
		fmt.Sprintf("%s: db %s, 0", name, val),
	)

}

func (c *Codegen) GenIntPrimitive(name string, val int) {
	c.data = append(c.data,
		fmt.Sprintf("%s: dq %d", name, val),
	)
}

func (c *Codegen) GenPrint() {
	//c.GenFuncInit("print")
	c.text = append(c.text,
		"mov rdi, rdi",
		"call puts",
	)
	//c.GenFuncExit(nil)
}

//TODO: remove
func (c *Codegen) TestArgs() {
	c.text = append(c.text, 
		"mov rax, 0",
		"add rax, [rdi]",
		"add rax, [rsi]",
		"add rax, [rdx]",
		"add rax, [rcx]",
		"add rax, [r8]",
		"add rax, [r9]",
		"add rax, [rbp + 16]",
	)

}

func (c *Codegen) GenPushArg(id string, idx int) {
	if idx > 5 {
		c.text = append(c.text, 
			fmt.Sprintf("mov rax, [rel %s]", id),
			"push rax",
		)
		return
	}
	c.text = append(c.text, 
		fmt.Sprintf("mov %s, %s", c.getArg(idx), id),
	)
}

func (c *Codegen) GenCallFunc(fnName string) {
	c.text = append(c.text, fmt.Sprintf("call %s", fnName))

}


func (c *Codegen) GenExpandStackFrame(dds int) {
	// if dds > 0 {
	// 	c.text = append(c.text, fmt.Sprintf("sub rsp, %d", dds * 8))
	// }
}

func (c *Codegen) GenShrinkStackFrame(dds int) {
	if dds > 0 {
		c.text = append(c.text, fmt.Sprintf("add rsp, %d", dds * 8))
	}
}



//TODO: smarter mov
// func (c *Codegen) GenMov(dst string, src string) {
// 	c.text = append(c.text, fmt.Sprintf("mov %s, %s", dst, src))
// }


// those 2 can be merged into one
// mov rax, rsi
// mov rax, addr

// mov rax, [rsi]
// mov [rel label], rax

func (c *Codegen) GenMovMemory(dst string, src string) {
	c.text = append(c.text, fmt.Sprintf("mov %s, %s", dst, src))
}

func (c *Codegen) GenMovIndirect(dst string, src string) {
	c.text = append(c.text, fmt.Sprintf("mov %s, [%s]", dst, src))
}


func (c *Codegen) GenMovRegRelative(dst string, src string) {
	c.text = append(c.text, fmt.Sprintf("mov %s, [rel %s]", dst, src))
}

func (c *Codegen) GenMovAddrRelative(dst string, src string) {
	c.text = append(c.text, fmt.Sprintf("mov [rel %s], %s", dst, src))
}

var registers = []string{"rdi","rsi","rdx","rcx","r8","r9"}
func (c *Codegen) getArg(idx int) string {
	// if idx < 6 {
		return registers[idx]
	// }
	// return fmt.Sprintf("[rbp + %d]", 16 + (idx - 6) * 8)
}


var generatedIdCount = 0
func (c *Codegen) CreateId() string {
	generatedIdCount++
	return fmt.Sprintf("__VAR__%d", generatedIdCount)
}

func (c *Codegen) CreateVarId(funcName string, varname string) string {
	return fmt.Sprintf("%s____%s", funcName, varname);
}



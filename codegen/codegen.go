package codegen

import (
	"fmt"
	"strconv"
)

type Instruction struct {
	op string
	dst *string
	src *string
}

func (i *Instruction) toString() string {
	instr := i.op
	if i.dst != nil {
		instr = fmt.Sprintf("%s %s", instr, *i.dst)
	}

	if i.src != nil {
		instr = fmt.Sprintf("%s, %s", instr, *i.src)
	}

	return instr
}

func instrArrToString(arr []Instruction) string {
	asm := ""
	for _, instr := range arr {
		asm = fmt.Sprintf("%s\n%s", asm, instr.toString())
	}
	return asm
}


func NewInstr(op string, dst *string, src *string) Instruction {
	return Instruction{op, dst, src}
}

type Codegen struct {
	global []Instruction
	text []Instruction
	data []Instruction
	CurrScope *string
};

func (c *Codegen) GetGlobal() []Instruction {
	return c.global
}

func (c *Codegen) GetText() [] Instruction {
	return c.text
}

func (c *Codegen) GetData() []Instruction {
	return c.data
}


func (c *Codegen) GenInit() {
	dst1 := ".text"
	c.global = append(c.global, NewInstr("section", &dst1, nil));
	//TODO: extern
	dst2 := "printf"
	c.global = append(c.global, NewInstr("extern", &dst2, nil));
	dst3 := ".data"
	c.data = append(c.data, NewInstr("section", &dst3, nil));
}

func (c *Codegen) StreamAsm() string{
	asm := fmt.Sprintf("%s\n%s\n%s", instrArrToString(c.data), instrArrToString(c.global), instrArrToString(c.text))
	return asm
}

func (c *Codegen) GenFuncInit(funcName string) {
	rbp := "rbp"
	rsp := "rsp"
	*c.CurrScope = funcName

	c.global = append(c.global, NewInstr("global", &funcName ,nil))
	c.text = append(c.text,
		NewInstr(fmt.Sprintf("%s:", funcName), nil, nil),
		NewInstr("push", &rbp, nil),
		NewInstr("mov", &rbp, &rsp),
	)
}

func (c *Codegen) GenFuncExit(retVal any) {
	rbp := "rbp"
	c.text = append(c.text,
		NewInstr("pop", &rbp ,nil),
		NewInstr("ret", nil, nil),
	)
}

func (c *Codegen) GenStrPrimitive(name string, val string) {
	name = fmt.Sprintf("%s__%s", *c.CurrScope, name)
	c.data = append(c.data,
		NewInstr(fmt.Sprintf("%s: db %s, 0", name, val), nil, nil),
	)

}

func (c *Codegen) GenIntPrimitive(name string, val int) {
	name = fmt.Sprintf("%s__%s", *c.CurrScope, name)
	c.data = append(c.data,
		NewInstr(fmt.Sprintf("%s: dq %d", name, val), nil, nil),
	)
}

// func (c *Codegen) GenPrint() {
// 	//c.GenFuncInit("print")
// 	rdi := "rdi"
// 	puts := "puts"
// 	c.text = append(c.text,
// 		NewInstr("mov", &rdi, &rdi),
// 		NewInstr("call", &puts, nil),
// 	)
// 	//c.GenFuncExit(nil)
// }

//TODO: remove
// func (c *Codegen) TestArgs() {
// 	c.text = append(c.text, 
// 		"mov rax, 0",
// 		"add rax, [rdi]",
// 		"add rax, [rsi]",
// 		"add rax, [rdx]",
// 		"add rax, [rcx]",
// 		"add rax, [r8]",
// 		"add rax, [r9]",
// 		"add rax, [rbp + 16]",
// 	)
//
// }

func (c *Codegen) GenPushArg(id string, idx int) {
	id = fmt.Sprintf("%s__%s", *c.CurrScope, id)
	if idx > 5 {
		rax := "rax"
		rel := fmt.Sprintf("[rel %s]", id)
		c.text = append(c.text, 
			NewInstr("mov", &rax, &rel),
			NewInstr("push", &rax, nil),
		)
		return
	}
	arg := c.getArg(idx)
	c.text = append(c.text, 
		NewInstr("mov", &arg, &id),
	)
}

func (c *Codegen) GenCallFunc(fnName string) {
	c.text = append(c.text, NewInstr("call", &fnName, nil))
}


func (c *Codegen) GenExpandStackFrame(dds int) {
	// if dds > 0 {
	// 	c.text = append(c.text, fmt.Sprintf("sub rsp, %d", dds * 8))
	// }
}

func (c *Codegen) GenShrinkStackFrame(dds int) {
	if dds > 0 {
		rsp := "rsp"
		offset := strconv.Itoa(dds * 8)
		c.text = append(c.text, NewInstr("add", &rsp, &offset))
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
	c.text = append(c.text, NewInstr("mov", &dst, &src))
}

func (c *Codegen) GenMovIndirect(dst string, src string) {
	src = fmt.Sprintf("[%s__%s]", *c.CurrScope, src)
	c.text = append(c.text, NewInstr("mov", &dst, &src))
}


func (c *Codegen) GenMovRegRelative(dst string, src string) {
	src = fmt.Sprintf("[rel %s__%s]", *c.CurrScope, src)
	c.text = append(c.text, NewInstr("mov", &dst, &src))
}

func (c *Codegen) GenMovAddrRelative(dst string, src string) {
	dst = fmt.Sprintf("[rel %s__%s]", *c.CurrScope, dst)
	c.text = append(c.text, NewInstr("mov", &dst, &src))
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



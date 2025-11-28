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
	name = fmt.Sprintf("%s", name)
	c.data = append(c.data,
		NewInstr(fmt.Sprintf("%s: db %s, 0", name, val), nil, nil),
	)

}

func (c *Codegen) GenIntPrimitive(name string, val int) {
	name = fmt.Sprintf("%s", name)
	c.data = append(c.data,
		NewInstr(fmt.Sprintf("%s: dq %d", name, val), nil, nil),
	)
}
func (c *Codegen) GenPush(val any) {
	switch v := val.(type) {
	case string:
		c.text = append(c.text, NewInstr("push", &v, nil))
	case int:
		asStr := strconv.Itoa(v)
		c.text = append(c.text, NewInstr("push", &asStr, nil))
	}
}

func (c *Codegen) GenAlignStack(dds int) {
	rsp := "rsp"
	offset := fmt.Sprintf("%d", dds * 8)
	c.text = append(c.text, NewInstr("add", &rsp, &offset))
}



func (c *Codegen) GenPushArg(offset int, idx int) {

	rel := fmt.Sprintf("[rbp - %d]", offset)
	if idx > 5 {
		rax := "rax"
		c.text = append(c.text, 
			NewInstr("mov", &rax, &rel),
			NewInstr("push", &rax, nil),
		)
		return
	}
	arg := c.getArg(idx)
	c.text = append(c.text, 
		NewInstr("mov", &arg, &rel),
	)
}

func (c *Codegen) GenCallFunc(fnName string) {
	c.text = append(c.text, NewInstr("call", &fnName, nil))
}



func (c *Codegen) GenShrinkStackFrame(dds int) {
	if dds > 0 {
		rsp := "rsp"
		offset := strconv.Itoa(dds * 8)
		c.text = append(c.text, NewInstr("add", &rsp, &offset))
	}
}




func (c *Codegen) GenMovMemory(dst string, src string) {
	c.text = append(c.text, NewInstr("mov", &dst, &src))
}

func (c *Codegen) GenMovIndirect(dst string, src string) {
	src = fmt.Sprintf("[%s]", src)
	c.text = append(c.text, NewInstr("mov", &dst, &src))
}


func (c *Codegen) GenMovRegRelative(dst string, off int) {
	src := fmt.Sprintf("[rbp - %d]", off)
	c.text = append(c.text, NewInstr("mov", &dst, &src))
}

func (c *Codegen) GenMovAddrRelative(off int, src any) {
	dst := fmt.Sprintf("[rbp - %d]", off)
	s, ok := src.(string)
	if ok {
		c.text = append(c.text, NewInstr("mov", &dst, &s))
	}
	i, ok := src.(int)
	if ok {
		s = strconv.Itoa(i)
		c.text = append(c.text, NewInstr("mov", &dst, &s))
	}
}


func (c *Codegen) GenAddRegRelative(dst string, off int) {
	src := fmt.Sprintf("[rbp - %d]", off)
	c.text = append(c.text, NewInstr("add", &dst, &src))
}


func (c *Codegen) GenAddAddrRelative(off int, src any) {
	dst := fmt.Sprintf("[rbp - %d]", off)
	s, ok := src.(string)
	if ok {
		c.text = append(c.text, NewInstr("add", &dst, &s))
	}
	i, ok := src.(int)
	if ok {
		s = strconv.Itoa(i)
		c.text = append(c.text, NewInstr("add", &dst, &s))
	}
}

func (c *Codegen) GenImul(off int) {
	dst := fmt.Sprintf("[rbp - %d]", off)
	c.text = append(c.text, NewInstr("imul qword", &dst, nil))
}





var registers = []string{"rdi","rsi","rdx","rcx","r8","r9"}
func (c *Codegen) getArg(idx int) string {
	return registers[idx]
}


var generatedIdCount = 0
func (c *Codegen) CreateId() string {
	generatedIdCount++
	return fmt.Sprintf("__VAR__%d", generatedIdCount)
}



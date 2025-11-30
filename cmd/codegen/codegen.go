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

func (i *Instruction) SetDst(dst string) {
	i.dst = &dst
}

func (i *Instruction) SetSrc(src string) {
	i.src = &src
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

func (c *Codegen) GetGlobal() *[]Instruction {
	return &c.global
}

func (c *Codegen) GetText() *[]Instruction {
	return &c.text
}

func (c *Codegen) GetData() *[]Instruction {
	return &c.data
}


func (c *Codegen) GenInit() {
	dst1 := ".text"
	c.global = append(c.global, NewInstr("section", &dst1, nil));
	dst3 := ".data"
	c.data = append(c.data, NewInstr("section", &dst3, nil));
}

func (c *Codegen) GenExtern(name string) {
	c.global = append(c.global, NewInstr("extern", &name, nil));
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
		NewInstr("sub", &rsp, nil),
	)
}

func (c *Codegen) GenFuncExit() {
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

func (c *Codegen) GenAlignStack(dds int) {
	rsp := "rsp"
	offset := fmt.Sprintf("%d", dds)
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
	arg := c.GetCallArg(idx)
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

func (c *Codegen) GenMovOffset(dst string, src string, off int) {
	src = fmt.Sprintf("[%s * %d]", src, off)
	c.text = append(c.text, NewInstr("lea", &dst, &src))

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

func (c *Codegen) GenAddMemory(dst string, src string) {
	c.text = append(c.text, NewInstr("add", &dst, &src))
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

func (c *Codegen) GenSubMemory(dst string, src string) {
	c.text = append(c.text, NewInstr("sub", &dst, &src))
}

func (c *Codegen) GenSubRegRelative(dst string, off int) {
	src := fmt.Sprintf("[rbp - %d]", off)
	c.text = append(c.text, NewInstr("sub", &dst, &src))
}


func (c *Codegen) GenSubAddrRelative(off int, src any) {
	dst := fmt.Sprintf("[rbp - %d]", off)
	s, ok := src.(string)
	if ok {
		c.text = append(c.text, NewInstr("sub", &dst, &s))
	}
	i, ok := src.(int)
	if ok {
		s = strconv.Itoa(i)
		c.text = append(c.text, NewInstr("sub", &dst, &s))
	}
}

func (c *Codegen) GenImul(off int) {
	dst := fmt.Sprintf("[rbp - %d]", off)
	c.text = append(c.text, NewInstr("imul qword", &dst, nil))
}

func (c *Codegen) GenXor(dst string, src string) {
	c.text = append(c.text, NewInstr("xor", &dst, &src))
}

func (c *Codegen) GenDiv() {
	rbx := "rbx"
	c.text = append(c.text, 
		NewInstr("cqo", nil, nil),
		NewInstr("idiv", &rbx, nil),
	)
}

func (c *Codegen) GenCmp(src string, dst string) {
	c.text = append(c.text, NewInstr("cmp", &dst, &src))
}

func (c *Codegen) GenCmpRegAddr(dst string, off int) {
	src := fmt.Sprintf("[rbp - %d]", off)
	c.text = append(c.text, NewInstr("cmp", &dst, &src))
}

func (c *Codegen) GenCmpAddrReg(off int, src string) {
	dst := fmt.Sprintf("[rbp - %d]", off)
	c.text = append(c.text, NewInstr("cmp", &dst, &src))
}

func (c *Codegen) GenSete(dst string) {
	c.text = append(c.text, NewInstr("sete", &dst, nil))
}

func (c *Codegen) GenSetne(dst string) {
	c.text = append(c.text, NewInstr("setne", &dst, nil))
}

func (c *Codegen) GenSetl(dst string) {
	c.text = append(c.text, NewInstr("setl", &dst, nil))
}

func (c *Codegen) GenSetle(dst string) {
	c.text = append(c.text, NewInstr("setle", &dst, nil))
}

func (c *Codegen) GenSetg(dst string) {
	c.text = append(c.text, NewInstr("setg", &dst, nil))
}

func (c *Codegen) GenSetge(dst string) {
	c.text = append(c.text, NewInstr("setge", &dst, nil))
}

func (c *Codegen) GenJz(dst string) {
	c.text = append(c.text, NewInstr("jz", &dst, nil))
}

func (c *Codegen) GenJmp(dst string) {
	c.text = append(c.text, NewInstr("jmp", &dst, nil))
}

func (c *Codegen) GenLabel(name string) {
	label := fmt.Sprintf("%s:", name)
	c.text = append(c.text, NewInstr(label, nil, nil))
}


var registers = []string{"rdi","rsi","rdx","rcx","r8","r9"}
func (c *Codegen) GetCallArg(idx int) string {
	if idx > 6 {
		return fmt.Sprintf("[rbp + %d]", (idx - 6)*8+16)
	}
	return registers[idx]
}


var generatedIdCount = 0
func (c *Codegen) CreateId() string {
	generatedIdCount++
	return fmt.Sprintf("__VAR__%d", generatedIdCount)
}




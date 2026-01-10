package semantics

import (
	"blunderbuss/cmd/codegen"
	"strconv"
	"fmt"
)

func PrepKey(funcName string, retType Type_, args []ScopeVar, cg *codegen.Codegen, scope *ScopeTree) {
	// args + fn addr + ret val
	keyPartsLen := strconv.Itoa((len(args) + 2))
	rax := "rax"
	rbx := "rbx"
	rdx := "rdx"
	rdi := "rdi"
	rsi := "rsi"
	zero := "qword 0"
	r12 := "r12"
	r13 := "r13"
	raxWrapped := cg.WrapEff("rax")
	rdxScaled := cg.WrapEff("rdx*8")
	malloc := "malloc"


	// mov r13, type
	intType := strconv.Itoa(IntFromTypeEnum(retType))

	cg.AddText(codegen.NewInstr("mov", &r13, &intType))

	// mov rdx, len
	cg.AddText(codegen.NewInstr("mov", &rdx, &keyPartsLen))
	// VALUES ARR
	// lea rdi, rdx*8
	cg.AddText(codegen.NewInstr("lea", &rdi, &rdxScaled))
	// call malloc
	cg.AddText(codegen.NewInstr("call", &malloc, nil))

	off1 := scope.GetOff()
	scope.IncrOff(QWORD)
	rbpOff := fmt.Sprintf("[rbp - %d]", off1)
	cg.AddText(codegen.NewInstr("mov", &rbpOff, &rax))

	// values[0] = fn address
	// lea rbx, [fname]
	cg.AddText(codegen.NewInstr("lea", &rbx, &funcName))
	// mov [rax], rbx
	cg.AddText(codegen.NewInstr("mov", &raxWrapped, &rbx))


	// values[i+1] = args[i].offset
	// for each args in scope, generate movs
	for i, arg := range args {
		off := arg.offset
		offEff := cg.WrapEff(fmt.Sprintf("rbp - %d", off))
		// mov rbx, [args[i].off]
		cg.AddText(codegen.NewInstr("mov", &rbx, &offEff))
		// mov [rax + (i+2)*8], rbx
		raxOff := cg.WrapEff(fmt.Sprintf("rax + %d * 8", i+1))
		cg.AddText(codegen.NewInstr("mov", &raxOff, &rbx))
	}


	// mov ret val and ret type
	raxOff := cg.WrapEff(fmt.Sprintf("rax + %d * 8", len(args) + 1))
	cg.AddText(codegen.NewInstr("mov", &raxOff, &r12))


	// mov rdx, len
	cg.AddText(codegen.NewInstr("mov", &rdx, &keyPartsLen))

	// TYPES ARR
	// lea rdi, rdx*8
	cg.AddText(codegen.NewInstr("lea", &rdi, &rdxScaled))
	// call malloc
	cg.AddText(codegen.NewInstr("call", &malloc, nil))

	// always zero for fn type
	// mov [rax], 0
	cg.AddText(codegen.NewInstr("mov", &raxWrapped, &zero))


	// types[i+1] = args[i].type
	// for each args in scope, generate movs

	for i, arg := range args {
		type_ := strconv.Itoa(IntFromTypeEnum(arg.type_))
		// mov rbx, [args[i].type (1, 2, 3)]
		cg.AddText(codegen.NewInstr("mov", &rbx, &type_))
		// mov [rax + (i+2)*8], rbx
		raxOff := cg.WrapEff(fmt.Sprintf("rax + %d * 8", i+1))
		cg.AddText(codegen.NewInstr("mov", &raxOff, &rbx))
	}


	// mov ret val and ret type
	raxOff = cg.WrapEff(fmt.Sprintf("rax + %d * 8", len(args) + 1))
	cg.AddText(codegen.NewInstr("mov", &raxOff, &r13))

	// prep for SetM
	// rdi = values, rsi = types, rdx = size
	// mov rdi, rcx
	// cg.AddText(codegen.NewInstr("mov", &rdi, &rcx))
	cg.AddText(codegen.NewInstr("mov", &rdi, &rbpOff))
	// mov rsi, rax
	cg.AddText(codegen.NewInstr("mov", &rsi, &rax))
	// mov rdx, len
	cg.AddText(codegen.NewInstr("mov", &rdx, &keyPartsLen))
	// call SetM
	// cg.AddText(codegen.NewInstr("call", &setm, nil))
}

func CallSetm(cg *codegen.Codegen) {
	setm := "____SetM"
	cg.AddText(codegen.NewInstr("call", &setm, nil))
}

func CallGetm(cg *codegen.Codegen, endLabel string, startLabel string) {
	getm := "____GetM"
	rax := "rax"
	r12 := "r12"
	zero := "0"
	raxAddr := "[rax]"
	cg.AddText(codegen.NewInstr("call", &getm, nil))
	cg.AddText(codegen.NewInstr("cmp", &rax, &zero))
	// jz rip + 16
	cg.AddText(codegen.NewInstr("jz", &startLabel, nil))
	// mov rax, [rax]
	// cg.AddText(codegen.NewInstr("mov", &rax, &raxAddr))
	cg.AddText(codegen.NewInstr("mov", &r12, &raxAddr))
	// cg.AddText(codegen.NewInstr("mov", &r12, &rax))
	// jmp endLabel
	cg.AddText(codegen.NewInstr("jne", &endLabel, nil))
}


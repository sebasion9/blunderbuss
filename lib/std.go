//go:build cgo
// +build cgo

package main

/*
#include <stdint.h>
*/

import "C"
import (
	"strconv"
	"sync"
	"unsafe"
)

var m = make(map[uint64]unsafe.Pointer)
var mu sync.Mutex

//export Itoa
func Itoa(num C.long) *C.char {
	s := strconv.FormatInt(int64(num), 10)
	return C.CString(s)
}

//export Test
func Test() *C.char {
	return C.CString("test")
}

type CType int
const (
	INT CType = 1
	STR CType = 2
	PTR CType = 3
)

func Hash(fn unsafe.Pointer, args ...any) uint64 {
	h := uint64(uintptr(fn))
	for i := range args {
		switch v := args[i].(type) {
		case int64:
			h = h*31 + uint64(v)
		case string:
			for i := 0; i < len(v); i++ {
				h = h*31 + uint64(v[i])
			}
		default:
			panic("unsupported type")
		}
	}
	return h
}


//export ____SetM
func ____SetM(values *C.long, types *C.long, length C.long) {
	var goArgs []any
	var i int64
	cptr := C.malloc(8)
	for i = 0; i < int64(length); i ++ {
		val := *(*C.long)(unsafe.Pointer(uintptr(unsafe.Pointer(values)) + uintptr(i)*unsafe.Sizeof(*values)))
		t := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(types)) + uintptr(i)*unsafe.Sizeof(*types)))

		switch t {
		case 1:
			if i == int64(length) - 1 {
				*(*int64)(cptr) = int64(val)
			}
			goArgs = append(goArgs, int64(val))
		case 2:
			str := C.GoString((*C.char)(unsafe.Pointer(uintptr(val))))
			goArgs = append(goArgs, str)
			if i == int64(length) - 1 {
				cstr := C.CString(str)
				*(*uintptr)(cptr) = uintptr(unsafe.Pointer(cstr))
			}
		case 3:
			ptr := unsafe.Pointer(uintptr(val))
			goArgs = append(goArgs, ptr)
		case 0:
			ptr := unsafe.Pointer(uintptr(val))
			goArgs = append(goArgs, ptr)
		default:
		}
	}
	var key uint64 
	if len(goArgs) == 1 {
		key = Hash(goArgs[0].(unsafe.Pointer))
	} else {
		key = Hash(goArgs[0].(unsafe.Pointer), goArgs[1:len(goArgs)-1]...)
	}


	m[key] = cptr
}

//export ____GetM
func ____GetM(values *C.long, types *C.long, length C.long) unsafe.Pointer {
	var goArgs []any
	var i int64
	for i = 0; i < int64(length); i ++ {
		val := *(*C.long)(unsafe.Pointer(uintptr(unsafe.Pointer(values)) + uintptr(i)*unsafe.Sizeof(*values)))
		t := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(types)) + uintptr(i)*unsafe.Sizeof(*types)))

		switch t {
		case 1:
			goArgs = append(goArgs, int64(val))
		case 2:
			str := C.GoString((*C.char)(unsafe.Pointer(uintptr(val))))
			goArgs = append(goArgs, str)
		case 3:
			ptr := unsafe.Pointer(uintptr(val))
			goArgs = append(goArgs, ptr)
		case 0:
			ptr := unsafe.Pointer(uintptr(val))
			goArgs = append(goArgs, ptr)
		default:
		}
	}

	var key uint64 
	if len(goArgs) == 1 {
		key = Hash(goArgs[0].(unsafe.Pointer))
	} else {
		key = Hash(goArgs[0].(unsafe.Pointer), goArgs[1:len(goArgs)-1]...)
	}

	if val, ok := m[key]; ok {
		return val
	}
	return nil

}

func main() {}

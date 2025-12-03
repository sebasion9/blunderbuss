// +build cgo
package main


import "C"
import "sync"
import "strconv"
import "unsafe"

var m = map[uint64]unsafe.Pointer{}
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
	for i, _ := range args {
		if i == 0 {
			continue
		}
		switch v := args[i].(type) {
		case int:
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

//export SetM
func SetM(values *C.long, types *C.long, length C.long) {
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
	key := Hash(goArgs[0].(unsafe.Pointer), goArgs)

	num := new(int)
	*num = 101
	mu.Lock()
	m[key] = unsafe.Pointer(num)
	mu.Unlock()
}

//export GetM
func GetM(values *C.long, types *C.long, length C.long) unsafe.Pointer {
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

	key := Hash(goArgs[0].(unsafe.Pointer), goArgs)
	mu.Lock()
	defer mu.Unlock()

	if val, ok := m[key]; ok {
		cptr := C.malloc(C.size_t(8))
		*(*unsafe.Pointer)(cptr) = val
		return cptr
	}
	return nil
}

func main() {}

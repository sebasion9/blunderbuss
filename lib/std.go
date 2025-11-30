// +build cgo
package main


import "C"
import "sync"
import "strconv"

var m = map[string]int{}
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

//export SetM
func SetM(key *C.char, val C.long) {
	goKey := C.GoString(key)
	mu.Lock()
	m[goKey] = int(val)
	mu.Unlock()
}

//export GetM
func GetM(key *C.char) C.long {
	goKey := C.GoString(key)
	mu.Lock()
	defer mu.Unlock()
	return C.long(m[goKey])
}

func main() {}

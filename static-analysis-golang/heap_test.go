package data

import (
	"../pkg_asm"
	"fmt"
	"runtime"
	"testing"
)

func Test_heap(t *testing.T) {
	h := NewHeap()
	h.Add(1)
	h.Add(2)
	h.Add(9)
	h.Add(4)
	h.Add(5)
	h.Add(10)
	h.DelMax()
	h.DelMax()
	h.DelMax()
}

func Test_asm(t *testing.T) {
	println(pkg_asm.Id)
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Println(runtime.FuncForPC(pc).Name(), file, line)
	}
}
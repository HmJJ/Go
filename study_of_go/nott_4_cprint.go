package main

/*
#include <stdio.h>
*/
import "C"
import "unsafe"

/*
wtf 什么都没发生
*/
func main() {
	cstr := C.CString("Hello World !")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
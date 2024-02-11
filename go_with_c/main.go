package main

/*
#cgo LDFLAGS: -L. -lexample
#include "wrapper.h"
*/
import "C"
import "fmt"

func main() {
	C.hello()
	C.showName()
	fmt.Println(C.addNum(2, 3))
	fmt.Println("Go with program")
}

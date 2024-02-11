package main

/*
#cgo LDFLAGS: -L. -lexample
#include "wrapper.h"
*/
import "C"

func main() {
	C.hello()
	C.showName()
}

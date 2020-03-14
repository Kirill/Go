package main

import (
	"fmt"
	"unsafe"
)

type x struct {
  a bool     // 1
  b bool     // 1
  c uint64   // 8
}


func main() {
	fmt.Println("a size: ", unsafe.Sizeof(x{}.a))
	fmt.Println("b size: ", unsafe.Sizeof(x{}.b))
	fmt.Println("c size: ", unsafe.Sizeof(x{}.c))
	fmt.Println("x size: ", unsafe.Sizeof(x{}))
	fmt.Println()
	
	fmt.Println(unsafe.Offsetof(x{}.a))
	fmt.Println(unsafe.Offsetof(x{}.b))
	fmt.Println(unsafe.Offsetof(x{}.c))
}

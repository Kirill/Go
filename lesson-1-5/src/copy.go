package main

import (
	"fmt"
)

func f1() {
	a := struct{ x, y int }{0, 0}
	b := a
	a.x = 1
	fmt.Println(b.x) // ?
}

func f2() {
	a := new(struct{ x, y int })
	b := a
	a.x = 1
	fmt.Println(b.x) // ?
}

func f3() {
	a := struct{ x *int }{new(int)}
	b := a
	*a.x = 1
	fmt.Println(*b.x) // ?
}

func main() {
	f1()
}

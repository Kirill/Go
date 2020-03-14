package main

import (
	"fmt"
)

type User struct {
	Age int
}

func (u User) callValue() {
	u.Age++
}

func (u *User) callPointer() {
	u.Age++
}


func main() {
	u := User{23}

	u.callValue()
	fmt.Println(u.Age)

	u.callPointer()
	fmt.Println(u.Age)

	//User{}.callValue()
	//User{}.callPointer()
}

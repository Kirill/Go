package main

import (
	"fmt"
)

type Address struct {
	City   string
	Street string
}

func (a Address) String() string {
	return a.City + ", " + a.Street
}

type Organization struct {
	Address
	Name  string
	Phone string
}

func main() {
	org := Organization{Address: Address{City: "Spb", Street: "Nevski"}}
	fmt.Println(org.String())
	fmt.Println(org.Address.String())
}

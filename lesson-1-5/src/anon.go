package main

import (
	"fmt"
)

type testCases []struct {
	name     string
	input    string
	expected string
}

type Resp struct {
	Ok        bool
	Documents []struct {
		Id    int
		Title string
	}
}

func main() {
	cases := testCases{
		{name: "test1", input: "StR", expected: "str"},
		{name: "empty"},
	}
	fmt.Printf("%+v\n", cases)

	r := Resp{
		Ok: true,
		Documents: []struct {
			Id    int
			Title string
		}{{0, "zero"}, {1, "one"}},
	}
	fmt.Printf("%+v\n", r)
}

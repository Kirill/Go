package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("go-go-go")
			ch <- struct{}{}
		}()
	}

	for i := 0; i < 5; i++ {
		<-ch
	}
}

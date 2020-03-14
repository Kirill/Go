package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	v := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			v = v + 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

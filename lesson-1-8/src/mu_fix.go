package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	v := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			v = v + 1
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

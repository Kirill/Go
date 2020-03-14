package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("go-go-go")
		}()
	}
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			fmt.Println("go-go-go")
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

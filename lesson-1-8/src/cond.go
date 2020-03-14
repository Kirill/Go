package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var cond *sync.Cond
var tasks []func()

func worker() {
	var task func()

	mu.Lock()
	for len(tasks) == 0 {
		cond.Wait()
	}
	task, tasks = tasks[0], tasks[1:]
	mu.Unlock()

	task()
}

func produce(task func()) {
	mu.Lock()
	tasks = append(tasks, task)
	mu.Unlock()

	cond.Broadcast()
}

func main() {
	cond = sync.NewCond(&mu)

	for i := 0; i < 5; i++ {
		go worker()
	}

	produce(func() { fmt.Println("1") })
	produce(func() { fmt.Println("2") })
	produce(func() { fmt.Println("3") })

	time.Sleep(time.Second)
}

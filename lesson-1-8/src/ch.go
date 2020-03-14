package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(jobsNum int, ch chan string) {
	for i := 0; i < jobsNum; i++ {
		ch <- fmt.Sprintf("job %d", i)
	}
	close(ch)
}

func worker(ch chan string) {
	for job := range ch {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println(job, " done")

	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	go producer(10, ch)
	for i := 0; i < 3; i++ {
		go worker(ch)
	}
	time.Sleep(time.Second)
}

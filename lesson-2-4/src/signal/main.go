package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signalHandler(c <-chan os.Signal) {
	for s := range c {
		fmt.Println("Got signal:", s)
	}
}

func main() {
	fmt.Println("pid:", os.Getpid())
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)
	signal.Ignore(syscall.SIGTERM)
	go signalHandler(c)

	for {
		time.Sleep(time.Second)
	}
}

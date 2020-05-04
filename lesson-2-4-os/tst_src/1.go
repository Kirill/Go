package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	defaultPort := 9090
	defaultPortStr, ok := os.LookupEnv("DEFAULT_PORT")
	if ok {
		var err error
		defaultPort, err = strconv.Atoi(defaultPortStr)
		if err != nil {
			log.Fatalf("failed to parse default port", err)
		}
	}

	port := flag.Int("port", defaultPort, "")
	flag.Parse()

	fmt.Println(*port)
}

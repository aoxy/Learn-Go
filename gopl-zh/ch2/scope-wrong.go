package main

import (
	"log"
	"os"
)

var cwd string

func main() {
	test()
	log.Printf("Working directory = %s", cwd)
}

func test() {
	cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

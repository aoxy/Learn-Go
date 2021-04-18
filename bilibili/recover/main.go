package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello!")
		time.Sleep(100 * time.Millisecond)
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误，恐慌😰")
		}
	}()
	var myMap map[string]int
	myMap["abc"] = 123
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main(), i = ", i)
		time.Sleep(150 * time.Millisecond)
	}
}

package main

import "fmt"

func main() {
	ch := make(chan int, 2) //改为1会死锁
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

package main

import "fmt"

func main() {
	fmt.Println("计数...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("完成")
}

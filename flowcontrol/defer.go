package main

import "fmt"

func main() {
	defer fmt.Println("世界")
	fmt.Println("你好")
}

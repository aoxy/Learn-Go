package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("你好", "世界")
	// var a, b string = swap("你好", "世界")
	// var a, b  = swap("你好", "世界")
	fmt.Println(a, b)
}

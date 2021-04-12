package main

import "fmt"

const Pi = 3.1415 //不能用 :=

func main() {
	const World = "世界咋瓦鲁多"
	fmt.Println("你好", World)
	fmt.Println("开心", Pi, "天")
	const Truth = true
	fmt.Println("Go 语法？", Truth)
}

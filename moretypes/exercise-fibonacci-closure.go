package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	g, h := 0, 1
	return func() int {
		ret := g
		h = h + g
		g = h - g
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

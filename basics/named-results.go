package main

import "fmt"

func sqlit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(sqlit(17))
}

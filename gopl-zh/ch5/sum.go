package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(3, 4, 5, 6))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

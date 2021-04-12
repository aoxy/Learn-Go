package main

import "fmt"

func main() {
	i, j := 42, 2701

	var q *int = &j
	fmt.Println(*q)

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

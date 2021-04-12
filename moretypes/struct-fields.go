package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{12, 34}
	v.X = 4
	fmt.Println(v.X)
}

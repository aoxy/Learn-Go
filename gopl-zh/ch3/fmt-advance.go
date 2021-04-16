package main

import (
	"fmt"
	"math"
)

const (
	pi  = 3.1415926
	num = 123
)

func main() {
	fmt.Printf("Π^2 = %8.3f\n", math.Pow(float64(pi), 2.0))
	fmt.Printf("%[1]x, %[1]o, %#[1]b", num)
}

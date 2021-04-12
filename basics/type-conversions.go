package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var f float64 = math.Sqrt(x*x + y*y)
	// cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, ": ", z)
	}
	return z
}

func SqrtCount(x float64) float64 {
	count := 1
	z := 1.0
	prez := z
	z -= (z*z - x) / (2 * z)
	fmt.Println(count, ": ", z)
	for prez-z > 1e-6 || z-prez > 1e-6 {
		prez = z
		z -= (z*z - x) / (2 * z)
		count++
		fmt.Println(count, ": ", z)
	}
	return z
}

func main() {
	fmt.Println(SqrtCount(2))
	fmt.Println(math.Sqrt(2))
}

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 这样不行，不能为内建类型定义接收者方法
// func (f float64) Abs2() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

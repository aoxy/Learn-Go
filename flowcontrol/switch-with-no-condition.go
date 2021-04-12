package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 5:
		fmt.Println("凌晨好！")
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("中午好！")
	default:
		fmt.Println("晚上！")
	}
}

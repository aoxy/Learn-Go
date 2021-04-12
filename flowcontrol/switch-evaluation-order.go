package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("哪天周六？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("是今天。")
	case today + 1:
		fmt.Println("是明天。")
	case today + 2:
		fmt.Println("是后天。")
	default:
		fmt.Println("是更远的以后。")
	}
}

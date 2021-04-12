package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go运行在")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS S。")
	case "liunx":
		fmt.Println("Linux。")
	default:
		fmt.Printf("%s。\n", os)
	}
}

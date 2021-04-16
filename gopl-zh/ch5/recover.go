package main

import (
	"fmt"
	"strconv"
)

type SyntaxType int

func parse(input string) (s SyntaxType, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("内部错误：%v", p)
		}
	}()

	if input == "" {
		v, e := strconv.Atoi(input)
		panic("就是panic")
		return SyntaxType(0 / v), e
	}
	return 1, nil
}

func main() {
	s, e := parse("123")
	fmt.Println(s, e)
	s, e = parse("")
	fmt.Println(s, e)
}

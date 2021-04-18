package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 123
	reflectDemo(num)
	reflectDemoUpdate(&num)
	fmt.Println()
	fmt.Println("num = ", num)
}

func reflectDemo(b interface{}) {
	rType := reflect.TypeOf(b) //reflect.Type
	fmt.Println("rType = ", rType)

	rValue := reflect.ValueOf(b) //reflect.Value
	fmt.Println("rValue = ", rValue)

	fmt.Println("kind = ", rType.Kind())
	fmt.Println("kind = ", rValue.Kind())

	n2 := 321 + rValue.Int()
	fmt.Println("n2 = ", n2)

	iValue := rValue.Interface()

	num2 := iValue.(int)
	fmt.Println("num2 = ", num2)
}

func reflectDemoUpdate(b interface{}) {
	rType := reflect.TypeOf(b) //reflect.Type
	fmt.Println("rType = ", rType)

	rValue := reflect.ValueOf(b) //reflect.Value
	fmt.Println("rValue = ", rValue)

	fmt.Println("kind = ", rType.Kind())
	rValue.Elem().SetInt(666) // Elem()重要，b是地址（指针）
}

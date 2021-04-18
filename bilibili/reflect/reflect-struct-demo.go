package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var stu Student = Student{"Tom", 19}
	reflectDemo(stu)
}

func reflectDemo(b interface{}) {
	rType := reflect.TypeOf(b) //reflect.Type
	fmt.Println("rType = ", rType)

	rValue := reflect.ValueOf(b) //reflect.Value
	fmt.Println("rValue = ", rValue)

	fmt.Println("kind = ", rType.Kind())
	fmt.Println("kind = ", rValue.Kind())

	iValue := rValue.Interface()
	fmt.Printf("iValue = %v (%T)\n", iValue, iValue)
	//fmt.Printf("Name = \n", iValue.Name) //无法编译

	stu2 := iValue.(Student)
	fmt.Println("stu2 = ", stu2)
}

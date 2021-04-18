package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s Student = Student{
		ID:   "T12345678",
		Name: "李华",
		GPA:  1.9,
	}
	TestStruct(s)
}

type Student struct {
	ID   string `json:"id"`
	Name string `json:"stu_name"`
	GPA  float64
	Sex  string `hsp:"🚹🚺"`
}

func (s Student) Print() {
	fmt.Println("-------------start-------------")
	fmt.Println(s)
	fmt.Println("--------------end--------------")
}

func (s Student) GetSum(n1, n2 int) (int, int) {
	return n1 + n2, 1235
}

func (s Student) Set(id string, name string, gpa float64, sex string) {
	s.ID = id
	s.Name = name
	s.GPA = gpa
	s.Sex = sex
}

func TestStruct(b interface{}) {
	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	kd := rValue.Kind()
	if kd != reflect.Struct {
		fmt.Println("不是struct")
		return
	}

	num := rValue.NumField()
	fmt.Printf("这个struct有%d个字段\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("第%d个字段值为 %v\n", i, rValue.Field(i))
		tagValue := rType.Field(i).Tag.Get("json")
		tagValue2 := rType.Field(i).Tag.Get("hsp")
		if tagValue != "" {
			fmt.Printf("第%d个字段json标签为 %v\n", i, tagValue)
		}
		if tagValue2 != "" {
			fmt.Printf("第%d个字段hsp标签为 %v\n", i, tagValue2)
		}
	}

	numOfMethod := rValue.NumMethod()
	fmt.Printf("这个struct有%d个方法\n", numOfMethod)

	rValue.Method(1).Call(nil) //调用Print()方法，方法编号按函数名编号

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rValue.Method(0).Call(params)
	fmt.Println("res = ", res[0].Int())
	fmt.Println("res[1] = ", res[1].Int())
}

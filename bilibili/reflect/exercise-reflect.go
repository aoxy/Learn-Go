package main

import (
	"fmt"
	"reflect"
)

func main() {
	var fnum float64 = 3.14
	reflectFloat(fnum)
}

func reflectFloat(b interface{}) {
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue = %[1]v (%[1]T)\n", rValue)

	rType := rValue.Type()
	kind := rValue.Kind()
	fValue := rValue.Float()
	fmt.Printf("rType = %[1]v (%[1]T)\n", rType)
	fmt.Printf("kind = %[1]v (%[1]T)\n", kind)
	fmt.Printf("fValue = %[1]v (%[1]T)\n", fValue)

	iValue := rValue.Interface()
	fmt.Printf("iValue = %[1]v (%[1]T)\n", iValue)

	fnum2 := iValue.(float64)
	fmt.Printf("fnum2 = %[1]v (%[1]T)\n", fnum2)

}

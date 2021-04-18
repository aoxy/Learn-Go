// 这是不可运行的程序
package main

import (
	"fmt"
	"mypkgdemo\tempconv"
)

type celsiusFlag struct {Celsius}

func (f *celsiusFlag) Set(s string) error{
	var unit string
	var value float64
	fmt.Sscanf(s,"%f%f",&value,&unit)
	switch unit{
	case "C","°C":
		f.Celsius=Celsius(value)
		return nil
	case "F","°F":
		f.Celsius=FToC(Fahrenheut(value))
		return nil
	}
	return fmt.Errorf("非法温度 %q",s)
}

func CelsiusFlag(name string,value Celsius,usage string) *Celsius{
	f:=celsiusFlag{value}
	flag.CommondLine.Var(&f,name,usage)
	return &f.Celsius
}

var temp=tempconv.CelsiusFlag("temp")
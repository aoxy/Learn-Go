package main

import (
	"fmt"
	"sort"
)

type SortInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	names := StringSlice{"Jack", "Jerry", "Jhon", "Mary", "Danel", "Luoxiang", "Jotaro"}
	fmt.Println(names)
	sort.Sort(names)
	fmt.Println(names)
}

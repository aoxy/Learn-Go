package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Union(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var s, t IntSet
	for i := 0; i < 10; i++ {
		x := rand.Intn(13)
		s.Add(x)
		fmt.Print(x, " ")
	}
	fmt.Println()
	fmt.Println(s.String())
	for i := 0; i < 10; i++ {
		x := rand.Intn(13)
		t.Add(x)
		fmt.Print(x, " ")
	}
	fmt.Println()
	fmt.Println(t.String()) //打印方式1

	s.Union(&t)
	fmt.Println()
	fmt.Println(&s) //打印方式2

	for i := 0; i < 10; i++ {
		x := rand.Intn(13)
		fmt.Println(x, s.Has(x))
	}
}

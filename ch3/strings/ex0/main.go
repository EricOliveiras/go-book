package main

import (
	"fmt"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println("Goodbye" + s[5:])
	fmt.Println(rune('b'))
	fmt.Println(string(rune('\u4e16')))
	a := "abc"
	b := []byte(a)
	s2 := string(b)
	fmt.Println(s2)
}

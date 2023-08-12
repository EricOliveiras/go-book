package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// reverse(a[:])
	// fmt.Println(a)
	// s := []int{0, 1, 2, 3, 4, 5}
	// reverse(s[:2])
	// reverse(s[2:])
	// reverse(s)
	// fmt.Println(s)
	c := make([]string, 5)
	for i, v := range c {
		fmt.Println(i, v)
	}
}

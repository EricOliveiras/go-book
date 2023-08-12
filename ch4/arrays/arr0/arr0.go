package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)
	
	r := [...]int{99: -1}
	fmt.Printf("%T\n", r)
}

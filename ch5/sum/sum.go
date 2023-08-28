package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1))
	fmt.Println(sum(0))
}

func sum(values ...int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum
}

package main

import "fmt"

const bollingF = 212.0

func main() {
	var f = bollingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("bolling point = %g°F or %g°C\n", f, c)
}

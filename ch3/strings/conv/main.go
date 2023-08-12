package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)

	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)

	a, err := strconv.Atoi("123")
	b, err := strconv.ParseInt("123", 10, 64)

	if err != nil {
		fmt.Printf("error: %s", err)
	}
	
	fmt.Printf("%d\n", a)
	fmt.Printf("%d", b)
}

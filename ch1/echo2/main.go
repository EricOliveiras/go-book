package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args {
		s += sep + arg[1:]
		sep = " "
		fmt.Println(i)
	}
	fmt.Println(s)
}
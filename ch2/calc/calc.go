package main

import (
	"fmt"
	"gobook/ch2/mat"
	"os"
	"strconv"
)

func main() {
	x, err := strconv.ParseFloat(os.Args[2], 64)
	y, err := strconv.ParseFloat(os.Args[3], 64)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "soma":
		mat.Soma(x, y)
	case "sub":
		mat.Subtracao(x, y)
	case "div":
		mat.Divisao(x, y)
	case "mult":
		mat.Multiplicacao(x, y)
	}
}

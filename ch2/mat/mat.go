package mat

import "fmt"

func Soma(x, y float64) {
	soma := x + y
	fmt.Printf("%2.f + %2.f = %2.f\n", x, y, soma)
}

func Subtracao(x, y float64) {
	sub := x - y
	fmt.Printf("%2.f - %2.f = %2.f\n", x, y, sub)
}

func Divisao(x, y float64) {
	div := x / y
	fmt.Printf("%2.f / %2.f = %2.f\n", x, y, div)
}

func Multiplicacao(x, y float64) {
	mult := x * y
	fmt.Printf("%2.f x %2.f = %2.f\n", x, y, mult)
}

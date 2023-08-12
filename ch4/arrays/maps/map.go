package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Eric":   34,
		"Yuri":   27,
		"Hudson": 34,
	}

	delete(ages, "Eric")

	for name, age := range ages {
		fmt.Println(name, age)
	}

	fmt.Println("")

	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var pessoa = make(map[string]int)
	pessoa["eric"] = 0

	if age, ok := pessoa["eric"]; ok { fmt.Println(age) }

	x := equals(map[string]int{"A": 1}, map[string]int{"A": 2})
	fmt.Println(x)
}

func equals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
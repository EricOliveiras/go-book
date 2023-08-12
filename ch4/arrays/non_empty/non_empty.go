package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(s, 2))
}

func nonempty(strings []string) []string {
	out := strings[:0]
	fmt.Println(out)
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	fmt.Println(slice[i:], slice[i+1:], slice[:len(slice)-1])
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

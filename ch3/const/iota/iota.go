package main

import "fmt"

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

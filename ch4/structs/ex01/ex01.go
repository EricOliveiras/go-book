package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct{ X, Y int }

var eric Employee

func main() {
	p := Point{1, 2}
	q := Point{1, 1}

	fmt.Println(p.X == q.X && p.Y == q.Y)
}

func aumentarSalario(e *Employee) {
	e.Salary = e.Salary + 500
}

package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	s := square{side: 5}
	var sh shape = s
	fmt.Println("Square Area:", sh.area())
	fmt.Println("Square Perimeter:", sh.perimeter())

	c := circle{radius: 3}
	sh = c
	fmt.Println("Circle Area:", sh.area())
	fmt.Println("Circle Perimeter:", sh.perimeter())

}

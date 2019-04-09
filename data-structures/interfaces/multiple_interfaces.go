package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{5.0}

	var s Shape = c
	var o Object = c

	fmt.Println("Area of s of interface type Shape is: ", s.Area())
	fmt.Println("Volume of o of interface type Object is: ", o.Volume())
}

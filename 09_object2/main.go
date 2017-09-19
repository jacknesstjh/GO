package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	hight, width float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.hight * r.width
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func main() {
	r1 := Rectangle{12, 3}
	c1 := Circle{12}
	fmt.Println("The area of r1 is \n", r1.area())
	fmt.Println("The area of c1 is \n", c1.area())
}

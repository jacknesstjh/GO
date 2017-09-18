package main

import (
	"fmt"
)

type Rectangle struct {
	hight, width float64
}

func Area(r Rectangle) float64 {
	return r.hight * r.width
}

func main() {
	r1 := Rectangle{12, 4}
	r2 := Rectangle{3, 8}
	fmt.Println("The Area of r1 is", Area(r1))
	fmt.Println("The Area of r2 is", Area(r2))
}

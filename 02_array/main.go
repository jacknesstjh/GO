package main

import (
	"fmt"
)

func main() {
	x := [5]int{2, 3, 5, 6, 8}
	x[0] = 32
	x[3] = 11
	fmt.Println("Element is \n", x[0])
	fmt.Println("Element is \n", x[2])
}
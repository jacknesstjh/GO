package main

import (
	"fmt"
)

func main() {
	i := 10
	switch i {
	case 1:
		fmt.Printf("i equals 1")
	case 2, 3, 4:
		fmt.Printf("i equals 2,3,or 4")
	case 10:
		fmt.Printf("i equals 10")
	default:
		fmt.Printf("all i know is that i is a integer")
	}

}

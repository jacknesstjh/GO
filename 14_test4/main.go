package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
fmt.Printf("\n")
	for i := 0; i < 5; i++ {
		defer func() { fmt.Printf("%d ", i) }()
	}
}

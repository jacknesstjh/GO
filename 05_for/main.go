package main

import (
	"fmt"
)

func main() {
	sum := 0
	//for expression1;expression2;expression3
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to", sum)
}

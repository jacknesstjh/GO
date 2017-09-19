package main

import (
	"fmt"
)

var x map[string]string

func main() {
	x := make(map[string]string)
	x["hello"] = "world"
	x["ni"] = "hao"
	for k, v := range x {
		fmt.Println("x's key ", k)
		fmt.Println("x's value ", v)
	}
}

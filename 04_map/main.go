package main

import (
	"fmt"
)

var x map[string]int

func main() {
	x := make(map[string]int)
	x["Harry"] = 11
	x["Potter"] = 23
	//x=append(x["John"],100)
	//x = append(x,map[string]int{"John":100})
	//m = append(m, map[foo]bar{k: v}...)
	fmt.Println("They are %s\n", x)
}

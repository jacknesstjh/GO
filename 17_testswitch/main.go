package main

import (
	"fmt"
)

func main() {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		//fallthrough强制执行后面的case代码
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

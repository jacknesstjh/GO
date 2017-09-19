package main

import (
	"fmt"
)
var x[10] int
func main(){
	x[0]=22
	x[2]=33
	fmt.Printf("数组第一个值为 %d\n",x[0])
	fmt.Printf("数组第三个值为 %d\n",x[2])
}
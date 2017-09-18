package main

import (
	"fmt"
)

func main() {
	slice := []byte{'a', 'b', 'c', 'd'}
	//fmt.Printf("第二个元素是 %s\n", slice[1])
	fmt.Printf("第二个元素是 %s\n",string(slice[1]))
}

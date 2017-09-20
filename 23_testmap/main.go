package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"key1": "value1"}
	fmt.Println(m)
	m["key2"] = "value2"
	fmt.Println(m)
	p := m["key1"]
	fmt.Println(p)
	delete(m, "key1")
	fmt.Println(m)
}

package main

import (
	"fmt"
)
//即使用下面的方式声明了变量之后，还是要用make初始化
var x map[string]int

func main() {
    //这句初始化是必须要有的
	x := make(map[string]int)
	x["Harry"] = 11
	x["Potter"] = 23
	//x=append(x["John"],100)
	//x = append(x,map[string]int{"John":100})
	//m = append(m, map[foo]bar{k: v}...)
	fmt.Println("They are %s\n", x)
}

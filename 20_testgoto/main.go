package main

import ("fmt")

func main() {
    i := 0

fuck:
    fmt.Printf("i=%d\n", i)
    i++
    if (100 == i) {
        return
    }

    goto fuck
}
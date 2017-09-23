package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	count := make(chan int)
	//add a count of two,one for each gorountine.
	wg.Add(2)
	fmt.Println("Start Goroutines")
	go printCounts("A", count)
	go printCounts("B", count)
	fmt.Println("Channel begin")
	count <- 1
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
func printCounts(lable string, count chan int) {
	defer wg.Done()
	for {
		val, ok := <-count
		if !ok {
			fmt.Println("channel was cloesd")
			return
		}
		fmt.Println("Count: %d received from %s \n", val, lable)
		if val == 10 {
			fmt.Printf("Channel Closed from %s \n", lable)
			close(count)
			return
		}
		val++
		count <- val
	}

}

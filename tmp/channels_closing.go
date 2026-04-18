//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	transferPoint := make(chan int)
	// miner
	go func() {
		// 3 4 5 6 7 8 9 10 11 12
		iterations := 3 + rand.Intn(10)
		fmt.Println("iterations:", iterations)
		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(100 * time.Millisecond)
		}
		close(transferPoint)
	}()

	coal := 0

	/*for {
		v, ok := <-transferPoint
		if !ok {
			fmt.Println("All iterations are finished")
			break
		}
		coal += v
		fmt.Println("coal:", coal)
	}*/

	for v := range transferPoint { // in case close(transferPoint) cycle is stopped
		coal += v
		fmt.Println("coal:", coal)
	}
	fmt.Println("Rotal coal=", coal)

	//var ch chan string //nil channel
	var ch chan string = make(chan string)

	go func() {
		ch <- "Hello" //block for write to nil chan
	}()

	s := <-ch //block for read from nil chan
	fmt.Println("s=", s)
}

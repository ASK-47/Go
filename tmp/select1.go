//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intChanale := make(chan int)
	strChanale := make(chan string)

	go func() {
		i := 0
		for {
			intChanale <- i
			i++
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		i := 0
		for {
			strChanale <- "hi" + strconv.Itoa(i)
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(50 * time.Millisecond)

	for {
		select {
		case number := <-intChanale:
			fmt.Println("intChanale=", number)
		case str := <-strChanale:
			fmt.Println("strChanale=", str)
		default:
			fmt.Println("NO Waiting Chanals")
		}
	}
}

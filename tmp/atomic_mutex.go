//go:build ignore

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//var n int = 0

var n atomic.Int64
var slice []int
var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done() //exit from wait group (instead deadlock)
	for i := 1; i <= 1000; i++ {
		//n++
		n.Add(1) //atomic incrirement - only for ONE goroutine
		mtx.Lock()
		slice = append(slice, 1) //locked safe append only for ONE goroutine
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)

	wg.Wait() //wait for finishing all goroutines in wg (unless wg counter becomes 0)
	fmt.Println("main завершился!", "n=", n.Load(), "len(slice)=", len(slice))

}

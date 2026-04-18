package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// fan‑in pattern
func main() {
	//var coal int //race for coal <-main, +goroutine {coal += v}
	var coal atomic.Int64 //to prevent the race

	var mails []string //race for coal <-main, +goroutine {mails = append(mails, v)}
	var mtx sync.Mutex //to prevent the race

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second) //working time limit for miner
		fmt.Println("MINER'S WORKING DAY IS OVER")

		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second) //working time limit for postman
		fmt.Println("POSTMAN'S WORKING DAY IS OVER")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 10)
	mailTransferPoint := postman.PostmanPool(postmanContext, 10)

	before := time.Now()

	/*isColalClosed := false //chan is open
	isMailClosed := false  //chan is open

	for !isColalClosed || !isMailClosed { //while one of chan is open
		select {
		case c, ok := <-coalTransferPoint:
			if !ok { //if c is default value
				isColalClosed = true
				continue
			}
			coal += c
		case m, ok := <-mailTransferPoint:
			if !ok {
				isMailClosed = true
				continue
			}
			mails = append(mails, m)
		}
	}*/

	//alt via goroutines + wg

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock() //to prevent the race
			mails = append(mails, v)
			mtx.Unlock() //to prevent the race
		}
	}(wg)

	wg.Wait()

	fmt.Println("Total Coal=", coal.Load())

	mtx.Lock() //to prevent the race
	fmt.Println("Total Mails", len(mails))
	mtx.Unlock()

	fmt.Println("Total time=", time.Since(before))
}

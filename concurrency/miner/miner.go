package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) { //sendOnly channel chan<-  int, forbid to get from channel
	defer wg.Done()
	for { //stop immediatly if ctx.Done during 1 sec delay  or during  transferPoint <- power
		fmt.Println("Miner #", n, "HAS STARTED to get the coal")
		select {
		case <-ctx.Done():
			fmt.Println("Miner #", n, "HAS FINIHED to get the coal")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Miner #", n, "HAS RECIEVED the the coal, P=", power)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Miner #", n, "HAS FINIHED to get the coal")
			return
		case transferPoint <- power:
			fmt.Println("Miner #", n, "HAS RECIEVED the the coal, P=", power)
		}
	}
	/*for {
		select {
		case <-ctx.Done():
			fmt.Println("Miner #", n, "HAS FINIHED to get the coal")
			return
		default:
			fmt.Println("Miner #", n, "HAS STARTED to get the coal")
			time.Sleep(1 * time.Second)
			fmt.Println("Miner #", n, "HAS RECIEVED the the coal, P=", power)
			transferPoint <- power
			fmt.Println("Miner #", n, "HAS LOADED the the coal, P=", power)
		}
	}*/

}

func MinerPool(ctx context.Context, minersNumber int) <-chan int { // <- read only chan
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 0; i < minersNumber; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() { //goroutine for chan closing - work independ form MinerPool
		wg.Wait()
		close(coalTransferPoint) //stop to read in rannge for in main
	}()

	return coalTransferPoint
}

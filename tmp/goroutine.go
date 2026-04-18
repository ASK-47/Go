//go:build ignore

package main

import (
	"fmt"
	"time"
)

func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался…")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")
	transferPoint <- 10
	fmt.Println("Поход номер", n, "Уголь в буфере")
}

func main() {
	coal := 0
	transferPoint := make(chan int, 6)
	initTime := time.Now()
	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)
	go mine(transferPoint, 4)
	go mine(transferPoint, 5)
	go mine(transferPoint, 6)

	coal += <-transferPoint
	time.Sleep(3 * time.Second)
	coal += <-transferPoint
	time.Sleep(3 * time.Second)
	coal += <-transferPoint
	time.Sleep(3 * time.Second)
	coal += <-transferPoint
	time.Sleep(3 * time.Second)
	coal += <-transferPoint
	time.Sleep(3 * time.Second)
	coal += <-transferPoint
	time.Sleep(3 * time.Second)

	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))
}

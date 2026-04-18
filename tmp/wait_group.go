//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done() //exit from wait group (instead deadlock)

	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон, я отнёс газету", text, "в", i, "раз")
		time.Sleep(250 * time.Millisecond)
	}
	//wg.Done() //exit from wait group (instead deadlock)
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1) //add gouroutine into waiting group
	go postman(wg, "Новости")

	wg.Add(1)
	go postman(wg, "Игровой журнал")

	wg.Add(1)
	go postman(wg, "Автомобильная хроника")

	wg.Wait() //wait for finishing all goroutines in wg (unless wg counter becomes 0)

	fmt.Println("main завершился!")
}

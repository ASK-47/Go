package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("PID:", os.Getpid())
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGTERM)
	go func() {
		for {
			s := <-sigchan
			fmt.Println("got system signal:", s)
		}
	}()
	time.Sleep(1 * time.Hour)
}

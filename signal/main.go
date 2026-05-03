package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func job(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("JOB is done")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("JOB is progress i=", i)
			i++
		}
	}
}

func main() {
	fmt.Println("PID:", os.Getpid())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	job(ctx)

	fmt.Println("Application is stopped by SIGTERM")

}

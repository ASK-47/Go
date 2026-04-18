//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done(): //if context is canceled
			fmt.Println("foo is finished", n)
			return
		default:
			fmt.Println("foo")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("boo is finished", n)
			return
		default:
			fmt.Println("boo")
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background()) //parentContext with Cancel from context.Background()
	childContext, childCancel := context.WithCancel(parentContext)          //parentContext with Cancel from context.Background()

	go foo(parentContext, 1)
	go boo(childContext, 1)
	go foo(parentContext, 2)
	go boo(childContext, 2)
	go foo(parentContext, 3)
	go boo(childContext, 3)
	go foo(parentContext, 4)
	go boo(childContext, 4)
	go foo(parentContext, 5)
	go boo(childContext, 5)

	time.Sleep(1000 * time.Millisecond)
	childCancel()

	time.Sleep(1000 * time.Millisecond)
	parentCancel()

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("main is finished")
}

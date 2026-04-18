//go:build ignore

package main

import (
	"time"

	"github.com/k0kubun/pp"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Друг 1",
				Text:   "Привет",
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			messageChan2 <- Message{
				Author: "Друг 2",
				Text:   "Пока",
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case message := <-messageChan1:
			pp.Println("Message has bee recieved from", message.Author, message.Text)
		case message := <-messageChan2:
			pp.Println("Message has bee recieved from", message.Author, message.Text)
			//default:
			//	pp.Println("NO MASSAGE")
		}
	}
}

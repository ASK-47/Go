package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, mail string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Postman #", n, "HAS FINISHED")
			return
		default:
			fmt.Println("Mail=", mail, "HAS BEEN RECIEVED by Postman #", n)
			time.Sleep(1 * time.Second)
			fmt.Println("Mail=", mail, "HAS BEEN DELIVERED to Post Office by Postman #", n)
			transferPoint <- mail
			fmt.Println("Mail=", mail, "HAS BEEN RECIEVED by Post Office")
		}
	}
}

func PostmanPool(ctx context.Context, posmanNumber int) <-chan string { // <- read only chan
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 0; i < posmanNumber; i++ {
		wg.Add(1)
		go postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() { //goroutine for chan closing - work independ form MinerPool
		wg.Wait()
		close(mailTransferPoint) //stop to read in rannge for in main
	}()

	return mailTransferPoint
}

func postmanToMail(n int) string {
	ptm := map[int]string{
		1: "Griteengs",
		2: "Invitation",
		3: "Info",
		5: "Account",
		6: "Alert",
		7: "News",
	}
	mail, ok := ptm[n]
	if !ok {
		return "Spam"
	}
	return mail
}

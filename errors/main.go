package main

import (
	"err/common"
	"err/transport"
	"errors"
	"fmt"
)

func main() {
	user, err := transport.GetUserTransport()
	if err != nil {
		if errors.Is(err, common.ErNotFound) {
			fmt.Println("Satus code 404: ", err)
		} else if errors.Is(err, common.ErBrokenConnection) {
			fmt.Println("Satus code 503: ", err)
		} else {
			fmt.Println("Satus code 500: ", err)
		}
	} else {
		fmt.Println("user", user, err)
	}
}

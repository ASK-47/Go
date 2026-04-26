package main

import (
	"docker/http_server"
	"fmt"
)

func main() {
	fmt.Println("HTTP server is starting")
	err := http_server.SratrHTTPserver()
	if err != nil {
		fmt.Println("Error in server operation:", err)
	} else {
		fmt.Println("Server operation is successful")
	}

	/*_, err := os.Create("out/newfile.txt")
	if err != nil {
		panic(err)
	}*/
}

//go:build ignore

package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("Handler отработал успешно!")
}

func handlerSleep(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	_, err := w.Write([]byte("HTTP-response"))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("Handler отработал успешно!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sleep", handlerSleep)

	fmt.Println("Запуск HTTP-cервера")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("Стоп HTTP-cервера")
}

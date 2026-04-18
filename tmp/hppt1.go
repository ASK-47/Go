//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Новый платёж обработан!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Bо время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно совершил оплату!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата отменена!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно отменил оплату!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, world!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Bо время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно обработал HTTP запрос!")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Запуск HTTP-cервера")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	}
	fmt.Println("Стоп HTTP-cервера")
}

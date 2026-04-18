//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

// var money int
var money atomic.Int64

// var bank int = 0
var bank atomic.Int64

var mtx sync.Mutex

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err) // no need to use manual err.Error()
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to int:", err) // no need to use manual err.Error()
		return
	}
	mtx.Lock()
	if money.Load()-int64(paymentAmount) >= 0 {
		//money -= paymentAmount //race for mnoney, since payHandler is gourutine
		money.Add(int64(-paymentAmount)) //atomic
		fmt.Println("Payment SUCCESSFUL")
		fmt.Println("money=", money.Load())
		fmt.Println("bank=", bank.Load())
	} else {
		fmt.Println("Not enough money")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err) // no need to use manual err.Error()
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to int:", err) // no need to use manual err.Error()
		return
	}
	mtx.Lock()
	if money.Load()-int64(saveAmount) >= 0 {
		money.Add(int64(-saveAmount)) //atomic
		bank.Add(int64(saveAmount))   //atomic
		fmt.Println("Save SUCCESSFUL")
		fmt.Println("money=", money.Load())
		fmt.Println("bank=", bank.Load())
	} else {
		fmt.Println("Not enough money")
	}
	mtx.Unlock()
}

func salaryHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err) // no need to use manual err.Error()
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	salaryAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to int:", err) // no need to use manual err.Error()
		return
	}
	mtx.Lock()
	money.Add(int64(salaryAmount)) //atomic
	fmt.Println("Salary SUCCESSFUL")
	fmt.Println("money=", money.Load())
	fmt.Println("bank=", bank.Load())
	mtx.Unlock()
}

func main() {

	money.Add(1000) //atomic init
	bank.Add(0)     //atomic init

	http.HandleFunc("/pay", payHandler)       //handler registration
	http.HandleFunc("/save", saveHandler)     //handler registration
	http.HandleFunc("/salary", salaryHandler) //handler registration

	fmt.Println("Запуск HTTP-cервера")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("Стоп HTTP-cервера")
}

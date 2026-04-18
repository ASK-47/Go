//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money int = 1000
var bank int = 0
var mtx sync.Mutex

func payHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTTP-method=", r.Method)
	for k, v := range r.Header {
		fmt.Println("key=", k, "\t", "value=", v)
	}
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wright HTTP respose:", err) // no need to use manual err.Error()
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "fail to convert HTTP body to int:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wright HTTP respose:", err) // no need to use manual err.Error()
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money -= paymentAmount
		msg := "Payment SUCCESSFUL, money=:" + strconv.Itoa(money) + " bank=" + strconv.Itoa(bank)
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wright HTTP respose:", err) // no need to use manual err.Error()
		}
	} else {
		fmt.Println("Not enough money")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wright HTTP respose:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "fail to convert HTTP body to int:" + err.Error()
		fmt.Println(msg)
		return
	}
	mtx.Lock()
	if money-saveAmount >= 0 {
		money -= saveAmount
		bank += saveAmount
		msg := "Save SUCCESSFUL, money=:" + strconv.Itoa(money) + " bank=" + strconv.Itoa(bank)
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wright HTTP respose:", err) // no need to use manual err.Error()
		}
	} else {
		fmt.Println("Not enough money")
	}
	mtx.Unlock()
}

func salaryHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wright HTTP respose:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	salaryAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "fail to convert HTTP body to int:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wright HTTP respose:", err)
		}
		return
	}
	mtx.Lock()
	money += salaryAmount
	msg := "Salary SUCCESSFUL, money=:" + strconv.Itoa(money) + " bank=" + strconv.Itoa(bank)
	fmt.Println(msg)
	_, err1 := w.Write([]byte(msg))
	if err1 != nil {
		fmt.Println("fail to wright HTTP respose:", err) // no need to use manual err.Error()
	}
	mtx.Unlock()
}

func main() {

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

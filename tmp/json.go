//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"desctiption"` //json teg
	USD         int    `json:"usd"`         //json teg
	FullName    string `json:"fullName"`    //json teg
	Adress      string `json:"adress"`      //json teg
	Time        time.Time
}

type HttpResponse struct {
	Money          int       `json:"My_Money"`   //json teg => My_Money see in response
	PaymentHistory []Payment `json:"My_Payment"` //json teg => My_Payment see in response
}

var paymentHistory = make([]Payment, 0)

var money int = 1000
var mtx sync.Mutex

func payHandler(w http.ResponseWriter, r *http.Request) {
	/*fmt.Println("HTTP-method=", r.Method)
	for k, v := range r.Header {
		fmt.Println("key=", k, "\t", "value=", v)
	}*/

	var payment Payment

	//json parsing1
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	//json parsing2
	//httpRequestBody, err := io.ReadAll(r.Body)
	/*if err := json.Unmarshal(httpRequestBody, &payment); err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	/*httpRequestBodyString := string(httpRequestBody)
	parts := strings.SplitN(httpRequestBodyString, ",", 2) //string slice
	if len(parts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	paymentAmount, err := strconv.Atoi(parts[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payment := Payment{
		Description: parts[1],
		USD:         paymentAmount,
	}*/

	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD

	} else {
		fmt.Println("Not enough money")
	}
	paymentHistory = append(paymentHistory, payment)

	//for response
	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	//response 1 over MarshalIndent
	/*b, err := json.MarshalIndent(httpResponse, "", "	")
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	//response 2 over Marshal
	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//for response
	if _, err := w.Write(b); err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("Money=", money)
	payment.Println()
	mtx.Unlock()
}

func (p *Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Adress:", p.Adress)
}

func main() {

	http.HandleFunc("/pay", payHandler) //handler registration

	fmt.Println("Start HTTP-server")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("HTTP-server error:", err.Error())
	}
	fmt.Println("Stop HTTP-cервера")
}

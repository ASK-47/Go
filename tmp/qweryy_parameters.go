//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

//localhost:9091/default?foo=x&boo=y

func Handler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")
	fmt.Println("fooParam=", fooParam)
	fmt.Println("booParam=", booParam)
}

func main() {

	http.HandleFunc("/default", Handler) //handler registration

	fmt.Println("Start HTTP-server")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("HTTP-server error:", err.Error())
	}
	fmt.Println("Stop HTTP-cервера")
}

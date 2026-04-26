package http_server

import (
	"errors"
	"fmt"
	"net/http"
)

func SratrHTTPserver() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Treating request in pattern /ping")
		w.Write([]byte("Hello from Docker\n"))
	})

	err := http.ListenAndServe(":5050", nil)

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}

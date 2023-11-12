package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

const (
	apiKey    = "ed070817-20b0-490e-a11a-aaa186a0df49"
	apiSecret = "0c4ead76-4ab6-47f8-921c-de00b2bbfbc5"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, your key is: %s\n", apiKey)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user := r.Form.Get("user")
		pw := r.Form.Get("password")

		slog.Info("Registering new user %s with password %s.\n", user, pw)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

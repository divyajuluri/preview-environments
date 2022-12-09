package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandler() http.HandlerFunc {
	now := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World 4! ( build date: %s)", now)
	})
}

func main() {
	http.HandleFunc("/", mainHandler())
	http.ListenAndServe(":8080", nil)
}

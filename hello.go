package main

import (
	"fmt"
	"log"
	"net/http"
)

func say(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "hello..."+r.RemoteAddr+"到访")
	if err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/say", say)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

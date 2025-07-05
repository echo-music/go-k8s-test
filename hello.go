package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func say(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	_, err := fmt.Fprintln(w, hostname+" 热烈欢迎 "+r.RemoteAddr+" 到访")
	if err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/say", say)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

package main

import (
	"fmt"
	"github.com/echo-music/go-k8s-test/order/client"

	//"github.com/echo-music/go-k8s-test/order/client"

	"log"
	"net/http"
	"os"
)

func say(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	res := client.CreateOrder()
	_, err := fmt.Fprintln(w, hostname+" 热烈欢迎 "+r.RemoteAddr+" 到访. grpc result"+res)
	if err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/say", say)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

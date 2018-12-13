package main

import (
	"fmt"
	"net/http"
)

const (
	listenPort = "8080"
)

func getMessage(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	fmt.Println("Got message from", r.RemoteAddr)
	fmt.Println("->", message)
	fmt.Fprintln(w, "Got message -> "+message)
}

func main() {
	http.HandleFunc("/", getMessage)
	fmt.Printf("Waiting for messages on port %v...\n", listenPort)
	err := http.ListenAndServe(":"+listenPort, nil)
	fmt.Println("ERROR: ", err)
}

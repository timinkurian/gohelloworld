package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", SayHello)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

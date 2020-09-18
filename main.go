package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/hello", SayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

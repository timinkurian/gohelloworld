package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/hello", SayHello)
	http.ListenAndServe("0.0.0.0:80", handler)
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello world`)
}

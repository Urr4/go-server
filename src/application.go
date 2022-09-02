package main

import (
	"go-is-cool/src/endpoints"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", endpoints.Greet)
	http.ListenAndServe(":8090", nil)
}
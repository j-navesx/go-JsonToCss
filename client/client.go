package main

import (
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(resp)
	fmt.Println(err)
}
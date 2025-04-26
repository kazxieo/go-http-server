package main

import (
	"net/http"
)

func main() {
	panic(http.ListenAndServe(":8080", nil))
}

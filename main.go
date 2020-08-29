package main

import (
	"./internal"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("A server running at 8080.")
	internal.DefineURIHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running API server...")

	router := router.CreateRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}

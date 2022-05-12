package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	router := router.CreateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}

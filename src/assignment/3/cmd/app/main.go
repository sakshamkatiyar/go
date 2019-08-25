package main

import (
	"log"
	"net/http"
	"assignment/3/pkg/router"
)

func main () {
	router.Routers()
	log.Println("Listening...")
	http.ListenAndServe(":8080", router.Mux)
}
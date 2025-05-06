package main

import (
	"log"
	"net/http"
)

func main() {
	//andler := ()
	http.HandleFunc("/health-check", healthCheckHandler)
	log.Println("Server started runnon on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

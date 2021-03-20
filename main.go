package main

import (
	"log"
	"net/http"

	h "hello/handlers"
	m "hello/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(m.AddContentType)
	router.HandleFunc("/", h.RequestHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GetHello struct {
	Message string `json:"message"`
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(&GetHello{
		Message: "Hello World",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func addContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(addContentType)
	router.HandleFunc("/", requestHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}

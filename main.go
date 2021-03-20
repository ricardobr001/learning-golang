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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", requestHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}

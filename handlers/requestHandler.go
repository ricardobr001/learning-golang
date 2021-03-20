package handlers

import (
	"encoding/json"
	"net/http"
)

type GetHello struct {
	Message string `json:"message"`
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(&GetHello{
		Message: "Hello World",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

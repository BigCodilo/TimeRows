package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//r.HandleFunc("/average", handlres.AverageHandlerGet).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}

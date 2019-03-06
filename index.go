package main

import (
	"net/http"

	"github.com/BigCodilo/TimeRows/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/average", handlers.AverageHandlerGet).Methods("GET")
	r.HandleFunc("/average", handlers.AverageHandlerPost).Methods("POST")

	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}

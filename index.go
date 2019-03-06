package main

import (
	"net/http"

	"github.com/BigCodilo/TimeRows/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./static/")))

	r.HandleFunc("/average", handlers.AverageHandlerGet).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}

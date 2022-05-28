package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var citys []City

func main() {
	r := mux.NewRouter()

	citys = loadCitys()
	_ = citys

	r.HandleFunc("/citys", getCitys).Methods("GET")
	r.HandleFunc("/citys/{name}", getCity).Methods("GET")
	r.HandleFunc("/citys", createCity).Methods("POST")
	r.HandleFunc("/citys/{name}", updateCity).Methods("PUT")
	r.HandleFunc("/citys/{name}", deleteCity).Methods("DELETE")

	fmt.Println("Api Started...")

	log.Fatal(http.ListenAndServe(":8000", r))
}

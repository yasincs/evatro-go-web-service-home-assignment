package main

//import packages
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var citys []City

func main() {

	// Init router
	r := mux.NewRouter()

	//Adding cities read from json to database
	citys = loadCitys()
	_ = citys

	// Route handles & endpoints
	r.HandleFunc("/citys", getCitys).Methods("GET")
	r.HandleFunc("/citys/{name}", getCity).Methods("GET")
	r.HandleFunc("/citys", createCity).Methods("POST")
	r.HandleFunc("/citys/{name}", updateCity).Methods("PUT")
	r.HandleFunc("/citys/{name}", deleteCity).Methods("DELETE")

	// Start server
	fmt.Println("Api Started...")
	log.Fatal(http.ListenAndServe(":8000", r))

}

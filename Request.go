package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getCitys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(citys)
}
func getCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range citys {
		if item.Name == params["name"] {
			var response Response
			response.Status = "success"
			response.Response = item
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	var status Status
	status.Status = "failure"
	json.NewEncoder(w).Encode(status)
}
func createCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var city City
	_ = json.NewDecoder(r.Body).Decode(&city)
	for _, item := range citys {
		if item.Name == city.Name {
			var status Status
			status.Status = item.Name + " city is exist"
			json.NewEncoder(w).Encode(status)
			return
		}
	}
	//city.Name = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	citys = append(citys, city)
	json.NewEncoder(w).Encode(city)
}

func updateCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range citys {
		if item.Name == params["name"] {
			citys = append(citys[:index], citys[index+1:]...)
			var city City
			_ = json.NewDecoder(r.Body).Decode(&city)
			city.Name = params["name"]
			citys = append(citys, city)
			var response Response
			response.Status = "success"
			response.Response = city
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	var status Status
	status.Status = params["name"] + " city is not exist"
	json.NewEncoder(w).Encode(status)
}

// Delete book
func deleteCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range citys {
		if item.Name == params["name"] {
			citys = append(citys[:index], citys[index+1:]...)
			json.NewEncoder(w).Encode(citys)
			return
		}
	}
	var status Status
	status.Status = params["name"] + " city is not exist"
	json.NewEncoder(w).Encode(status)
}

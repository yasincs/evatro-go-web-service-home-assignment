package main

//import packages
import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Get all cities
Takes no parameters
*/
func getCitys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(citys)
}

/*
Get single city
Takes city name as parameter
Return city
If city is not exist return failure
*/
func getCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through citys and find one with the name from the params
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

/*
Add new city
Takes city struct as parametre
Return Response struct {success and city}
If city is exist return failure
*/
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
	citys = append(citys, city)
	json.NewEncoder(w).Encode(city)
}

/*
Update city
Takes city name and city struct as parametre
Return Response struct {success and city}
If city is not exist return failure
*/
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

/*
Delete city
Return cities
If city is not exist return failure
*/
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

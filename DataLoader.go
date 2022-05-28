package main

//import packages
import (
	"encoding/json"
	"io/ioutil"
)

/*
Read city.json
Convert city model
Return cities array
*/
func loadCitys() []City {
	bytes, _ := ioutil.ReadFile("city.json")
	var citys []City
	json.Unmarshal(bytes, &citys)
	return citys
}

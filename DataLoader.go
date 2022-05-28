package main

import (
	"encoding/json"
	"io/ioutil"
)

func loadCitys() []City {
	bytes, _ := ioutil.ReadFile("city.json")
	var citys []City
	json.Unmarshal(bytes, &citys)
	return citys
}

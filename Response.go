package main

//Response struct
type Response struct {
	Status   string `json:"status"`
	Response City   `json:"city"`
}

package main

type Response struct {
	Status   string `json:"status"`
	Response City   `json:"city"`
}

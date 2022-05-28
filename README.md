# evatro-go-web-service-home-assignment
a web service that returns the current temperature of a given city name 

## About Implementation Details
- Programming Language : Go
- **Postman** collection file is provided to test study.

## Setup & Run
- Download and install [Go](https://go.dev/dl/) 
- Install Go Extensions
- Clone repository to your computer.
- Open the terminal and go to the repository folder location.
  
## Json file format

```sh
# Json format example
# [
#	{"Name": "Istanbul", "Temperature": "8", "Unit": "8"},
#	{"Name": "Prag", "Temperature": "10", "Unit": "8"},
#	{"Name": "Los Angeles", "Temperature": "9", "Unit": "8"},
#	{"Name": "Kiev", "Temperature": "-5", "Unit": "8"}
# ]
```

## Build Project

```sh
go build
```
## Run API

```sh
./EvatroGo
```


## Endpoints

### Get All Cities

```sh
http://localhost:8000/citys
```
### Get Single City 

```sh
http://localhost:8000/citys/Istanbul
```
### Create New City

```sh
http://localhost:8000/citys

# Request sample
# {
#    "name": "kuala lumpur",
#    "temperature": "12",
#    "unit": "celcius"
# }
```

### Update City
```sh
http://localhost:8000/citys/Istanbul

# Request sample
# {
#   "name": "Istanbul",
#   "temperature": "8",
#   "unit": "celcius"
# }
```

### Delete City

```sh
http://localhost:8000/citys/Istanbul
```

#### Testing


[Click here](https://github.com/yasincs/evatro-go-web-service-home-assignment/blob/main/Evatro-Go.postman_collection.json) to download Postman Collection.

You can open the collection in [Postman Application](https://www.postman.com/downloads/) and test with different parameters.

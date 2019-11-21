package main

import (
	route "eajardini/gin/gocrud/route"
)

func main() {
	//curl -v -X GET http://localhost:8080/

	//curl --header "Content-Type: application/json" --request POST --data '{"username":"xyz","password":"xyz"}' http://localhost:8080/cliente

	route.IniciaServidor()

}

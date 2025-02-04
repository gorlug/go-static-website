package main

import (
	"go-static-website/routes"
	"log"
	"net/http"
)

func main() {
	api := routes.CreateApi()

	log.Fatal(http.ListenAndServe(":8090", api))
}

package main

import (
	"flag"
	"fmt"
	"go-lana/pkg/application"
	"go-lana/pkg/handlers"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {

	app := application.New()
	port := flag.String("port", defaultPort, "Set port for run application")
	flag.Parse()

	portNumber := fmt.Sprintf(":%s", *port)

	app.Handle("/products/([A-Z]+)$", handlers.ProductHandler, "GET")
	app.Handle("/products/$", handlers.ProductsHandler, "GET")
	app.Handle("/cart/$", handlers.CartHandler, "POST")
	app.Handle("/get-cart/$", handlers.GetCartHandler, "GET")
	app.Handle("/add-to-cart/$", handlers.AddItemCartHandler, "PUT")
	app.Handle("/delete-cart/$", handlers.DeleteCartHandler, "DELETE")

	fmt.Println("Running Application on port" + portNumber)
	err := http.ListenAndServe(portNumber, app)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

package main

import (
	"fmt"
	"go-lana/pkg/application"
	"go-lana/pkg/handlers"
	"log"
	"net/http"
)

const port = "8081"

func main() {

	app := application.New()
	portNumber := fmt.Sprintf(":%s", port)

	app.Handle("/products/([A-Z]+)$", handlers.ProductHandler, "GET")
	app.Handle("/products/$", handlers.ProductsHandler, "GET")
	app.Handle("/cart/$", handlers.CartHandler, "POST")
	app.Handle("/add-to-cart/$", handlers.AddItemCartHandler, "PUT")
	app.Handle("/delete-cart/$", handlers.DeleteCartHandler, "DELETE")

	fmt.Println("Running Application on port:" + port)
	err := http.ListenAndServe(portNumber, app)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

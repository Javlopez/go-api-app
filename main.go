package main

import (
	"fmt"
	"go-lana/pkg/app"
	"go-lana/pkg/handlers"
	"go-lana/pkg/infrastructure"
	"log"
	"net/http"
)

const port = "8081"

func main() {

	c := infrastructure.Container{}
	lanaApp := &app.ApplicationContext{Version: app.Version, Container: c}

	http.Handle("/products", app.ApplicationHandler{lanaApp, handlers.ProductsHandler})

	portNumber := fmt.Sprintf(":%s", port)
	log.Printf("Server running on port %s......", port)
	http.ListenAndServe(portNumber, nil)

}

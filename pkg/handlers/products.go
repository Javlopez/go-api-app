package handlers

import (
	"fmt"
	"go-lana/pkg/app"
	"net/http"
)

//ProductsHandler method
func ProductsHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {

	resp := struct {
		Active string
	}{Active: "Running"}

	result, _ := a.Container.ProductService().GetAll()

	fmt.Printf("Loading data %+v", result)

	return 200, resp
}

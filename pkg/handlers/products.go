package handlers

import (
	"go-lana/pkg/app"
	"net/http"
)

//ProductsHandler method
func ProductsHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	resp := struct {
		Active string
	}{Active: "Running"}
	return 200, resp
}

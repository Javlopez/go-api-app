package handlers

import (
	app "go-lana/pkg/application"
	"net/http"
)

//ProductsHandler method
func ProductsHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	result, _ := a.Container.ProductService().GetAll()
	return 200, result
}

func ProductHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {

	code := a.Params[0]
	result, err := a.Container.ProductService().GetProductByCode(code)

	if err != nil {
		return http.StatusNotFound, struct{ Message string }{Message: err.Error()}
	}
	return http.StatusOK, result
}

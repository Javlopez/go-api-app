package handlers

import (
	app "go-lana/pkg/application"
	jsonResponse "go-lana/pkg/response/json"
	"net/http"
)

//ProductsHandler method returns a catalog of the products
func ProductsHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) *jsonResponse.Response {
	result, _ := a.Container.ProductService().GetAll()
	return jsonResponse.NewSuccessResponse(http.StatusOK, result)
}

//ProductHandler returns a product by using a identifier (productCode)
func ProductHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) *jsonResponse.Response {

	code := a.Params[0]
	result, err := a.Container.ProductService().GetProductByCode(code)

	if err != nil {
		return jsonResponse.NewErrorResponse(http.StatusNotFound, err.Error())
	}
	return jsonResponse.NewSuccessResponse(http.StatusOK, result)
}

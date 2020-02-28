package handlers

import (
	"encoding/json"
	"errors"
	app "go-lana/pkg/application"
	jsonResponse "go-lana/pkg/response/json"
	"io/ioutil"
	"log"
	"net/http"
)

//CartHandler method
func CartHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) *jsonResponse.Response {

	cart := a.Container.CartService().CreateCart()

	return jsonResponse.NewSuccessResponse(http.StatusCreated, cart)
}

func AddItemCartHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) *jsonResponse.Response {

	var cartReader struct {
		Cart  string
		Items []string
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf(err.Error())
		return jsonResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(reqBody, &cartReader)
	if err != nil {
		log.Printf(err.Error())
		return jsonResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
	}

	cart, err := a.Container.CartService().AddItem(cartReader.Cart, cartReader.Items)

	if err != nil {
		log.Printf(err.Error())
		return jsonResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return jsonResponse.NewSuccessResponse(http.StatusOK, cart)
}

//DeleteCartHandler method
func DeleteCartHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) *jsonResponse.Response {

	var cartReader struct {
		Cart string
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf(err.Error())
		return jsonResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(reqBody, &cartReader)
	if err != nil {
		log.Printf(err.Error())
		return jsonResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
	}

	err = a.Container.CartService().DeleteCart(cartReader.Cart)

	if err != nil {
		log.Printf(err.Error())
		errorDeleted := errors.New("The cart cannot be deleted, maybe the cart no longer exists")
		return jsonResponse.NewErrorResponse(http.StatusInternalServerError, errorDeleted.Error())
	}

	deleteOk := struct {
		Deleted string
	}{
		Deleted: "ok",
	}
	return jsonResponse.NewSuccessResponse(http.StatusOK, deleteOk)
}

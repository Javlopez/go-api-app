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
	a.Lock()
	cart := a.Container.CartService().CreateCart()
	a.Unlock()
	return jsonResponse.NewSuccessResponse(http.StatusCreated, cart)
}

func GetCartHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) *jsonResponse.Response {
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

	a.RLock()
	cart, err := a.Container.CartService().GetCart(cartReader.Cart)
	a.RUnlock()
	if err != nil {
		log.Printf(err.Error())
		erroGetCart := errors.New("The cart does not exists")

		return jsonResponse.NewErrorResponse(http.StatusInternalServerError, erroGetCart.Error())
	}

	return jsonResponse.NewSuccessResponse(http.StatusOK, cart)
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

	a.Lock()
	cart, err := a.Container.CartService().AddItem(cartReader.Cart, cartReader.Items)
	a.Unlock()

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

	a.Lock()
	err = a.Container.CartService().DeleteCart(cartReader.Cart)
	a.Unlock()

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

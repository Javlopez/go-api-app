package app

import (
	"encoding/json"
	"go-lana/pkg/infrastructure"
	"net/http"
)

const (
	//Version 1.0 beta
	Version = "1.0"
)

//ApplicationContext struct
type ApplicationContext struct {
	Version   string
	Container infrastructure.Container
}

//ApplicationHandler struct
type ApplicationHandler struct {
	*ApplicationContext
	H func(*ApplicationContext, http.ResponseWriter, *http.Request) (int, interface{})
}

func (ah ApplicationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode int
		data       interface{}
	)

	statusCode, data = ah.H(ah.ApplicationContext, w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
	return
}

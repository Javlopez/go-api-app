package application

import (
	"encoding/json"
	"fmt"
	"go-lana/pkg/infrastructure"
	"net/http"
	"regexp"
)

const (
	//Version 1.0 beta
	Version         = "1.0"
	ContentTypeJSON = "application/json"
)

type Handler func(*ApplicationContext, http.ResponseWriter, *http.Request) (int, interface{})

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
}

type Context struct {
	http.ResponseWriter
	*http.Request
}

//ApplicationContext struct
type ApplicationContext struct {
	Version   string
	Container infrastructure.Container
	Routes    []Route
	Params    []string
}

//New
func New() *ApplicationContext {
	return &ApplicationContext{
		Version: Version,
	}
}

func (ac *ApplicationContext) Handle(pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Handler: handler}
	ac.Routes = append(ac.Routes, route)
}

func (ah ApplicationContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := ah.dispatch(w, r)
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
	return
}

func (ah ApplicationContext) dispatch(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	ctx := &Context{Request: r, ResponseWriter: w}

	var statusCode = 200
	var data interface{}

	for _, rt := range ah.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				ah.Params = matches[1:]
			}
			statusCode, data = rt.Handler(&ah, w, r)
			return statusCode, data
		}
	}

	fmt.Printf("data: %#v", data)

	if data == nil {
		data = struct{ Response string }{Response: "Not found"}
		statusCode = http.StatusNotFound
	}
	return statusCode, data
}

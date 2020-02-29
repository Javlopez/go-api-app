//Package application is to manage the application state, transform handler is routes and executer them as server
package application

import (
	"encoding/json"
	"go-lana/pkg/infra"
	jsonResponse "go-lana/pkg/response/json"
	"net/http"
	"regexp"
	"sync"
)

const (
	//Version of application
	Version = "1.0"
	//ContentTypeJSON The application for now only is able to response JSON content
	ContentTypeJSON = "application/json"
)

//Handler is a type able to manage custom handler keeping http.ResponseWriter, *http.Request
type Handler func(*ApplicationContext, http.ResponseWriter, *http.Request) *jsonResponse.Response

//Route is a struct to manage each handler
type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
	Methods []string
}

//Context is a struct to manage the context
type Context struct {
	http.ResponseWriter
	*http.Request
}

//ApplicationContext struct for manbage the application state
type ApplicationContext struct {
	sync.RWMutex
	Version   string
	Container infra.Container
	Routes    []Route
	Params    []string
	Context   *Context
}

//New this function instanciates a new ApplicationContext
func New() *ApplicationContext {
	return &ApplicationContext{
		Version: Version,
	}
}

//Handle method is to create routes from handlers
func (ac *ApplicationContext) Handle(pattern string, handler Handler, methods ...string) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Handler: handler, Methods: methods}
	ac.Routes = append(ac.Routes, route)
}

//ServeHTTP method is to make the application able to create a server
func (ac *ApplicationContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	ac.Context = &Context{Request: r, ResponseWriter: w}
	resp := ac.Dispatch(w, r)
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(resp.Code)

	if resp.Success {
		output["data"] = resp.Data
	} else {
		output["error"] = resp.Data
	}

	json.NewEncoder(w).Encode(output)
	return
}

//Dispatch method is to validate and execute the handler is it exists otherwise will be returned an 404 error
func (ac *ApplicationContext) Dispatch(w http.ResponseWriter, r *http.Request) *jsonResponse.Response {
	ctx := ac.Context

	for _, rt := range ac.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {

			if !ac.MethodIsAllowed(rt.Methods, ctx.Method) {
				return jsonResponse.NewErrorResponse(http.StatusMethodNotAllowed, "Method Not Allowed")
			}

			if len(matches) > 1 {
				ac.Params = matches[1:]
			}
			return rt.Handler(ac, w, r)
		}
	}

	return jsonResponse.NewErrorResponse(http.StatusNotFound, "Url not found")
}

//MethodIsAllowed validates is the method used by the request is allowed or not
func (ac *ApplicationContext) MethodIsAllowed(methods []string, method string) bool {
	for _, methodAllowed := range methods {
		if methodAllowed == method {
			return true
		}
	}
	return false
}

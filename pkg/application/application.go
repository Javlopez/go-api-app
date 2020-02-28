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
	//Version 1.0 beta
	Version         = "1.0"
	ContentTypeJSON = "application/json"
)

type Handler func(*ApplicationContext, http.ResponseWriter, *http.Request) *jsonResponse.Response

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
	Methods []string
}

type Context struct {
	http.ResponseWriter
	*http.Request
}

//ApplicationContext struct
type ApplicationContext struct {
	sync.RWMutex
	Version   string
	Container infra.Container
	Routes    []Route
	Params    []string
	Context   *Context
}

//New
func New() *ApplicationContext {
	return &ApplicationContext{
		Version: Version,
	}
}

//Handle method
func (ac *ApplicationContext) Handle(pattern string, handler Handler, methods ...string) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Handler: handler, Methods: methods}
	ac.Routes = append(ac.Routes, route)
}

func (ah ApplicationContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	ah.Context = &Context{Request: r, ResponseWriter: w}
	resp := ah.dispatch(w, r)
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

func (ah ApplicationContext) dispatch(w http.ResponseWriter, r *http.Request) *jsonResponse.Response {
	ctx := ah.Context

	for _, rt := range ah.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {

			if !ah.MethodIsAllowed(rt.Methods, ctx.Method) {
				return jsonResponse.NewErrorResponse(http.StatusMethodNotAllowed, "Method Not Allowed")
			}

			if len(matches) > 1 {
				ah.Params = matches[1:]
			}
			return rt.Handler(&ah, w, r)
		}
	}

	return jsonResponse.NewErrorResponse(http.StatusNotFound, "Url not found")
}

func (ah ApplicationContext) MethodIsAllowed(methods []string, method string) bool {
	for _, methodAllowed := range methods {
		if methodAllowed == method {
			return true
		}
	}
	return false
}

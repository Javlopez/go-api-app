package application_test

import (
	app "go-lana/pkg/application"
	"go-lana/pkg/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func FakeHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return 200, nil
}

func FakeHandlerFail(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return 404, nil
}

// The application app is able to dispatch any handler by using Handle() method
// also it can manage any method
func Example_app() {
	ap := app.New()
	ap.Handle("/endpoint/([A-Z]+)$", handlers.ProductHandler, "GET", "POST", "PUT")
	err := http.ListenAndServe(":8000", ap)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func Example_route() {
	var routes []app.Route
	pattern := "/endpoint/"
	methods := []string{"GET", "POST", "PUT"}
	re := regexp.MustCompile(pattern)
	route := app.Route{Pattern: re, Handler: handlers.ProductHandler, Methods: methods}
	routes = append(routes, route)
}

func TestApplicationHandlers(t *testing.T) {

	t.Run("Handler should show not found when the handler does not exists", func(t *testing.T) {
		a := app.New()
		r := httptest.NewRequest("GET", "/not_exists", nil)
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusNotFound {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusNotFound, status)
		}
	})

	t.Run("Handler should accept only methods allowed", func(t *testing.T) {
		a := app.New()
		req := httptest.NewRequest("POST", "/products/", nil)
		rr := httptest.NewRecorder()

		a.Handle("/products/", handlers.ProductsHandler, "GET")

		a.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusMethodNotAllowed, status)
		}
	})

	t.Run("Handler should show sucess when the handler does exists", func(t *testing.T) {
		a := app.New()
		req := httptest.NewRequest("GET", "/products/", nil)
		rr := httptest.NewRecorder()

		a.Handle("/products/", handlers.ProductsHandler, "GET")

		a.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})

	t.Run("Handler should be able to manage parameters", func(t *testing.T) {
		a := app.New()
		req := httptest.NewRequest("GET", "/products/PEN", nil)
		rr := httptest.NewRecorder()

		a.Handle("/products/([A-Z]+)$", handlers.ProductHandler, "GET")

		a.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})

}

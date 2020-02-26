package application_test

import (
	app "go-lana/pkg/application"
	"go-lana/pkg/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FakeHandler(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return 200, nil
}

func FakeHandlerFail(a *app.ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return 404, nil
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

	t.Run("Handler should show sucess when the handler does exists", func(t *testing.T) {
		a := app.New()
		req := httptest.NewRequest("GET", "/products/", nil)
		rr := httptest.NewRecorder()

		a.Handle("/products/", handlers.ProductsHandler)

		a.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})

	t.Run("Handler should be able to manage parameters", func(t *testing.T) {
		a := app.New()
		req := httptest.NewRequest("GET", "/products/PEN", nil)
		rr := httptest.NewRecorder()

		a.Handle("/products/([A-Z]+)$", handlers.ProductHandler)

		a.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})

}

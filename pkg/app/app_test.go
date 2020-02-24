package app

import (
	"go-lana/pkg/infrastructure"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FakeHandler(a *ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return 200, nil
}

func FakeHandlerFail(a *ApplicationContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return 404, nil
}

func TestServeHTTP(t *testing.T) {

	t.Run("ServeHTTP should be able to manage request", func(t *testing.T) {
		container := infrastructure.Container{}
		a := &ApplicationContext{Version: "v0.1.0", Container: container}
		ha := ApplicationHandler{a, FakeHandler}

		r := httptest.NewRequest("GET", "/foo", nil)
		w := httptest.NewRecorder()
		ha.ServeHTTP(w, r)
		if status := w.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})

	t.Run("ServeHTTP should be able to manage invalid requests", func(t *testing.T) {
		container := infrastructure.Container{}
		a := &ApplicationContext{Version: "v0.1.0", Container: container}
		ha := ApplicationHandler{a, FakeHandlerFail}

		r := httptest.NewRequest("GET", "/foo", nil)
		w := httptest.NewRecorder()
		ha.ServeHTTP(w, r)
		if status := w.Code; status != http.StatusNotFound {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})
}

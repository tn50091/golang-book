package home

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET","/",nil)
	HomePageHandler(res, req)

	if res.Code != 200 {
		t.Fatalf("Expect status to == 200, but got %d", res.Code)
	}
}

func TestHomePageHandlerShouldReturnHelloWorld(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET","/",nil)
	HomePageHandler(res, req)

	if res.Body.String() != "Hello World!" {
		t.Fatalf("Expect body to be Hello World!, but got %s", res.Body.String())
	}
}


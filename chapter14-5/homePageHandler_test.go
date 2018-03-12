package home

import (
	"encoding/json"
	"strings"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST","/json",strings.NewReader(`{"first_name": "Suriya","LastName": "Dubey","Email": "tn@tnis.com"}`))
	home := HomePageHandler{}
	home.ServeHTTP(res,req)

	if res.Code != 201 {
		t.Fatalf("Expect status to == 201, but got %d", res.Code)
	}

	user := new(User)
	json.NewDecoder(res.Body).Decode(user)

	if user.FirstName != "Suriya" {
		t.Fatalf("Expect user first to == Suriya, but got %s", user.FirstName)
	}

}

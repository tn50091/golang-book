package home

import (
	"net/http"
	"fmt"
)

func HomePageHandler (res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}
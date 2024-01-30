package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	RestApi(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

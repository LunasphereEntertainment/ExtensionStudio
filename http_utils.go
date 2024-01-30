package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type httpErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func serialize(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

func deserialize[T interface{}](w http.ResponseWriter, r *http.Request) *T {
	out := new(T)

	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		httpError(w, http.StatusBadRequest, err)
	}

	return out
}

func httpError(w http.ResponseWriter, status int, errs ...error) {
	err := httpErrorResponse{
		Status:  status,
		Message: errs[0].Error(),
	}

	if len(errs) > 1 {
		err.Detail = errs[1].Error()
	}

	serialize(w, err)
}

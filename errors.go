package main

import (
	"fmt"
	"net/http"
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	err := writeJSON(w, status, response{"error": message})
	if err != nil {
		w.WriteHeader(500)
	}
}

func internalErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println("error: ", err.Error())
	message := "Internal Server Error"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "Resource Not Found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

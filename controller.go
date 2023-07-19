package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func RouteBuilder() *mux.Router{
    r := mux.NewRouter()
    r.HandleFunc("/health-check", HealthCheck)

    return r
}

func HealthCheck(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "OK")
}
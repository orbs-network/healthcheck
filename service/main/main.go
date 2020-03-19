package main

import (
	"github.com/orbs-network/healthcheck/service"
	"net/http"
)

func main() {
	service.Dummy()

	http.ListenAndServe(":8080", nil)
}

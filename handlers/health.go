package handlers

import (
	"fmt"
	"net/http"
)

type HealthHandler struct {

}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

package handlers

import (
	"bytes"
	"log"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	recorder := httptest.NewRecorder()
	recorder.Body = bytes.NewBuffer(nil)

	handler := HealthHandler{}

	handler.ServeHTTP(recorder, nil)

	if recorder.Body.String() != "pong" {
		log.Fatalf("response does not match expected")
	}
}

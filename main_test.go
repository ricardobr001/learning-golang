package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	RequestResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler)
	handler.ServeHTTP(RequestResponse, request)

	expected := `{"message":"Hello World"}`
	if RequestResponse.Body.String() != expected {
		t.Errorf(
			`handler returned unexpected body: got "%v" want "%v"`,
			RequestResponse.Body.String(),
			expected,
		)
	}
}

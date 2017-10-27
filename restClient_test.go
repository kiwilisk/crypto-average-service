package main

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
)

func TestParseAndReturnResponse(t *testing.T) {
	expectedBody := "{\"someKey\": \"someValue\"}"
	echoHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedBody)
	}
	ts := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer ts.Close()
	restClient := DefaultRestClient{http.Client{}}

	response, err := restClient.Get(ts.URL)

	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if response.StatusCode != 200 {
		t.Fatalf("unsuccesful request: %d", response.StatusCode)
	}
	body := string(response.Body)
	if body != expectedBody {
		t.Fatalf("body did not match")
	}
}

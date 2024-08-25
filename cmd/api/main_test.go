package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "status: available"
	var responseFirstLine string

	req, err := http.NewRequest(http.MethodGet, "http://localhost:4000/v1/healthcheck", nil)
	if err != nil {
		t.Errorf("error making http request: %s\n", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("client: error making http request: %s\n", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if i == 0 {
			responseFirstLine = line
		}
	}

	if responseFirstLine != expected {
		t.Errorf("Expected %v but got %v", string(expected), string(responseFirstLine))
	}
}

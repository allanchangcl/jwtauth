package main

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// expected := "status: available"
	// defer os.Unsetenv("ES_")
	// os.Setenv("ES_URL", "localhost:4000/v1/healthcheck")

	// req := httptest.NewRequest(http.MethodGet, "localhost:4000/v1/healthcheck", nil)

	// Create a response recorder so you can inspect the response
	// w := httptest.NewRecorder()
	// healthcheckHandler(w, req)
	// res := w.Result()
	// defer res.Body.Close()

	// Perform the request
	// req.ServeHTTP(w, req)
	// fmt.Println(res)

	// data, err := io.ReadAll(res.Body)

	// fmt.Println(req)

	// if string(data) != expected {
	// 	t.Errorf("expected ABC got %v", string(data))
	// }
	req, err := http.NewRequest(http.MethodGet, "http://localhost:4000/v1/healthcheck", nil)
	if err != nil {
		log.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)
	log.Printf("client: got response!\n")
	log.Printf("client: status code: %d\n", res.StatusCode)

	log.Println(res.StatusCode)

	// os.Exit(t.Run())
}

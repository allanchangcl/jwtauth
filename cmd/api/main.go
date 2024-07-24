package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare a string containing the application version number
const version = "1.0.0"

// Define a config struct.
type config struct {
	env  string
	port int
}

// Define an application struct to hold dependencides for our HTTP handlers, helpers, and
// middleware.
type application struct {
	logger *log.Logger
	config config
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

//	func stripSlashes(next http.Handler) http.Handler {
//	  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	}
//
// func (app *application) stripSlash(next http.Handler) http.Handler {
func stripSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("OK")
		log.Print("Executing middlewareOne")
		path := r.URL.Path

		if len(path) > 1 && path[len(path)-1] == '/' {
			newPath := path[:len(path)-1]
			r.URL.Path = newPath
		}

		fmt.Println(path)
		next.ServeHTTP(w, r)
	})
}

// func (app *application) routes() http.Handler {
// 	// mux := http.NewServeMux()
// 	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
//
// 	// mux.HandleFunc("GET /{$}", app.home)
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Server", "Go")
// 		mux.ServeHTTP(w, r)
// 	})
// }

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	// We default to using the port number 4000 and the environment "development" if no
	// corresponding flags are provided.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production")
	flag.Parse()

	// Initialize a new logger which writes messages to the STDOUT stream, prefixed
	// with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of the application struct, containing the config struct and the logger.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a new servemux and add a /v1/healthcheck route which dispatches requests
	// to the healthcheckHandler method.
	mux := http.NewServeMux()

	// HandleFunc expects a function where Handle expects a Handler interface.
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Declare an HTTP server with some sensitive timeout settings, which listens on the
	// on the post provided in the config struct and uses the servemux we created above.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      stripSlash(mux),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("starting the %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {

	// We will register one function for our HTTP server
	// When a http request is sent to /, the following function will be executed
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		m := map[string]interface{}{
			"hello": "Cisco Live!",
		}
		data, _ := json.Marshal(&m)
		w.Write(data)
	})

	// We initialize our HTTP server with some sane defaults
	// Handler is set to nil, because we will you the default ServerMux provided by the net/http package
	s := &http.Server{
		Addr:           "localhost:8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// log any fatal errors
	log.Fatal(s.ListenAndServe())
}

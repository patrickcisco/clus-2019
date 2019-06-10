// https://github.com/julienschmidt/httprouter
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Echo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	for k, v := range r.Header {
		fmt.Println(k, v)
	}

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/echo", Echo)
	log.Fatal(http.ListenAndServe("localhost:8070", router))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "Hello World!",
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, string(b))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	resp := struct {
		Greeting string `json:"greeting"`
	}{
		Greeting: fmt.Sprintf("Hello, %s!", ps.ByName("name")),
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, string(b))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

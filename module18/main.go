package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Understading MOD in go")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey There!")
}

func serverHome(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("<h1>Kasa Kaaye</h1>"))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo Get")
}

func postUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo Post")
}

func putUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo Put")
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo Delete")
}

func main() {

	// enrutador
	r := mux.NewRouter().StrictSlash(false) //con false es ditinto api/algo que api/algo/ (barra al final)

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", postUsers).Methods("POST")
	r.HandleFunc("/api/users", putUsers).Methods("PUT")
	r.HandleFunc("/api/users", deleteUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("escuchando...")
	server.ListenAndServe()
}

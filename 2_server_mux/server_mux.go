package main

import (
	"fmt"
	"net/http"
)

func main() {
	// crea el server mux
	mux := http.NewServeMux()

	// manejo de peticiones
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola mundo</h1>")
	})

	// ruta desde /Prueba
	mux.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola mundo desde /prueba</h1>")
	})

	// ruta desde /usuario
	mux.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola mundo desde /usuario</h1>")
	})

	// pasamos el mux
	http.ListenAndServe(":8080", mux)
}

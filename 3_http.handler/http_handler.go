package main

import (
	"fmt"
	"net/http"
)

func holaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo</h1>")
}

func prueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo desde /prueba</h1>")
}

func usuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo desde /usuario</h1>")
}

func main() {
	// crea el server mux
	mux := http.NewServeMux()

	// Se le pasa el objeto, sin parentesis para que no se ejecute
	mux.HandleFunc("/", holaMundo)

	// ruta desde /Prueba
	mux.HandleFunc("/prueba", prueba)

	// ruta desde /usuario
	mux.HandleFunc("/usuario", usuario)

	// pasamos el mux
	http.ListenAndServe(":8080", mux)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// maneja una peticion
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola mundo</h1>")
	})

	// crear el server
	http.ListenAndServe(":8080", nil)
}

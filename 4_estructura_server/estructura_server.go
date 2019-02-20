package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

// Manejador (usa la funcion ServerHTTP)
type mensaje struct {
	msj string
}

// el recibidor asocia la funcion a la estructura "mensaje"
func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msj)
}

func main() {
	msj := mensaje{
		msj: "Hola mundo de nuevo",
	}

	// crea el server mux
	mux := http.NewServeMux()

	mux.HandleFunc("/", holaMundo)
	mux.HandleFunc("/prueba", prueba)
	mux.HandleFunc("/usuario", usuario)

	mux.Handle("/hola", msj)

	// estructura server propia
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second, //tiempo de espera para escribir
		WriteTimeout:   10 * time.Second, //tiempo de espera para leer antes de cortar la peticion
		MaxHeaderBytes: 1 << 20,          //operador shift, multiplica por 2 el Nro de la izquierda (1) la cantidad de veces de la derecha (20)
	}

	log.Println("Escuchando...")
	// server.ListenAndServe() //asi funciona

	log.Fatal(server.ListenAndServe()) //asi corta el server si hay error
}

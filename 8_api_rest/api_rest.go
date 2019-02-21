package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Note : estructura para simular la entidad a guardar en una "BD"
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// map para almacenar las notas, simula la tabla de una BD
var noteStore = make(map[string]Note)

// id del map
var id int

// GetNoteHandler - GET - /api/notes
// recorre el map de notas y lo pasa a json
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// objeto que va a contener las notas
	var notes []Note //slice de Note

	for _, value := range noteStore {
		notes = append(notes, value)
	}

	// cabecera para la solicitud
	// el seteo de la cabecera debe ser ANTES que WriteHeader (sino WriteHeader no lo tiene en cuenta)
	w.Header().Set("Content-Type", "application/json")

	// convertir a json
	j, err := json.Marshal(notes) //devuelve un slice de byte

	if err != nil {
		// no usar panic en produccion
		panic(err)
	}

	// agrega codigo de status a la respuesta
	w.WriteHeader(http.StatusOK) //200

	// cuerpo de la respuesta
	w.Write(j)
}

// PostNoteHandler - POST - /api/notes
// toma el json de la peticion y lo decodifica a la estructura Note
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	// objeto que va a contener la nota
	var note Note

	// decodifica el json (el json viene en el body del request -r-)
	err := json.NewDecoder(r.Body).Decode(&note) //con el puntero, lo decodificado se guarda ennote directamente

	// maneja el error si lo hay
	if err != nil {
		panic(err)
	}

	// agrega la fecha al Note
	note.CreatedAt = time.Now()

	// incrementa el id
	id++

	// pasa el int a string
	k := strconv.Itoa(id)

	// guarda la nota en el map de almacenamiento (en un indice en particular)
	noteStore[k] = note

	// respuesta json
	w.Header().Set("Content-Type", "application/json")

	// convertir a json
	j, err := json.Marshal(note) //devuelve un slice de byte

	if err != nil {
		// no usar panic en produccion
		panic(err)
	}

	// agrega codigo de status a la respuesta
	w.WriteHeader(http.StatusCreated) //201

	// cuerpo de la respuesta
	w.Write(j)

}

// PutNoteHandler - PUT - /api/notes
// toma el json de la peticion, lo decodifica y busca por el paramaetro, lo actualiza
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	// extraer los parametros de la ruta
	params := mux.Vars(r) //devuelve un slice

	// id de la nota
	k := params["id"]

	// mantiene los datos enviados para actualizar
	var noteUpdate Note

	// extrae lo enviado en el request
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)

	// manejo de error
	if err != nil {
		panic(err)
	}

	// sentencia antes de conficion (ok)
	if buscarNota, ok := noteStore[k]; ok {
		// mantiene la fecha de creacion
		noteUpdate.CreatedAt = buscarNota.CreatedAt

		// borra la nota anterior
		delete(noteStore, k)

		// pone la nota nueva (actualizada)
		noteStore[k] = noteUpdate
	} else {
		log.Printf("No se encontro la nota con el id %s", k)
	}

	// respuesta de id no encontrado
	w.WriteHeader(http.StatusNoContent) // 204 //se devuelve por convencion en update y delete, que no hay contenido que devolver
}

// DeleteNoteHandler - DELETE - /api/notes
// borra la nota
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// extraer los parametros de la ruta
	params := mux.Vars(r) //devuelve un slice

	// id de la nota
	k := params["id"]

	// sentencia antes de conficion (ok)
	if _, ok := noteStore[k]; ok {
		// borra la nota anterior
		delete(noteStore, k)
	} else {
		log.Printf("No se encontro la nota con el id %s", k)
	}

	// respuesta de id no encontrado
	w.WriteHeader(http.StatusNoContent) // 204 //se devuelve por convencion en update y delete, que no hay contenido que devolver
}

func main() {
	// enrutador
	r := mux.NewRouter().StrictSlash(false)

	// manejadores
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// muestra este texto en la consola cuando se ejecuta
	log.Println("Escuchando desde http://localhost:8080 ...")

	// inicia el server
	server.ListenAndServe()
}

package main

import (
	"html/template"
	"os"
)

// el template recibe data para renderizar donde se le indique

// Persona : estructura para data
type Persona struct {
	Nombre string
	Edad   int
	Pais   string
}

func main() {
	// creo una persona
	persona := Persona{Nombre: "Pablo", Edad: 27, Pais: "Argentina"}

	// creo el template
	t := template.New("nombre del template")

	// parsea la plantilla al template creado
	t, err := t.ParseGlob("templates/*.txt")

	if err != nil {
		panic(err)
	}

	// ejecutar el template (ver en consola)
	err = t.ExecuteTemplate(os.Stdout, "visitante", persona)
	if err != nil {
		panic(err)
	}
}

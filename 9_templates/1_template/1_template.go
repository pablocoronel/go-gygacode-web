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
}

// plantilla para el template
const tp = `Nombre: {{.Nombre}} Edad: {{.Edad}}`

func main() {
	// creo una persona
	persona := Persona{Nombre: "Pablo", Edad: 27}

	// creo el template
	t := template.New("nombre del template")

	// parsea la plantilla al template creado
	t, err := t.Parse(tp)

	if err != nil {
		panic(err)
	}

	// ejecutar el template (ver en consola)
	err = t.Execute(os.Stdout, persona)
	if err != nil {
		panic(err)
	}
}

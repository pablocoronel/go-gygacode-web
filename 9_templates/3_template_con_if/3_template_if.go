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
// el punto (.) es el slice que recibe, para iterarlo
const tp = `
{{range .}}
	{{if .Edad}}
		Nombre: {{.Nombre}} Edad: {{.Edad}} - Correcto
	{{else if .Nombre}}
		Nombre: {{.Nombre}} Edad: {{.Edad}} - Incorrecto
	{{else}}
		Nombre: {{.Nombre}} Edad: {{.Edad}} - Está vacio
	{{end}}
{{end}}`

func main() {
	// creo una persona
	persona := []Persona{
		{Nombre: "Pablo", Edad: 27},
		{Nombre: "Nadia", Edad: 0},
		{Nombre: "", Edad: 0},
	}

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

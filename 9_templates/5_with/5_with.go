package main

import (
	"html/template"
	"os"
)

// template
const hero = `
Hero Name: {{.Name}}
{{range .Emails}}
	Email: {{.}}
{{end}}

{{with .Friends}}
	{{range .}}
		Friend name: {{.Name}}
	{{end}}
{{end}}
`

// Friend ...
type Friend struct {
	Name string
}

// Hero ...
type Hero struct {
	Name    string
	Emails  []string
	Friends []Friend
}

func main() {

	f1 := Friend{"Luciana"}
	f2 := Friend{"Lucas"}

	t := template.New("le pongo un nombre al template")

	t, err := t.Parse(hero)
	if err != nil {
		panic(err)
	}

	hero := Hero{Name: "Pablo",
		Emails:  []string{"pablo@gmail.com", "otroemail@gmail.com"},
		Friends: []Friend{f1, f2},
	}

	err = t.Execute(os.Stdout, hero)
	if err != nil {
		panic(err)
	}
}

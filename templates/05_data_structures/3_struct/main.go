package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}

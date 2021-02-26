package main

import (
	"log"
	"os"
	"text/template"
)

func parseToStdout() {
	// *Template is a container holding all the templates that are parsed.
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func parseToHTMLFile() {
	// *Template is a container holding all the templates that are parsed.
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseMultipleTemplates() {
	// *Template is a container holding all the templates that are parsed.
	tpl, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "templates/one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "templates/two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func parseGlob() {
	// *Template is a container holding all the templates that are parsed.
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	// parseToStdout()
	// parseToHTMLFile()
	// parseMultipleTemplates()
	parseGlob()
}

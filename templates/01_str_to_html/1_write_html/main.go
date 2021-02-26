package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Todd McLeod"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	<body>
	<h1>` + name + `</h1>
	</body>
	</head>
	</html>
	`
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
}

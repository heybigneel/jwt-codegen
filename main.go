package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/neel229/jwt-codegen/codegen"
)

func main() {
	// 1. read example.json file
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}
	// 2. parse values in struct
	var cc codegen.CustomClaims
	if err = json.Unmarshal(data, &cc); err != nil {
		log.Fatalf("error unmarshaling data: %v\n", err)
	}

	// 3. generate jwt file from template
	tmpl, err := template.ParseFiles("./codegen/jwt.tmpl")
	if err != nil {
		log.Fatalf("error parsing template file: %v\n", err)
	}
	var processed bytes.Buffer
	if err = tmpl.Execute(&processed, cc); err != nil {
		log.Fatalf("error executing template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("error formatting code: %v\n", err)
	}

	file, err := os.Create("jwt.go")
	if err != nil {
		log.Fatalf("error creating file: %v\n", err)
	}
	_, err = file.Write(formatted)
	if err != nil {
		log.Fatalf("error writing data to file: %v\n", err)
	}
}

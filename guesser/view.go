package guesser

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, tmpl string,r *http.Request) {
	addOn := fmt.Sprintf("public/%s",tmpl)
	t, err := template.ParseFiles(addOn)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Insult struct {
	InsultText string
}

//go:embed index.html
var f embed.FS

func main() {

	fmt.Println("Server running: http://localhost:8000")
	fmt.Println("CTRL +C to stop server")

	httpGetInsult := func(w http.ResponseWriter, r *http.Request) {
		// alternative to read from local filesystem
		// tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl := template.Must(template.ParseFS(f, "index.html"))
		tmpl.Execute(w, Insult{InsultText: getInsult()})
	}

	httpGetInsultUpdate := func(w http.ResponseWriter, r *http.Request) {
		// tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl := template.Must(template.ParseFS(f, "index.html"))
		tmpl.ExecuteTemplate(w, "insult-element", Insult{InsultText: getInsult()})
	}

	http.HandleFunc("/", httpGetInsult)
	http.HandleFunc("/get/", httpGetInsultUpdate)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

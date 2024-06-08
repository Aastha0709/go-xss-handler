package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template = template.Must(template.ParseFiles("./templates/input.html"))
var p *template.Template = template.Must(template.ParseFiles("./templates/parsed.html"))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", submit)
	http.HandleFunc("/unsafeparsed", unsafeparsed)
	http.HandleFunc("/safeparsed", safeparsed)

	fmt.Println("Listening on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Invalid path", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/submit", http.StatusSeeOther)
}

func submit(w http.ResponseWriter, r *http.Request) {
	t.Execute(w, nil)
}

func unsafeparsed(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	userInput := r.FormValue("textfield")
	p.Execute(w, template.HTML(userInput))
}

func safeparsed(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	userInput := r.FormValue("textfield")
	p.Execute(w, userInput)
}

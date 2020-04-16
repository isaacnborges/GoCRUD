package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "lala", Preco: 15, Quantidade: 45},
		{"Tenis", "aaasdsa", 199, 15},
		{"Mouse", "wqewq", 59, 40},
		{"Tecladada", "ewqeqw", 999, 12},
	}

	templates.ExecuteTemplate(w, "index", produtos)
}

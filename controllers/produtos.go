package controllers

import (
	"GoCRUD/models"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.ListarProdutos()
	templates.ExecuteTemplate(w, "index", produtos)
}

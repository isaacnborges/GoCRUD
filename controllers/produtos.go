package controllers

import (
	"GoCRUD/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Index é rota principal
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.ListarProdutos()
	templates.ExecuteTemplate(w, "index", produtos)
}

// New é rota para a tela de produto
func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new", nil)
}

// Insert é rota para obter os dados do formulario
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		precoStr := r.FormValue("preco")
		quantidadeStr := r.FormValue("quantidade")

		preco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			log.Println("Erro conversão preço")
		}

		quantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			log.Println("Erro conversão quantidade")
		}

		models.CriarProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}

// Delete é para obter o id da rota
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.ExcluirProduto(idProduto)

	http.Redirect(w, r, "/", 301)
}

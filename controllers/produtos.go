package controllers

import (
	"GoCRUD/models"
	"fmt"
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

//Get para obter o id da rota e enviar para o model editar
func Get(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.ObterProduto(idProduto)

	templates.ExecuteTemplate(w, "get", produto)
}

func acaoProduto(r *http.Request, acao string) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		precoStr := r.FormValue("preco")
		quantidadeStr := r.FormValue("quantidade")

		preco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			fmt.Println("Erro conversão preço")
		}

		quantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			fmt.Println("Erro conversão quantidade")
		}

		if acao == "insert" {
			models.CriarProduto(nome, descricao, preco, quantidade)
		} else {
			idStr := r.FormValue("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Erro conversão id")
			}

			models.EditarProduto(id, nome, descricao, preco, quantidade)
		}
	}
}

// Insert é rota para obter os dados do formulario
func Insert(w http.ResponseWriter, r *http.Request) {
	acaoProduto(r, "insert")

	http.Redirect(w, r, "/", 301)
}

// Update é rota para obter os dados do formulario e editar o produto
func Update(w http.ResponseWriter, r *http.Request) {
	acaoProduto(r, "update")
	http.Redirect(w, r, "/", 301)
}

// Delete é para obter o id da rota
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.ExcluirProduto(idProduto)

	http.Redirect(w, r, "/", 301)
}

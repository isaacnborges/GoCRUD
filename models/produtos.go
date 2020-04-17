package models

import (
	"GoCRUD/db"
	"fmt"
)

// Produto é representação da tabela no banco
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func ListarProdutos() []Produto {
	db := db.ConectarDB()

	listarProdutos, err := db.Query("SELECT * FROM Produtos")
	if err != nil {
		fmt.Println(err.Error())
	}

	produtos := []Produto{}
	p := Produto{}

	for listarProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = listarProdutos.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			fmt.Println(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

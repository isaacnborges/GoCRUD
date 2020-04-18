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

// ListarProdutos serve para listar os produtos do DB
func ListarProdutos() []Produto {
	db := db.ConectarDB()

	produtosDb, err := db.Query("SELECT * FROM Produtos ORDER BY Descricao")
	if err != nil {
		fmt.Println(err.Error())
	}

	produtos := []Produto{}
	p := Produto{}

	for produtosDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtosDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			fmt.Println(err.Error())
		}

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

// ObterProduto para obter um produto específico
func ObterProduto(id string) Produto {
	db := db.ConectarDB()

	produtoDb, err := db.Query("SELECT * FROM Produtos WHERE id = $1", id)
	if err != nil {
		fmt.Println(err.Error())
	}

	produto := Produto{}

	for produtoDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			fmt.Println(err.Error())
		}

		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()
	return produto
}

// CriarProduto insere no banco
func CriarProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectarDB()

	inserirProduto, err := db.Prepare("INSERT INTO Produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		fmt.Println(err.Error())
	}

	inserirProduto.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

// EditarProduto editar um produto no banco
func EditarProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectarDB()

	editarProduto, err := db.Prepare("UPDATE Produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		fmt.Println(err.Error())
	}

	editarProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

// ExcluirProduto exclui o produto do banco
func ExcluirProduto(id string) {
	db := db.ConectarDB()

	deletarProduto, err := db.Prepare("DELETE FROM Produtos WHERE id = $1")
	if err != nil {
		fmt.Println(err.Error())
	}

	deletarProduto.Exec(id)
	defer db.Close()
}

# GoCRUD
Projeto CRUD utilizando Golang


Comando para rodar o projeto 
- go run main.go

Necessário instalar o bando de dados postgresql

Base de dados: loja_demo

CREATE TABLE Produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)


Libs utilizadas:
go get github.com/lib/pq
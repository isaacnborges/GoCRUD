# GoCRUD
Projeto CRUD utilizando Golang

Está rodando na porta 8000
http://localhost:8000/

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
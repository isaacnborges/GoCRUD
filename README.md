# GoCRUD
Projeto CRUD utilizando Golang

Porta 8000 para rodar o projeto: http://localhost:8000/

Comando para rodar o projeto 
- go run main.go

Necessário instalar o bando de dados postgresql

Administrador Banco de dados: pgAdmin

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

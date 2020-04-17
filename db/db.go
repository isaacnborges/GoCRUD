package db

import (
	"database/sql"

	// Utilizado para realizar a conexao com banco
	_ "github.com/lib/pq"
)

func ConectarDB() *sql.DB {
	conexao := "user=postgres dbname=loja_demo password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

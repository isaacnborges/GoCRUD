package db

import (
	"database/sql"
)

// ConectarDB é conexão com banco postgres
func ConectarDB() *sql.DB {
	conexao := "user=postgres dbname=loja_demo password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

package main

import (
	"GoCRUD/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}

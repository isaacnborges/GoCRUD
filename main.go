package main

import (
	"GoCRUD/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregarRota()
	http.ListenAndServe(":8000", nil)
}

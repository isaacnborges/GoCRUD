package routes

import (
	"GoCRUD/controllers"
	"net/http"
)

func CarregarRota() {
	http.HandleFunc("/", controllers.Index)
}

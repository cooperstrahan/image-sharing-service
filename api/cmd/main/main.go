package main

import (
	"log"
	"net/http"

	"github.com/cooperstrahan/image-sharing-service/api/pkg/routes"
	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/cooperstrahan/image-sharing-service/api/pkg/models"
)

func main() {
	models.ResetDatabase()

	r := mux.NewRouter()
	routes.RegisterImageRoutes(r)
	http.Handle("/images", r)
	log.Fatal(http.ListenAndServe("localhost:5000", r))
}

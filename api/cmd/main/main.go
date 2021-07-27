package main

import (
	"log"
	"net/http"

	"github.com/cooperstrahan/image-sharing-service/api/pkg/routes"
	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterImageRoutes(r)
	http.Handle("/images", r)
	log.Fatal(http.ListenAndServe("localhost:5000", r))
}

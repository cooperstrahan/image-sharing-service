package routes

import (
	"github.com/cooperstrahan/image-sharing-service/api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterImageRoutes = func(router *mux.Router) {
	router.HandleFunc("/images", controllers.GetImages).Methods("GET")
	router.HandleFunc("/images/{id}", controllers.GetImageById).Methods("GET")
	router.HandleFunc("/images/", controllers.UploadImage).Methods("POST")
	router.HandleFunc("/images/{id}", controllers.DeleteImage).Methods("DELETE")
	router.HandleFunc("/images/{id}", controllers.IgnoreOptions).Methods("OPTIONS")
}

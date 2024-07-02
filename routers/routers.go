package routers

import (
	"github.com/gorilla/mux"
	controller "github.com/lucifer126741/mongo_app/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.AddMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")
	return router
}

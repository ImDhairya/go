package router

import (
	"github.com/ImDhairya/go/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMoveis).Methods("GET")

	router.HandleFunc("/api/creat-movie", controllers.CreateMovie).Methods("POST")

	router.HandleFunc("/api/watched/{id}", controllers.MarkedAsWatched).Methods("PUT")

	router.HandleFunc("/api/delete/{id}", controllers.DeleteAMovie).Methods("DELETE")

	router.HandleFunc("/api/delete-all", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}

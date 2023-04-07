package routers

import (
	"github.com/geekysaurabh001/gomongo/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controllers.InsertOneMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.UpdateOneMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovie).Methods("DELETE")
	return router
}

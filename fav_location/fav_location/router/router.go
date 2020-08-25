package router

import (
	"fav_location/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/fav_location/{id}", middleware.GetFavLoc).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/fav_location", middleware.GetAllFavLoc).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/fav_location", middleware.CreateFavLoc).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/fav_location/{id}", middleware.UpdateFavLoc).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/fav_location/{id}", middleware.DeleteFavLoc).Methods("DELETE", "OPTIONS")

	return router
}

package router

import (
	controller "api-basic/Controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", controller.HealthCheck).Methods("GET")
	r.HandleFunc("/api/movies", controller.GetAllMovie).Methods("GET")
	r.HandleFunc("/api/create", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/update", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/delete/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/delete-all", controller.DeleteAllMovie).Methods("DELETE")
	r.HandleFunc("/api/all-courses", controller.GetAllClasses).Methods("GET")

	return r
}

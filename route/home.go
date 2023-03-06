package route

import (
	"github.com/opensaucerer/gotemplate/controller"

	mux "github.com/gorilla/mux"
)

func RegisterHomeRoutes(r *mux.Router) {

	r.HandleFunc("/", controller.Home).Methods("GET")
}

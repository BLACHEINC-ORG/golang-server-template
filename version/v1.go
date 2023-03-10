package version

import (
	mux "github.com/gorilla/mux"
	"github.com/opensaucerer/gotemplate/route"
	"github.com/opensaucerer/gotemplate/route/v1/base"
)

// Version1Routes registers all routes for the v1 version
func Version1Routes(r *mux.Router) {

	//home does not need versioning
	route.RegisterHomeRoutes(r)

	// V1 routes
	router := r.PathPrefix("/v1").Subrouter()

	//health info for version 1
	base.RegisterHealthRoutes(router)

}

package service_instances

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

func NewRouter(serviceInstances *domain.ServiceInstances) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/service_instances", createHandler{serviceInstances}).Methods("POST")
	router.Handle("/v2/service_instances/{guid}", getHandler{serviceInstances}).Methods("GET")

	return router
}

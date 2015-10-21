package applications

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/invalid"
)

func NewRouter(apps *domain.Applications) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/apps", createHandler{apps}).Methods("POST")
	router.Handle("/v2/apps/very-bad-guid", invalid.Handler{}).Methods("GET")
	router.Handle("/v2/apps/{guid}", getHandler{apps}).Methods("GET")
	router.Handle("/v2/apps/{guid}", deleteHandler{apps}).Methods("DELETE")

	return router
}

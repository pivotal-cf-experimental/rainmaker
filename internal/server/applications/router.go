package applications

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/invalid"
)

type guidGenerator func(string) string

func NewRouter(guids guidGenerator, apps *domain.Applications) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/apps", createHandler{guids, apps}).Methods("POST")
	router.Handle("/v2/apps", listHandler{apps}).Methods("GET")
	router.Handle("/v2/apps/very-bad-guid", invalid.Handler{}).Methods("GET")
	router.Handle("/v2/apps/{guid}", getHandler{apps}).Methods("GET")
	router.Handle("/v2/apps/{guid}", deleteHandler{apps}).Methods("DELETE")

	return router
}

package spaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/invalid"
)

func NewRouter(orgs *domain.Organizations, spaces *domain.Spaces, users *domain.Users) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/spaces", createHandler{orgs, spaces}).Methods("POST")
	router.Handle("/v2/spaces", listHandler{spaces}).Methods("GET")
	router.Handle("/v2/spaces/very-bad-guid", invalid.Handler{}).Methods("GET")
	router.Handle("/v2/spaces/{guid}", getHandler{spaces}).Methods("GET")
	router.Handle("/v2/spaces/very-bad-guid", invalid.Handler{}).Methods("DELETE")
	router.Handle("/v2/spaces/{guid}", deleteHandler{spaces}).Methods("DELETE")

	router.Handle("/v2/spaces/{guid}/developers/{developer_guid}", associateDeveloperHandler{spaces, users}).Methods("PUT")
	router.Handle("/v2/spaces/{guid}/developers", getDevelopersHandler{spaces}).Methods("GET")

	return router
}

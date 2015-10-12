package spaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

func NewRouter(spaces *domain.Spaces, users *domain.Users) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/spaces", createHandler{spaces}).Methods("POST")
	router.Handle("/v2/spaces/{guid}", getHandler{spaces}).Methods("GET")
	router.Handle("/v2/spaces/{guid}", deleteHandler{spaces}).Methods("DELETE")

	router.Handle("/v2/spaces/{guid}/developers/{developer_guid}", associateDeveloperHandler{spaces, users}).Methods("PUT")
	router.Handle("/v2/spaces/{guid}/developers", getDevelopersHandler{spaces}).Methods("GET")

	return router
}

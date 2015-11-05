package buildpacks

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

func NewRouter(buildpacks *domain.Buildpacks) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/buildpacks", createHandler{buildpacks}).Methods("POST")

	return router
}

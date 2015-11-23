package buildpacks

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type guidGenerator func(string) string

func NewRouter(guids guidGenerator, buildpacks *domain.Buildpacks) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/buildpacks", createHandler{guids, buildpacks}).Methods("POST")
	router.Handle("/v2/buildpacks/{guid}", getHandler{buildpacks}).Methods("GET")
	router.Handle("/v2/buildpacks/{guid}", deleteHandler{buildpacks}).Methods("DELETE")

	return router
}

package buildpacks

import (
	"strings"

	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type deleteHandler struct {
	buildpacks *domain.Buildpacks
}

func (h deleteHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	guid := strings.TrimPrefix(path, "/v2/buildpacks/")
	h.buildpacks.Remove(guid)

	w.WriteHeader(http.StatusNoContent)
}

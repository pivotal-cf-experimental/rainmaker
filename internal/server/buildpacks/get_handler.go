package buildpacks

import (
	"strings"

	"encoding/json"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type getHandler struct {
	buildpacks *domain.Buildpacks
}

func (h getHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	guid := strings.TrimPrefix(path, "/v2/buildpacks/")
	buildpack, ok := h.buildpacks.Get(guid)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(buildpack)
}

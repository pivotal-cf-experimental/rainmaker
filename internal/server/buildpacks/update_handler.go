package buildpacks

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type updateHandler struct {
	buildpacks *domain.Buildpacks
}

func (h updateHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	var document documents.UpdateBuildpackRequest

	err := json.NewDecoder(request.Body).Decode(&document)
	if err != nil {
		common.JSONError(w, http.StatusBadRequest, common.ErrorResponse{
			Code:        1001,
			Description: "Request invalid due to parse error",
			ErrorCode:   "CF-MessageParseError",
		})
		return
	}

	path := request.URL.Path
	guid := strings.TrimPrefix(path, "/v2/buildpacks/")
	buildpack, ok := h.buildpacks.Get(guid)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if document.Name != nil {
		buildpack.Name = *document.Name
	}
	if document.Position != nil {
		buildpack.Position = *document.Position
	}
	if document.Enabled != nil {
		buildpack.Enabled = *document.Enabled
	}
	if document.Locked != nil {
		buildpack.Locked = *document.Locked
	}
	if document.Filename != nil {
		buildpack.Filename = *document.Filename
	}

	h.buildpacks.Add(buildpack)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buildpack)
}

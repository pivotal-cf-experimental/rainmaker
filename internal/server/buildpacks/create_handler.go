package buildpacks

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type createHandler struct {
	generateGUID guidGenerator
	buildpacks   *domain.Buildpacks
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	var document documents.CreateBuildpackRequest

	err := json.NewDecoder(request.Body).Decode(&document)
	if err != nil {
		common.JSONError(w, http.StatusBadRequest, common.ErrorResponse{
			Code:        1001,
			Description: "Request invalid due to parse error",
			ErrorCode:   "CF-MessageParseError",
		})
		return
	}

	buildpack := domain.NewBuildpack(h.generateGUID("buildpack"))
	buildpack.Name = document.Name
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

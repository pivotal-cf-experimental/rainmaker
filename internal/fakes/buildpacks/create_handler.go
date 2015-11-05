package buildpacks

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type createHandler struct {
	buildpacks *domain.Buildpacks
}

func (handler createHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	var document documents.CreateBuildpackRequest

	err := json.NewDecoder(request.Body).Decode(&document)
	if err != nil {
		panic(err)
	}

	buildpack := domain.NewBuildpack(domain.NewGUID("buildpack"))

	buildpack.Name = document.Name

	response, err := json.Marshal(buildpack)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

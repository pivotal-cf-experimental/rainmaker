package applications

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type createHandler struct {
	applications *domain.Applications
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var document documents.CreateApplicationRequest

	err := json.NewDecoder(req.Body).Decode(&document)
	if err != nil {
		panic(err)
	}

	application := domain.NewApplication(domain.NewGUID("app"))

	application.Name = document.Name
	application.SpaceGUID = document.SpaceGUID
	application.Diego = document.Diego

	h.applications.Add(application)

	response, err := json.Marshal(application)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

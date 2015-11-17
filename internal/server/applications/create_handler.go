package applications

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type createHandler struct {
	generateGUID guidGenerator
	applications *domain.Applications
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var document documents.CreateApplicationRequest

	err := json.NewDecoder(req.Body).Decode(&document)
	if err != nil {
		panic(err)
	}

	application := domain.NewApplication(h.generateGUID("app"))
	application.Name = document.Name
	application.SpaceGUID = document.SpaceGUID
	application.Diego = document.Diego

	h.applications.Add(application)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(application)
}

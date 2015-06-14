package fakes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

func (fake *CloudController) createSpace(w http.ResponseWriter, req *http.Request) {
	var document documents.CreateSpaceRequest
	err := json.NewDecoder(req.Body).Decode(&document)
	if err != nil {
		panic(err)
	}
	now := time.Now().UTC()

	space := domain.NewSpace(domain.NewGUID("space"))
	space.Name = document.Name
	space.OrganizationGUID = document.OrganizationGUID
	space.CreatedAt = now
	space.UpdatedAt = now

	fake.Spaces.Add(space)

	response, err := json.Marshal(space)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

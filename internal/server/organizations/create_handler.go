package organizations

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type createHandler struct {
	organizations *domain.Organizations
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var document documents.CreateOrganizationRequest
	now := time.Now().UTC()
	err := json.NewDecoder(req.Body).Decode(&document)
	if err != nil {
		panic(err)
	}

	organization := domain.NewOrganization(domain.NewGUID("org"))

	if document.GUID != "" {
		organization.GUID = document.GUID
	}

	organization.Name = document.Name
	organization.CreatedAt = now
	organization.UpdatedAt = now

	h.organizations.Add(organization)

	response, err := json.Marshal(organization)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

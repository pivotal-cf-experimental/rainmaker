package organizations

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type updateHandler struct {
	organizations *domain.Organizations
}

func (h updateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	organizationGUID := strings.TrimPrefix(req.URL.Path, "/v2/organizations/")
	decoder := json.NewDecoder(req.Body)
	updateRequest := documents.UpdateOrganizationRequest{}

	err := decoder.Decode(&updateRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	organization, ok := h.organizations.Get(organizationGUID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	organization.Name = updateRequest.Name
	organization.Status = updateRequest.Status
	organization.QuotaDefinitionGUID = updateRequest.QuotaDefinitionGUID
	h.organizations.Add(organization)

	response, err := json.Marshal(organization)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

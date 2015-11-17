package spaces

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type createHandler struct {
	orgs   *domain.Organizations
	spaces *domain.Spaces
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var document documents.CreateSpaceRequest
	now := time.Now().UTC()
	err := json.NewDecoder(req.Body).Decode(&document)
	if err != nil {
		panic(err)
	}

	org, ok := h.orgs.Get(document.OrganizationGUID)
	if !ok {
		common.NotFound(w)
		return
	}

	space := domain.NewSpace(domain.NewGUID("space"))

	if document.GUID != "" {
		space.GUID = document.GUID
	}

	space.Name = document.Name
	space.OrganizationGUID = document.OrganizationGUID
	space.CreatedAt = now
	space.UpdatedAt = now

	h.spaces.Add(space)

	org.Spaces.Add(space)

	response, err := json.Marshal(space)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

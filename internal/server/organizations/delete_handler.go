package organizations

import (
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type deleteHandler struct {
	organizations *domain.Organizations
}

func (h deleteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	organizationGUID := strings.TrimPrefix(req.URL.Path, "/v2/organizations/")

	h.organizations.Delete(organizationGUID)

	w.WriteHeader(http.StatusNoContent)
}

package organizations

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type associateAuditorHandler struct {
	organizations *domain.Organizations
	users         *domain.Users
}

func (h associateAuditorHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/organizations/(.*)/auditors/(.*)$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	org, ok := h.organizations.Get(matches[1])
	if !ok {
		common.NotFound(w)
		return
	}

	auditor, ok := h.users.Get(matches[2])
	if !ok {
		common.NotFound(w)
		return
	}

	org.Auditors.Add(auditor)

	response, err := json.Marshal(org)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

package organizations

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type associateManagerHandler struct {
	organizations *domain.Organizations
	users         *domain.Users
}

func (h associateManagerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/organizations/(.*)/managers/(.*)$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	org, ok := h.organizations.Get(matches[1])
	if !ok {
		common.NotFound(w)
		return
	}

	manager, ok := h.users.Get(matches[2])
	if !ok {
		common.NotFound(w)
		return
	}

	org.Managers.Add(manager)

	response, err := json.Marshal(org)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

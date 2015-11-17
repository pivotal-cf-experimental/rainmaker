package spaces

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type associateDeveloperHandler struct {
	spaces *domain.Spaces
	users  *domain.Users
}

func (h associateDeveloperHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/spaces/(.*)/developers/(.*)$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	space, ok := h.spaces.Get(matches[1])
	if !ok {
		common.NotFound(w)
		return
	}

	developer, ok := h.users.Get(matches[2])
	if !ok {
		common.NotFound(w)
		return
	}

	space.Developers.Add(developer)

	response, err := json.Marshal(space)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

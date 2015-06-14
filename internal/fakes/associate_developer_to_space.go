package fakes

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
)

func (fake *CloudController) associateDeveloperToSpace(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/spaces/(.*)/developers/(.*)$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	space, ok := fake.Spaces.Get(matches[1])
	if !ok {
		common.NotFound(w)
		return
	}

	developer, ok := fake.Users.Get(matches[2])
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

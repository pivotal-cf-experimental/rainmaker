package fakes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

func (fake *CloudController) getUsers(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	pageNum := common.ParseInt(query.Get("page"), 1)
	perPage := common.ParseInt(query.Get("results-per-page"), 10)

	page := domain.NewPage(fake.filteredUsers(query.Get("q")), req.URL.Path, pageNum, perPage)
	response, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (fake *CloudController) filteredUsers(query string) *domain.Users {
	switch {
	case strings.Contains(query, "space_guid:"):
		spaceGUID := strings.TrimPrefix(query, "space_guid:")
		space, ok := fake.Spaces.Get(spaceGUID)
		if !ok {
			return domain.NewUsers()
		}

		return space.Developers
	default:
		return fake.Users
	}
}

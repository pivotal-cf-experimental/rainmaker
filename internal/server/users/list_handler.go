package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type listHandler struct {
	users  *domain.Users
	spaces *domain.Spaces
}

func (h listHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	pageNum := common.ParseInt(query.Get("page"), 1)
	perPage := common.ParseInt(query.Get("results-per-page"), 10)

	page := domain.NewPage(h.filteredUsers(query.Get("q")), req.URL.Path, pageNum, perPage)
	response, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h listHandler) filteredUsers(query string) *domain.Users {
	switch {
	case strings.Contains(query, "space_guid:"):
		spaceGUID := strings.TrimPrefix(query, "space_guid:")
		space, ok := h.spaces.Get(spaceGUID)
		if !ok {
			return domain.NewUsers()
		}

		return space.Developers
	default:
		return h.users
	}
}

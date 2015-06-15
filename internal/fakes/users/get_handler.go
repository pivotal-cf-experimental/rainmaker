package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type getHandler struct {
	users *domain.Users
}

func (h getHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	userGUID := strings.TrimPrefix(req.URL.Path, "/v2/users/")

	user, ok := h.users.Get(userGUID)
	if !ok {
		common.NotFound(w)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

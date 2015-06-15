package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type createHandler struct {
	users *domain.Users
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var document documents.CreateUserRequest
	now := time.Now().UTC()
	err := json.NewDecoder(req.Body).Decode(&document)
	if err != nil {
		panic(err)
	}

	user := domain.NewUser(domain.NewGUID("user"))
	user.GUID = document.GUID
	user.DefaultSpaceGUID = document.DefaultSpaceGUID
	user.CreatedAt = now
	user.UpdatedAt = now

	h.users.Add(user)

	response, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

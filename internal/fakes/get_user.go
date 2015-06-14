package fakes

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (fake *CloudController) getUser(w http.ResponseWriter, req *http.Request) {
	userGUID := strings.TrimPrefix(req.URL.Path, "/v2/users/")

	user, ok := fake.Users.Get(userGUID)
	if !ok {
		fake.notFound(w)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

package fakes

import (
	"net/http"
	"strings"
)

func (fake *CloudController) GetSpace(w http.ResponseWriter, req *http.Request) {
	spaceGUID := strings.TrimPrefix(req.URL.Path, "/v2/spaces/")

	space, ok := fake.Spaces.Get(spaceGUID)
	if !ok {
		fake.NotFound(w)
		return
	}

	response, err := space.MarshalJSON()
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

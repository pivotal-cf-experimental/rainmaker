package fakes

import (
	"net/http"
	"strings"
)

func (fake *CloudController) GetOrganization(w http.ResponseWriter, req *http.Request) {
	organizationGUID := strings.TrimPrefix(req.URL.Path, "/v2/organizations/")

	organization, ok := fake.Organizations.Get(organizationGUID)
	if !ok {
		fake.NotFound(w)
		return
	}

	response, err := organization.ToJSON()
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

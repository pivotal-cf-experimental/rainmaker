package fakes

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func (fake *CloudController) associateAuditorToOrganization(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/organizations/(.*)/auditors/(.*)$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	org, ok := fake.Organizations.Get(matches[1])
	if !ok {
		fake.notFound(w)
		return
	}

	auditor, ok := fake.Users.Get(matches[2])
	if !ok {
		fake.notFound(w)
		return
	}

	org.Auditors.Add(auditor)

	response, err := json.Marshal(org)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

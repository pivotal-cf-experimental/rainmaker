package fakes

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

func (fake *CloudController) getOrganizationAuditors(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/organizations/(.*)/auditors$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	query := req.URL.Query()
	pageNum := parseInt(query.Get("page"), 1)
	perPage := parseInt(query.Get("results-per-page"), 10)

	org, ok := fake.Organizations.Get(matches[1])
	if !ok {
		fake.notFound(w)
		return
	}

	page := domain.NewPage(org.Auditors, req.URL.Path, pageNum, perPage)
	response, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

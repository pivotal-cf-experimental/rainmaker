package organizations

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type getSpacesHandler struct {
	organizations *domain.Organizations
}

func (h getSpacesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/organizations/(.*)/spaces$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	query := req.URL.Query()
	pageNum := common.ParseInt(query.Get("page"), 1)
	perPage := common.ParseInt(query.Get("results-per-page"), 10)

	org, ok := h.organizations.Get(matches[1])
	if !ok {
		common.NotFound(w)
		return
	}

	page := domain.NewPage(org.Spaces, req.URL.Path, pageNum, perPage)
	response, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

package spaces

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type getDevelopersHandler struct {
	spaces *domain.Spaces
}

func (h getDevelopersHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`^/v2/spaces/(.*)/developers$`)
	matches := r.FindStringSubmatch(req.URL.Path)

	query := req.URL.Query()
	pageNum := common.ParseInt(query.Get("page"), 1)
	perPage := common.ParseInt(query.Get("results-per-page"), 10)

	space, ok := h.spaces.Get(matches[1])
	if !ok {
		common.NotFound(w)
		return
	}

	page := domain.NewPage(space.Developers, req.URL.Path, pageNum, perPage)
	response, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

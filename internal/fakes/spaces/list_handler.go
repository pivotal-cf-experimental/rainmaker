package spaces

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type listHandler struct {
	spaces *domain.Spaces
}

func (h listHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	pageNum := common.ParseInt(query.Get("page"), 1)
	perPage := common.ParseInt(query.Get("results-per-page"), 10)

	page := domain.NewPage(h.spaces, req.URL.Path, pageNum, perPage)
	response, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

package applications

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type summaryHandler struct {
	applications *domain.Applications
}

func (h summaryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	application, ok := h.applications.Get(vars["guid"])
	if !ok {
		common.NotFound(w)
		return
	}

	response, err := json.Marshal(application.GetSummary())
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

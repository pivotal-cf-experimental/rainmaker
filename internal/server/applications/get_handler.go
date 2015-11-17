package applications

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type getHandler struct {
	applications *domain.Applications
}

func (h getHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	app, ok := h.applications.Get(vars["guid"])
	if !ok {
		common.NotFound(w)
		return
	}

	response, err := json.Marshal(app)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

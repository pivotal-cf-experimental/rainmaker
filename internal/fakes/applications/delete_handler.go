package applications

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type deleteHandler struct {
	applications *domain.Applications
}

func (h deleteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	h.applications.Delete(vars["guid"])

	w.WriteHeader(http.StatusNoContent)
}

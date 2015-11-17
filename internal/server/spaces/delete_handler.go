package spaces

import (
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type deleteHandler struct {
	spaces *domain.Spaces
}

func (h deleteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	spaceGUID := strings.TrimPrefix(req.URL.Path, "/v2/spaces/")

	h.spaces.Delete(spaceGUID)

	w.WriteHeader(http.StatusNoContent)
}

package spaces

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
)

type getHandler struct {
	spaces *domain.Spaces
}

func (h getHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	spaceGUID := strings.TrimPrefix(req.URL.Path, "/v2/spaces/")

	space, ok := h.spaces.Get(spaceGUID)
	if !ok {
		common.NotFound(w)
		return
	}

	response, err := json.Marshal(space)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

package service_instances

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/common"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type getHandler struct {
	serviceInstances *domain.ServiceInstances
}

func (h getHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	instanceGUID := strings.TrimPrefix(req.URL.Path, "/v2/service_instances/")

	instance, ok := h.serviceInstances.Get(instanceGUID)
	if !ok {
		common.NotFound(w)
	}

	response, err := json.Marshal(instance)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

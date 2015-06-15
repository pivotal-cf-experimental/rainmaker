package service_instances

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

type createHandler struct {
	serviceInstances *domain.ServiceInstances
}

func (h createHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var document documents.CreateServiceInstanceRequest
	err := json.NewDecoder(req.Body).Decode(&document)
	if err != nil {
		panic(err)
	}

	now := time.Now().UTC()
	instance := domain.NewServiceInstance(domain.NewGUID("service-instance"))
	instance.Name = document.Name
	instance.PlanGUID = document.PlanGUID
	instance.SpaceGUID = document.SpaceGUID
	instance.CreatedAt = now
	instance.UpdatedAt = now

	h.serviceInstances.Add(instance)

	response, err := json.Marshal(instance)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

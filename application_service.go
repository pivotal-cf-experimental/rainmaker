package rainmaker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/network"
)

type ApplicationsService struct {
	config Config
}

func NewApplicationsService(config Config) ApplicationsService {
	return ApplicationsService{
		config: config,
	}
}

func (service ApplicationsService) Create(application Application, token string) (Application, error) {
	resp, err := newNetworkClient(service.config).MakeRequest(network.Request{
		Method: "POST",
		Path:   "/v2/apps",
		Body: network.NewJSONRequestBody(documents.CreateApplicationRequest{
			Name:      application.Name,
			SpaceGUID: application.SpaceGUID,
		}),
		Authorization:         network.NewTokenAuthorization(token),
		AcceptableStatusCodes: []int{http.StatusCreated},
	})
	if err != nil {
		return Application{}, translateError(err)
	}

	var response documents.ApplicationCreateResponse
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		panic(err)
	}

	return newApplicationFromCreateResponse(service.config, response), nil
}

func (service ApplicationsService) Summary(guid, token string) (Application, error) {
	resp, err := newNetworkClient(service.config).MakeRequest(network.Request{
		Method:                "GET",
		Path:                  fmt.Sprintf("/v2/apps/%s/summary", guid),
		Body:                  network.NewJSONRequestBody(documents.CreateApplicationRequest{}),
		Authorization:         network.NewTokenAuthorization(token),
		AcceptableStatusCodes: []int{http.StatusOK},
	})
	if err != nil {
		return Application{}, translateError(err)
	}

	var response documents.ApplicationSummaryResponse
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		return Application{}, translateError(err)
	}

	return newApplicationFromSummaryResponse(service.config, response), nil
}

func (service ApplicationsService) Delete(guid, token string) error {
	_, err := newNetworkClient(service.config).MakeRequest(network.Request{
		Method:                "DELETE",
		Path:                  fmt.Sprintf("/v2/apps/%s", guid),
		Body:                  network.NewJSONRequestBody(documents.CreateApplicationRequest{}),
		Authorization:         network.NewTokenAuthorization(token),
		AcceptableStatusCodes: []int{http.StatusNoContent},
	})
	if err != nil {
		return translateError(err)
	}

	return nil
}

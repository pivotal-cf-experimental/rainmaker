package rainmaker

import (
	"encoding/json"

	"github.com/pivotal-golang/rainmaker/internal/documents"
)

type OrganizationsService struct {
	client Client
}

func NewOrganizationsService(client Client) *OrganizationsService {
	return &OrganizationsService{
		client: client,
	}
}

func (service OrganizationsService) Get(guid string) Organization {
	_, body, err := service.client.makeRequest("GET", "/v2/organizations/"+guid, nil)
	if err != nil {
		panic(err)
	}

	var response documents.OrganizationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return NewOrganizationFromResponse(response)
}

func (service OrganizationsService) ListUsers(guid string) UsersList {
	_, body, err := service.client.makeRequest("GET", "/v2/organizations/"+guid+"/users", nil)
	if err != nil {
		panic(err)
	}

	var response documents.UsersListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return NewUsersListFromResponse(response)
}

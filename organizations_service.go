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

func (service OrganizationsService) Get(guid, token string) (Organization, error) {
	_, body, err := service.client.makeRequest(requestArguments{
		Method: "GET",
		Path:   "/v2/organizations/" + guid,
		Token:  token,
	})
	if err != nil {
		return Organization{}, err
	}

	var response documents.OrganizationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Organization{}, err
	}

	return NewOrganizationFromResponse(response), nil
}

func (service OrganizationsService) ListUsers(guid, token string) (UsersList, error) {
	_, body, err := service.client.makeRequest(requestArguments{
		Method: "GET",
		Path:   "/v2/organizations/" + guid + "/users",
		Token:  token,
	})
	if err != nil {
		return UsersList{}, err
	}

	var response documents.UsersListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return UsersList{}, err
	}

	return NewUsersListFromResponse(response), nil
}

func (service OrganizationsService) ListBillingManagers(guid, token string) (UsersList, error) {
	_, body, err := service.client.makeRequest(requestArguments{
		Method: "GET",
		Path:   "/v2/organizations/" + guid + "/billing_managers",
		Token:  token,
	})
	if err != nil {
		return UsersList{}, err
	}

	var response documents.UsersListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return UsersList{}, err
	}

	return NewUsersListFromResponse(response), nil
}

func (service OrganizationsService) ListAuditors(guid, token string) (UsersList, error) {
	_, body, err := service.client.makeRequest(requestArguments{
		Method: "GET",
		Path:   "/v2/organizations/" + guid + "/auditors",
		Token:  token,
	})
	if err != nil {
		return UsersList{}, err
	}

	var response documents.UsersListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return UsersList{}, err
	}

	return NewUsersListFromResponse(response), nil
}

func (service OrganizationsService) ListManagers(guid, token string) (UsersList, error) {
	_, body, err := service.client.makeRequest(requestArguments{
		Method: "GET",
		Path:   "/v2/organizations/" + guid + "/managers",
		Token:  token,
	})
	if err != nil {
		return UsersList{}, err
	}

	var response documents.UsersListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return UsersList{}, err
	}

	return NewUsersListFromResponse(response), nil
}

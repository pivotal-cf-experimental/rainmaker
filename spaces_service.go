package rainmaker

import (
	"encoding/json"

	"github.com/pivotal-golang/rainmaker/internal/documents"
)

type SpacesService struct {
	client Client
}

func NewSpacesService(client Client) *SpacesService {
	return &SpacesService{
		client: client,
	}
}

func (service SpacesService) Get(guid, token string) (Space, error) {
	_, body, err := service.client.makeRequest(requestArguments{
		Method: "GET",
		Path:   "/v2/spaces/" + guid,
		Token:  token,
	})
	if err != nil {
		return Space{}, err
	}

	var response documents.SpaceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Space{}, err
	}

	return NewSpaceFromResponse(response), nil
}

func (service SpacesService) ListUsers(guid, token string) (UsersList, error) {
	_, body, err := service.client.makeRequest(requestArguments{
		Method: "GET",
		Path:   "/v2/users?q=space_guid:" + guid,
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

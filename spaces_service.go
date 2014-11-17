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

func (service SpacesService) Get(guid string) Space {
	_, body, err := service.client.makeRequest("GET", "/v2/spaces/"+guid, nil)
	if err != nil {
		panic(err)
	}

	var response documents.SpaceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return NewSpaceFromResponse(response)
}

func (service SpacesService) ListUsers(guid string) UsersList {
	_, body, err := service.client.makeRequest("GET", "/v2/users?q=space_guid:"+guid, nil)
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

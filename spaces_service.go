package rainmaker

import (
	"fmt"
	"net/url"
)

type SpacesService struct {
	config Config
}

func NewSpacesService(config Config) *SpacesService {
	return &SpacesService{
		config: config,
	}
}

func (service SpacesService) Get(guid, token string) (Space, error) {
	return FetchSpace(service.config, "/v2/spaces/"+guid, token)
}

func (service SpacesService) ListUsers(guid, token string) (UsersList, error) {
	query := url.Values{}
	query.Set("q", fmt.Sprintf("space_guid:%s", guid))

	return FetchUsersList(service.config, NewRequestPlan("/v2/users", query), token)
}

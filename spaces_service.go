package rainmaker

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
	return FetchUsersList(service.config, "/v2/users?q=space_guid:"+guid, token)
}

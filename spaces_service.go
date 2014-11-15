package rainmaker

type SpacesService struct {
	client Client
}

func NewSpacesService(client Client) *SpacesService {
	return &SpacesService{
		client: client,
	}
}

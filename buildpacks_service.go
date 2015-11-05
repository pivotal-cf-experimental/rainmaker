package rainmaker

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/network"
)

type BuildpacksService struct {
	config Config
}

func (b BuildpacksService) Create(name string, token string) (Buildpack, error) {
	resp, err := newNetworkClient(b.config).MakeRequest(network.Request{
		Method: "POST",
		Path:   "/v2/buildpacks",
		Body: network.NewJSONRequestBody(documents.CreateBuildpackRequest{
			Name: name,
		}),
		Authorization:         network.NewTokenAuthorization(token),
		AcceptableStatusCodes: []int{http.StatusCreated},
	})
	if err != nil {
		panic(err)
	}

	var response documents.BuildpackResponse
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		panic(err)
	}

	return newBuildpackFromResponse(b.config, response), nil
}

func NewBuildpacksService(config Config) BuildpacksService {
	return BuildpacksService{
		config: config,
	}
}

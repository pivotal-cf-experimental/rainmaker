package rainmaker

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pivotal-golang/rainmaker/internal/network"
)

type Client struct {
	Config        Config
	Organizations *OrganizationsService
	Spaces        *SpacesService
}

func NewClient(config Config) Client {
	client := Client{
		Config: config,
	}

	client.Organizations = NewOrganizationsService(client)
	client.Spaces = NewSpacesService(client)

	return client
}

func (client Client) makeRequest(method, path string, body interface{}) (int, []byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return 0, []byte{}, err
	}

	request, err := http.NewRequest(method, client.Config.Host+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, []byte{}, err
	}

	networkClient := network.GetClient(network.Config{SkipVerifySSL: client.Config.SkipVerifySSL})
	response, err := networkClient.Do(request)
	if err != nil {
		return 0, []byte{}, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, []byte{}, err
	}

	return response.StatusCode, responseBody, nil
}

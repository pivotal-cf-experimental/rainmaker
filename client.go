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

func (client Client) makeRequest(requestArgs requestArguments) (int, []byte, error) {
	jsonBody, err := json.Marshal(requestArgs.Body)
	if err != nil {
		return 0, []byte{}, NewRequestBodyMarshalError(err)
	}

	request, err := http.NewRequest(requestArgs.Method, client.Config.Host+requestArgs.Path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, []byte{}, NewRequestConfigurationError(err)
	}

	request.Header.Set("Authorization", "Bearer "+requestArgs.Token)

	networkClient := network.GetClient(network.Config{SkipVerifySSL: client.Config.SkipVerifySSL})
	response, err := networkClient.Do(request)
	if err != nil {
		return 0, []byte{}, NewRequestHTTPError(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, []byte{}, NewResponseReadError(err)
	}

	return response.StatusCode, responseBody, nil
}

func (client Client) unmarshal(body []byte, response interface{}) error {
	err := json.Unmarshal(body, response)
	if err != nil {
		return NewResponseBodyUnmarshalError(err)
	}
	return nil
}

type requestArguments struct {
	Method string
	Path   string
	Token  string
	Body   interface{}
}

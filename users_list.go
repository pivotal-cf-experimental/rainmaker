package rainmaker

import (
	"encoding/json"

	"github.com/pivotal-golang/rainmaker/internal/documents"
)

type UsersList struct {
	TotalResults int
	TotalPages   int
	NextURL      string
	client       Client
	Users        []User
}

func NewUsersList(client Client) UsersList {
	return UsersList{
		client: client,
	}
}

func (list UsersList) Next(token string) (UsersList, error) {
	return FetchUsersList(list.client, list.NextURL, token)
}

func FetchUsersList(client Client, path, token string) (UsersList, error) {
	_, body, err := client.makeRequest(requestArguments{
		Method: "GET",
		Path:   path,
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

func NewUsersListFromResponse(response documents.UsersListResponse) UsersList {
	list := UsersList{
		TotalResults: response.TotalResults,
		TotalPages:   response.TotalPages,
		Users:        make([]User, 0),
	}

	for _, userResponse := range response.Resources {
		list.Users = append(list.Users, NewUserFromResponse(userResponse))
	}

	return list
}

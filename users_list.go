package rainmaker

import (
	"encoding/json"

	"github.com/pivotal-golang/rainmaker/internal/documents"
)

type UsersList struct {
	config       Config
	TotalResults int
	TotalPages   int
	NextURL      string
	PrevURL      string
	Users        []User
}

func NewUsersList(config Config) UsersList {
	return UsersList{
		config: config,
	}
}

func NewUsersListFromResponse(config Config, response documents.UsersListResponse) UsersList {
	list := NewUsersList(config)
	list.TotalResults = response.TotalResults
	list.TotalPages = response.TotalPages
	list.PrevURL = response.PrevURL
	list.NextURL = response.NextURL
	list.Users = make([]User, 0)

	for _, userResponse := range response.Resources {
		list.Users = append(list.Users, NewUserFromResponse(userResponse))
	}

	return list
}

func FetchUsersList(config Config, path, token string) (UsersList, error) {
	_, body, err := NewClient(config).makeRequest(requestArguments{
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

	return NewUsersListFromResponse(config, response), nil
}

func (list UsersList) HasNextPage() bool {
	return list.NextURL != ""
}

func (list UsersList) HasPrevPage() bool {
	return list.PrevURL != ""
}

func (list UsersList) Next(token string) (UsersList, error) {
	return FetchUsersList(list.config, list.NextURL, token)
}

func (list UsersList) Prev(token string) (UsersList, error) {
	return FetchUsersList(list.config, list.PrevURL, token)
}

func (list UsersList) AllUsers(token string) ([]User, error) {
	l := list
	users := make([]User, 0)

	for l.HasPrevPage() {
		var err error
		l, err = list.Prev(token)
		if err != nil {
			return []User{}, err
		}

		users = append(l.Users, users...)
	}

	users = append(users, list.Users...)

	for l.HasNextPage() {
		var err error
		l, err = list.Next(token)
		if err != nil {
			return []User{}, err
		}

		users = append(users, l.Users...)
	}

	return users, nil
}

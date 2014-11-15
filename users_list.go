package rainmaker

import "github.com/pivotal-golang/rainmaker/internal/documents"

type UsersList struct {
	TotalResults int
	TotalPages   int
	Users        []User
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

package fakes

import (
	"encoding/json"
	"net/http"
)

func (fake *CloudController) GetOrganizationUsers(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	pageNumber := req.Form.Get("page")

	var page Page
	var nextURL, prevURL string

	switch pageNumber {
	case "", "1":
		page = Page{
			Number: 1,
			GUIDs:  []string{"user-123", "user-456"},
		}
		nextURL = req.URL.Path + "?page=2"
	case "2":
		page = Page{
			Number: 2,
			GUIDs:  []string{"user-next"},
		}
		prevURL = req.URL.Path + "?page=1"
		nextURL = req.URL.Path + "?page=3"
	case "3":
		page = Page{
			Number: 3,
			GUIDs:  []string{"user-last"},
		}
		prevURL = req.URL.Path + "?page=2"
	}

	document := map[string]interface{}{
		"total_results": 4,
		"total_pages":   3,
		"prev_url":      prevURL,
		"next_url":      nextURL,
		"resources":     make([]map[string]interface{}, 0),
	}

	for _, userGUID := range page.GUIDs {
		document["resources"] = append(document["resources"].([]map[string]interface{}), map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid":       userGUID,
				"url":        "/v2/users/" + userGUID,
				"created_at": "2014-11-11T18:22:51+00:00",
				"updated_at": nil,
			},
			"entity": map[string]interface{}{
				"admin":                             false,
				"active":                            true,
				"default_space_guid":                nil,
				"spaces_url":                        "/v2/users/" + userGUID + "/spaces",
				"organizations_url":                 "/v2/users/" + userGUID + "/organizations",
				"managed_organizations_url":         "/v2/users/" + userGUID + "/managed_organizations",
				"billing_managed_organizations_url": "/v2/users/" + userGUID + "/billing_managed_organizations",
				"audited_organizations_url":         "/v2/users/" + userGUID + "/audited_organizations",
				"managed_spaces_url":                "/v2/users/" + userGUID + "/managed_spaces",
				"audited_spaces_url":                "/v2/users/" + userGUID + "/audited_spaces",
			},
		})
	}

	response, err := json.Marshal(document)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

type Page struct {
	Number int
	GUIDs  []string
}

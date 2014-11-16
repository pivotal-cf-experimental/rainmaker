package fakes

import (
	"encoding/json"
	"net/http"
)

func (fake *CloudController) GetOrganizationUsers(w http.ResponseWriter, req *http.Request) {
	userGUIDs := []string{"user-123", "user-456"}

	document := map[string]interface{}{
		"total_results": 2,
		"total_pages":   1,
		"prev_url":      nil,
		"next_url":      nil,
		"resources":     make([]map[string]interface{}, 0),
	}

	for _, userGUID := range userGUIDs {
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

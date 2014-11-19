package fakes

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (fake *CloudController) GetOrganization(w http.ResponseWriter, req *http.Request) {
	organizationGUID := "org-001"

	requestedOrganizationGUID := strings.TrimPrefix(req.URL.Path, "/v2/organizations/")
	if requestedOrganizationGUID != organizationGUID {
		errorBody, err := json.Marshal(map[string]interface{}{
			"code":        10000,
			"description": "Unknown request",
			"error_code":  "CF-NotFound",
		})
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(errorBody)
		return
	}

	organizationName := "rainmaker-organization"
	quotaDefinitionGUID := "quota-definition-guid"
	billingEnabled := false
	organizationStatus := "active"

	response, err := json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid":       organizationGUID,
			"url":        "/v2/organizations/" + organizationGUID,
			"created_at": "2014-11-11T18:34:16+00:00",
			"updated_at": nil,
		},
		"entity": map[string]interface{}{
			"name":                        organizationName,
			"billing_enabled":             billingEnabled,
			"quota_definition_guid":       quotaDefinitionGUID,
			"status":                      organizationStatus,
			"quota_definition_url":        "/v2/quota_definitions/" + quotaDefinitionGUID,
			"spaces_url":                  "/v2/organizations/" + organizationGUID + "/spaces",
			"domains_url":                 "/v2/organizations/" + organizationGUID + "/domains",
			"private_domains_url":         "/v2/organizations/" + organizationGUID + "/private_domains",
			"users_url":                   "/v2/organizations/" + organizationGUID + "/users",
			"managers_url":                "/v2/organizations/" + organizationGUID + "/managers",
			"billing_managers_url":        "/v2/organizations/" + organizationGUID + "/billing_managers",
			"auditors_url":                "/v2/organizations/" + organizationGUID + "/auditors",
			"app_events_url":              "/v2/organizations/" + organizationGUID + "/app_events",
			"space_quota_definitions_url": "/v2/organizations/" + organizationGUID + "/space_quota_definitions",
		},
	})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

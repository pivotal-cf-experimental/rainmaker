package fakes

import (
	"encoding/json"
	"net/http"
)

func (fake *CloudController) GetSpace(w http.ResponseWriter, req *http.Request) {
	spaceGUID := "space-001"
	spaceName := "development"
	organizationGUID := "org-001"

	response, err := json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid":       spaceGUID,
			"url":        "/v2/spaces/" + spaceGUID,
			"created_at": "2014-10-09T22:02:26+00:00",
			"updated_at": nil,
		},
		"entity": map[string]interface{}{
			"name":                        spaceName,
			"organization_guid":           organizationGUID,
			"space_quota_definition_guid": nil,
			"organization_url":            "/v2/organizations/" + organizationGUID,
			"developers_url":              "/v2/spaces/" + spaceGUID + "/developers",
			"managers_url":                "/v2/spaces/" + spaceGUID + "/managers",
			"auditors_url":                "/v2/spaces/" + spaceGUID + "/auditors",
			"apps_url":                    "/v2/spaces/" + spaceGUID + "/apps",
			"routes_url":                  "/v2/spaces/" + spaceGUID + "/routes",
			"domains_url":                 "/v2/spaces/" + spaceGUID + "/domains",
			"service_instances_url":       "/v2/spaces/" + spaceGUID + "/service_instances",
			"app_events_url":              "/v2/spaces/" + spaceGUID + "/app_events",
			"events_url":                  "/v2/spaces/" + spaceGUID + "/events",
			"security_groups_url":         "/v2/spaces/" + spaceGUID + "/security_groups",
		},
	})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

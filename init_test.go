package rainmaker_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var fakeCloudController *FakeCloudController

func TestRainmakerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rainmaker Suite")
}

var _ = BeforeSuite(func() {
	fakeCloudController = NewFakeCloudController()
	fakeCloudController.Start()
})

var _ = AfterSuite(func() {
	fakeCloudController.Close()
})

type FakeCloudController struct {
	server *httptest.Server
}

func NewFakeCloudController() *FakeCloudController {
	fake := &FakeCloudController{}
	router := mux.NewRouter()
	router.HandleFunc("/v2/organizations/{guid}", fake.GetOrganization)
	router.HandleFunc("/v2/organizations/{guid}/users", fake.GetOrganizationUsers)

	fake.server = httptest.NewUnstartedServer(router)
	return fake
}

func (fake *FakeCloudController) Start() {
	fake.server.Start()
}

func (fake *FakeCloudController) Close() {
	fake.server.Close()
}

func (fake *FakeCloudController) URL() string {
	return fake.server.URL
}

func (fake *FakeCloudController) GetOrganization(w http.ResponseWriter, req *http.Request) {
	organizationGUID := "org-001"
	organizationName := "rainmaker-organization"
	quotaDefinitionGUID := "quota-definition-guid"
	billingEnabled := false
	organizationStatus := "active"

	response, err := json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid":       organizationGUID,
			"url":        "/v2/organizations/" + organizationGUID,
			"created_at": "2014-11-11T18:34:16+00:00",
			"updated_at": "2014-11-14T20:02:09+00:00",
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

func (fake *FakeCloudController) GetOrganizationUsers(w http.ResponseWriter, req *http.Request) {
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

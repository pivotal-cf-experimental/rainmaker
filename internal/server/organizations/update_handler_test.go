package organizations_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/organizations"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("updateHandler", func() {
	var (
		orgs     *domain.Organizations
		users    *domain.Users
		router   http.Handler
		recorder *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		orgs = domain.NewOrganizations()
		users = domain.NewUsers()
		router = organizations.NewRouter(orgs, users)
		recorder = httptest.NewRecorder()

	})

	It("returns a 404 when the organization does not exist", func() {
		request, err := http.NewRequest("PUT", "/v2/organizations/does-not-exist", strings.NewReader("{}"))
		Expect(err).NotTo(HaveOccurred())
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusNotFound))
	})

	It("returns an error when passed malformed JSON", func() {
		malformedJSON := "{iaskjdh}"
		request, err := http.NewRequest("PUT", "/v2/organizations/some-guid", strings.NewReader(malformedJSON))
		Expect(err).NotTo(HaveOccurred())

		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
	})

	It("updates the organization", func() {
		org := domain.NewOrganization("some-guid")
		orgs.Add(org)

		requestBody, err := json.Marshal(map[string]string{
			"name":                  "some-updated-name",
			"status":                "some-updated-status",
			"quota_definition_guid": "some-updated-quota-definition-guid",
		})
		Expect(err).NotTo(HaveOccurred())

		request, err := http.NewRequest("PUT", "/v2/organizations/some-guid", bytes.NewBuffer(requestBody))
		Expect(err).NotTo(HaveOccurred())

		router.ServeHTTP(recorder, request)

		Expect(recorder.Code).To(Equal(http.StatusCreated))
		Expect(recorder.Body).To(MatchJSON(`{
			"metadata": {
				"guid": "some-guid",
				"url": "/v2/organizations/some-guid",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			},
			"entity": {
				"name": "some-updated-name",
				"billing_enabled": false,
				"quota_definition_guid": "some-updated-quota-definition-guid",
				"status": "some-updated-status",
				"quota_definition_url": "/v2/quota_definitions/some-updated-quota-definition-guid",
				"spaces_url": "/v2/organizations/some-guid/spaces",
				"domains_url": "/v2/organizations/some-guid/domains",
				"private_domains_url": "/v2/organizations/some-guid/private_domains",
				"users_url": "/v2/organizations/some-guid/users",
				"managers_url": "/v2/organizations/some-guid/managers",
				"billing_managers_url": "/v2/organizations/some-guid/billing_managers",
				"auditors_url": "/v2/organizations/some-guid/auditors",
				"app_events_url": "/v2/organizations/some-guid/app_events",
				"space_quota_definitions_url": "/v2/organizations/some-guid/space_quota_definitions"
			}
		}`))

		organization, ok := orgs.Get("some-guid")
		Expect(ok).To(BeTrue())
		Expect(organization.Name).To(Equal("some-updated-name"))
		Expect(organization.Status).To(Equal("some-updated-status"))
		Expect(organization.QuotaDefinitionGUID).To(Equal("some-updated-quota-definition-guid"))
	})
})

package rainmaker_test

import (
	"time"

	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Organization", func() {
	var config rainmaker.Config

	Describe("FetchOrganization", func() {
		BeforeEach(func() {
			config = rainmaker.Config{
				Host: fakeCloudController.URL(),
			}
		})

		It("retrieves the organization", func() {
			organization, err := rainmaker.FetchOrganization(config, "/v2/organizations/org-001", "token-123")
			Expect(err).NotTo(HaveOccurred())

			createdAt, err := time.Parse(time.RFC3339, "2014-11-11T18:34:16+00:00")
			if err != nil {
				panic(err)
			}
			expectedOrganization := rainmaker.NewOrganization(config)
			expectedOrganization.GUID = "org-001"
			expectedOrganization.Name = "rainmaker-organization"
			expectedOrganization.URL = "/v2/organizations/org-001"
			expectedOrganization.BillingEnabled = false
			expectedOrganization.Status = "active"
			expectedOrganization.QuotaDefinitionGUID = "quota-definition-guid"
			expectedOrganization.QuotaDefinitionURL = "/v2/quota_definitions/quota-definition-guid"
			expectedOrganization.SpacesURL = "/v2/organizations/org-001/spaces"
			expectedOrganization.DomainsURL = "/v2/organizations/org-001/domains"
			expectedOrganization.PrivateDomainsURL = "/v2/organizations/org-001/private_domains"
			expectedOrganization.UsersURL = "/v2/organizations/org-001/users"
			expectedOrganization.ManagersURL = "/v2/organizations/org-001/managers"
			expectedOrganization.BillingManagersURL = "/v2/organizations/org-001/billing_managers"
			expectedOrganization.AuditorsURL = "/v2/organizations/org-001/auditors"
			expectedOrganization.AppEventsURL = "/v2/organizations/org-001/app_events"
			expectedOrganization.SpaceQuotaDefinitionsURL = "/v2/organizations/org-001/space_quota_definitions"
			expectedOrganization.CreatedAt = createdAt
			expectedOrganization.UpdatedAt = time.Time{}

			Expect(organization).To(Equal(expectedOrganization))
		})

		It("handles NotFound errors", func() {
			_, err := rainmaker.FetchOrganization(config, "/v2/organizations/something", "token")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
		})

		It("handles unauthorized use", func() {
			_, err := rainmaker.FetchOrganization(config, "/v2/organizations/org-001", "")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(rainmaker.UnauthorizedError{}))

		})
	})
})

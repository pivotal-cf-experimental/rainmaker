package rainmaker_test

import (
	"time"

	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrganizationsService", func() {
	var service *rainmaker.OrganizationsService

	BeforeEach(func() {
		client := rainmaker.NewClient(rainmaker.Config{
			Host: fakeCloudController.URL(),
		})
		service = client.Organizations
	})

	Describe("Get", func() {
		It("returns the organization matching the given GUID", func() {
			createdAt, err := time.Parse(time.RFC3339, "2014-11-11T18:34:16+00:00")
			if err != nil {
				panic(err)
			}

			organization := service.Get("org-001")

			Expect(organization).To(Equal(rainmaker.Organization{
				GUID:                     "org-001",
				Name:                     "rainmaker-organization",
				URL:                      "/v2/organizations/org-001",
				BillingEnabled:           false,
				Status:                   "active",
				QuotaDefinitionGUID:      "quota-definition-guid",
				QuotaDefinitionURL:       "/v2/quota_definitions/quota-definition-guid",
				SpacesURL:                "/v2/organizations/org-001/spaces",
				DomainsURL:               "/v2/organizations/org-001/domains",
				PrivateDomainsURL:        "/v2/organizations/org-001/private_domains",
				UsersURL:                 "/v2/organizations/org-001/users",
				ManagersURL:              "/v2/organizations/org-001/managers",
				BillingManagersURL:       "/v2/organizations/org-001/billing_managers",
				AuditorsURL:              "/v2/organizations/org-001/auditors",
				AppEventsURL:             "/v2/organizations/org-001/app_events",
				SpaceQuotaDefinitionsURL: "/v2/organizations/org-001/space_quota_definitions",
				CreatedAt:                createdAt,
				UpdatedAt:                time.Time{},
			}))
		})
	})

	Describe("ListUsers", func() {
		It("returns the users belonging to the organization", func() {
			usersList := service.ListUsers("org-001")
			Expect(usersList.TotalResults).To(Equal(2))
			Expect(usersList.TotalPages).To(Equal(1))
			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-11T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			users := usersList.Users
			Expect(len(users)).To(Equal(2))
			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-123",
				URL:                            "/v2/users/user-123",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-123/spaces",
				OrganizationsURL:               "/v2/users/user-123/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-123/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-123/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-123/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-123/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-123/audited_spaces",
			}))

			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-456",
				URL:                            "/v2/users/user-456",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-456/spaces",
				OrganizationsURL:               "/v2/users/user-456/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-456/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-456/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-456/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-456/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-456/audited_spaces",
			}))
		})
	})
})

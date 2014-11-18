package rainmaker_test

import (
	"time"

	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpacesService", func() {
	var service *rainmaker.SpacesService

	BeforeEach(func() {
		client := rainmaker.NewClient(rainmaker.Config{
			Host: fakeCloudController.URL(),
		})
		service = rainmaker.NewSpacesService(client)
	})

	Describe("Get", func() {
		It("returns the space matching the given GUID", func() {
			space, err := service.Get("space-001", "token-asd")
			Expect(err).NotTo(HaveOccurred())
			createdAt, err := time.Parse(time.RFC3339, "2014-10-09T22:02:26+00:00")
			if err != nil {
				panic(err)
			}

			Expect(space).To(Equal(rainmaker.Space{
				GUID:                     "space-001",
				URL:                      "/v2/spaces/space-001",
				CreatedAt:                createdAt,
				UpdatedAt:                time.Time{},
				Name:                     "development",
				OrganizationGUID:         "org-001",
				SpaceQuotaDefinitionGUID: "",
				OrganizationURL:          "/v2/organizations/org-001",
				DevelopersURL:            "/v2/spaces/space-001/developers",
				ManagersURL:              "/v2/spaces/space-001/managers",
				AuditorsURL:              "/v2/spaces/space-001/auditors",
				AppsURL:                  "/v2/spaces/space-001/apps",
				RoutesURL:                "/v2/spaces/space-001/routes",
				DomainsURL:               "/v2/spaces/space-001/domains",
				ServiceInstancesURL:      "/v2/spaces/space-001/service_instances",
				AppEventsURL:             "/v2/spaces/space-001/app_events",
				EventsURL:                "/v2/spaces/space-001/events",
				SecurityGroupsURL:        "/v2/spaces/space-001/security_groups",
			}))
		})
	})

	Describe("ListUsers", func() {
		It("returns the users belonging to the space", func() {
			usersList, err := service.ListUsers("space-001", "token-abc")
			Expect(err).NotTo(HaveOccurred())
			Expect(usersList.TotalResults).To(Equal(2))
			Expect(usersList.TotalPages).To(Equal(1))
			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-01T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			users := usersList.Users
			Expect(len(users)).To(Equal(2))
			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-abc",
				URL:                            "/v2/users/user-abc",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-abc/spaces",
				OrganizationsURL:               "/v2/users/user-abc/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-abc/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-abc/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-abc/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-abc/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-abc/audited_spaces",
			}))

			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-xyz",
				URL:                            "/v2/users/user-xyz",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-xyz/spaces",
				OrganizationsURL:               "/v2/users/user-xyz/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-xyz/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-xyz/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-xyz/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-xyz/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-xyz/audited_spaces",
			}))
		})
	})
})

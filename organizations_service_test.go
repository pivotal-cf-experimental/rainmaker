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

			organization, err := service.Get("org-001", "token-123")
			Expect(err).NotTo(HaveOccurred())

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
			usersList, err := service.ListUsers("org-001", "token-456")
			Expect(err).NotTo(HaveOccurred())
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

	Describe("ListBillingManagers", func() {
		It("returns the billing managers belonging to the organization", func() {
			usersList, err := service.ListBillingManagers("org-002", "token-987")
			Expect(err).NotTo(HaveOccurred())
			Expect(usersList.TotalResults).To(Equal(2))
			Expect(usersList.TotalPages).To(Equal(1))
			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-04T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			users := usersList.Users
			Expect(len(users)).To(Equal(2))
			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-987",
				URL:                            "/v2/users/user-987",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-987/spaces",
				OrganizationsURL:               "/v2/users/user-987/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-987/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-987/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-987/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-987/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-987/audited_spaces",
			}))

			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-654",
				URL:                            "/v2/users/user-654",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-654/spaces",
				OrganizationsURL:               "/v2/users/user-654/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-654/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-654/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-654/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-654/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-654/audited_spaces",
			}))
		})
	})

	Describe("ListAuditors", func() {
		It("returns the auditors belonging to the organization", func() {
			usersList, err := service.ListAuditors("org-003", "token-012")
			Expect(err).NotTo(HaveOccurred())
			Expect(usersList.TotalResults).To(Equal(2))
			Expect(usersList.TotalPages).To(Equal(1))
			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-05T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			users := usersList.Users
			Expect(len(users)).To(Equal(2))
			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-asd",
				URL:                            "/v2/users/user-asd",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-asd/spaces",
				OrganizationsURL:               "/v2/users/user-asd/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-asd/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-asd/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-asd/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-asd/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-asd/audited_spaces",
			}))

			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-jkl",
				URL:                            "/v2/users/user-jkl",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-jkl/spaces",
				OrganizationsURL:               "/v2/users/user-jkl/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-jkl/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-jkl/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-jkl/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-jkl/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-jkl/audited_spaces",
			}))
		})
	})

	Describe("ListManagers", func() {
		It("returns the managers belonging to the organization", func() {
			usersList, err := service.ListManagers("org-004", "token-345")
			Expect(err).NotTo(HaveOccurred())
			Expect(usersList.TotalResults).To(Equal(2))
			Expect(usersList.TotalPages).To(Equal(1))
			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-21T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			users := usersList.Users
			Expect(len(users)).To(Equal(2))
			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-aaa",
				URL:                            "/v2/users/user-aaa",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-aaa/spaces",
				OrganizationsURL:               "/v2/users/user-aaa/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-aaa/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-aaa/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-aaa/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-aaa/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-aaa/audited_spaces",
			}))

			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-bbb",
				URL:                            "/v2/users/user-bbb",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-bbb/spaces",
				OrganizationsURL:               "/v2/users/user-bbb/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-bbb/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-bbb/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-bbb/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-bbb/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-bbb/audited_spaces",
			}))
		})
	})
})

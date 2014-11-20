package rainmaker_test

import (
	"time"

	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UsersList", func() {
	var config rainmaker.Config
	var list rainmaker.UsersList

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		list = rainmaker.NewUsersList(config)
	})

	Describe("FetchUsersList", func() {
		It("returns a UsersList object", func() {
			var err error
			list, err = rainmaker.FetchUsersList(config, "/v2/organizations/org-001/users", "token")
			Expect(err).NotTo(HaveOccurred())

			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-11T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			Expect(list.TotalResults).To(Equal(4))
			Expect(list.TotalPages).To(Equal(3))
			Expect(list.PrevURL).To(Equal(""))
			Expect(list.NextURL).To(Equal("/v2/organizations/org-001/users?page=2"))

			users := list.Users
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

	Describe("Next", func() {
		It("returns the next UserList result for the paginated set", func() {
			list.NextURL = "/v2/organizations/org-001/users?page=2"
			nextList, err := list.Next("token")
			Expect(err).NotTo(HaveOccurred())

			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-11T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			Expect(nextList.TotalResults).To(Equal(4))
			Expect(nextList.TotalPages).To(Equal(3))
			Expect(len(nextList.Users)).To(Equal(1))
			Expect(nextList.Users).To(ContainElement(rainmaker.User{
				GUID:                           "user-next",
				URL:                            "/v2/users/user-next",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-next/spaces",
				OrganizationsURL:               "/v2/users/user-next/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-next/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-next/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-next/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-next/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-next/audited_spaces",
			}))
		})
	})

	Describe("Prev", func() {
		It("returns the previous UserList result for the paginated set", func() {
			list.PrevURL = "/v2/organizations/org-001/users?page=1"
			nextList, err := list.Prev("token")
			Expect(err).NotTo(HaveOccurred())

			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-11T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			Expect(nextList.TotalResults).To(Equal(4))
			Expect(nextList.TotalPages).To(Equal(3))
			Expect(nextList.Users).To(HaveLen(2))
			Expect(nextList.Users).To(ContainElement(rainmaker.User{
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

			Expect(nextList.Users).To(ContainElement(rainmaker.User{
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

	Describe("HasNextPage", func() {
		It("indicates whether or not there is a next page of results", func() {
			list.NextURL = "/v2/organizations/org-001/users?page=2"
			Expect(list.HasNextPage()).To(BeTrue())

			list.NextURL = ""
			Expect(list.HasNextPage()).To(BeFalse())
		})
	})

	Describe("HasPrevPage", func() {
		It("indicates whether or not there is a previous page of results", func() {
			list.PrevURL = "/v2/organizations/org-001/users?page=1"
			Expect(list.HasPrevPage()).To(BeTrue())

			list.PrevURL = ""
			Expect(list.HasPrevPage()).To(BeFalse())
		})
	})

	Describe("AllUsers", func() {
		It("returns a slice of all of users", func() {
			var err error
			list, err = rainmaker.FetchUsersList(config, "/v2/organizations/org-001/users?page=2", "token")
			Expect(err).NotTo(HaveOccurred())

			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-11T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			users, err := list.AllUsers("token")
			Expect(err).NotTo(HaveOccurred())

			Expect(users).To(HaveLen(4))
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

			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-next",
				URL:                            "/v2/users/user-next",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-next/spaces",
				OrganizationsURL:               "/v2/users/user-next/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-next/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-next/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-next/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-next/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-next/audited_spaces",
			}))

			Expect(users).To(ContainElement(rainmaker.User{
				GUID:                           "user-last",
				URL:                            "/v2/users/user-last",
				CreatedAt:                      userCreatedAt,
				UpdatedAt:                      time.Time{},
				Admin:                          false,
				Active:                         true,
				DefaultSpaceGUID:               "",
				SpacesURL:                      "/v2/users/user-last/spaces",
				OrganizationsURL:               "/v2/users/user-last/organizations",
				ManagedOrganizationsURL:        "/v2/users/user-last/managed_organizations",
				BillingManagedOrganizationsURL: "/v2/users/user-last/billing_managed_organizations",
				AuditedOrganizationsURL:        "/v2/users/user-last/audited_organizations",
				ManagedSpacesURL:               "/v2/users/user-last/managed_spaces",
				AuditedSpacesURL:               "/v2/users/user-last/audited_spaces",
			}))
		})
	})
})

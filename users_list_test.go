package rainmaker_test

import (
	"time"

	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UsersList", func() {
	Describe("Next", func() {
		var list rainmaker.UsersList

		BeforeEach(func() {
			client := rainmaker.NewClient(rainmaker.Config{
				Host: fakeCloudController.URL(),
			})
			list = rainmaker.NewUsersList(client)
		})

		It("returns the next UserList result for the paginated set", func() {
			list.NextURL = "/v2/organizations/org-001/users?page=2"
			nextList, err := list.Next("token")
			if err != nil {
				panic(err)
			}

			userCreatedAt, err := time.Parse(time.RFC3339, "2014-11-11T18:22:51+00:00")
			if err != nil {
				panic(err)
			}

			Expect(nextList.TotalResults).To(Equal(3))
			Expect(nextList.TotalPages).To(Equal(2))
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
})

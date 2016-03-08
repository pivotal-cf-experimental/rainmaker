package rainmaker_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrganizationsService", func() {
	var (
		config       rainmaker.Config
		token        string
		service      rainmaker.OrganizationsService
		organization rainmaker.Organization
	)

	BeforeEach(func() {
		var err error

		token = "token"
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		client := rainmaker.NewClient(config)
		service = client.Organizations

		organization, err = service.Create("test-org", token)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Create", func() {
		It("creates a new organization that can be fetched from the API", func() {
			organization, err := service.Create("my-new-org", token)
			Expect(err).NotTo(HaveOccurred())
			Expect(organization.Name).To(Equal("my-new-org"))

			fetchedOrg, err := service.Get(organization.GUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedOrg).To(Equal(organization))
		})
	})

	Describe("List", func() {
		var (
			org1 rainmaker.Organization
		)

		BeforeEach(func() {
			var err error

			org1, err = service.Create("first org", token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("retrieves a list of all orgs from the cloud controller", func() {
			list, err := service.List(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(2))
			Expect(list.TotalPages).To(Equal(1))
			Expect(list.Organizations).To(HaveLen(2))

			orgGuids := []string{
				list.Organizations[0].GUID,
				list.Organizations[1].GUID,
			}
			Expect(orgGuids).To(ConsistOf([]string{org1.GUID, organization.GUID}))
		})
	})

	Describe("Get", func() {
		It("returns the organization matching the given GUID", func() {
			var err error
			orgGUID := organization.GUID

			organization, err = service.Get(orgGUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(organization.GUID).To(Equal(orgGUID))
		})

		Context("when the request errors", func() {
			BeforeEach(func() {
				config.Host = ""
				service = rainmaker.NewOrganizationsService(config)
			})

			It("returns the error", func() {
				_, err := service.Create("org-name", token)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when unmarshalling fails", func() {
			It("returns an error", func() {
				_, err := service.Get("very-bad-guid", token)
				Expect(err).To(BeAssignableToTypeOf(rainmaker.Error{}))
			})
		})
	})

	Describe("Update", func() {
		It("updates the organization", func() {
			organization.Name = "some-updated-organization"
			organization.Status = "some-updated-status"
			organization.QuotaDefinitionGUID = "some-quota-definition-guid"

			updatedOrg, err := service.Update(organization, token)
			Expect(err).NotTo(HaveOccurred())

			fetchedOrg, err := service.Get(organization.GUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedOrg).To(Equal(updatedOrg))
		})

		Context("failure cases", func() {
			It("returns an error when malformed JSON is returned", func() {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("banana"))
				}))

				config = rainmaker.Config{
					Host: server.URL,
				}
				service = rainmaker.NewClient(config).Organizations

				_, err := service.Update(organization, token)
				Expect(err).To(BeAssignableToTypeOf(rainmaker.Error{}))
			})
			It("returns an error when receiving an unexpected status code", func() {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
				}))

				config = rainmaker.Config{
					Host: server.URL,
				}
				service = rainmaker.NewClient(config).Organizations

				_, err := service.Update(organization, token)
				Expect(err).To(BeAssignableToTypeOf(rainmaker.Error{}))
			})
		})
	})

	Describe("Delete", func() {
		It("deletes the organization", func() {
			err := service.Delete(organization.GUID, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = service.Get(organization.GUID, token)
			Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
		})

		Context("when the response status is unexpected", func() {
			It("returns an error", func() {
				err := service.Delete("very-bad-guid", token)
				Expect(err).To(BeAssignableToTypeOf(rainmaker.Error{}))
			})
		})
	})

	Context("when listing related spaces", func() {
		var space1, space2 rainmaker.Space

		BeforeEach(func() {
			var err error
			spacesService := rainmaker.NewSpacesService(config)

			space1, err = spacesService.Create("space-123", organization.GUID, token)
			Expect(err).NotTo(HaveOccurred())

			space2, err = spacesService.Create("space-456", organization.GUID, token)
			Expect(err).NotTo(HaveOccurred())
		})

		Describe("ListSpaces", func() {
			It("returns the spaces belonging to the organization", func() {
				list, err := service.ListSpaces(organization.GUID, token)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.TotalResults).To(Equal(2))
				Expect(list.TotalPages).To(Equal(1))
				Expect(list.Spaces).To(HaveLen(2))

				var spaceGUIDs []string
				for _, space := range list.Spaces {
					spaceGUIDs = append(spaceGUIDs, space.GUID)
				}

				Expect(spaceGUIDs).To(ConsistOf([]string{space1.GUID, space2.GUID}))
			})

			Context("when the organization does not exist", func() {
				It("returns an error", func() {
					_, err := service.ListSpaces("org-does-not-exist", token)
					Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
				})
			})
		})
	})

	Context("when listing related users", func() {
		var user1, user2, user3 rainmaker.User

		BeforeEach(func() {
			var err error
			usersService := rainmaker.NewUsersService(config)

			user1, err = usersService.Create("user-123", token)
			Expect(err).NotTo(HaveOccurred())

			user2, err = usersService.Create("user-456", token)
			Expect(err).NotTo(HaveOccurred())

			user3, err = usersService.Create("user-789", token)
			Expect(err).NotTo(HaveOccurred())
		})

		Describe("ListUsers", func() {
			BeforeEach(func() {
				err := organization.Users.Associate(user1.GUID, token)
				Expect(err).NotTo(HaveOccurred())

				err = organization.Users.Associate(user2.GUID, token)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the users belonging to the organization", func() {
				list, err := service.ListUsers(organization.GUID, token)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.TotalResults).To(Equal(2))
				Expect(list.TotalPages).To(Equal(1))
				Expect(list.Users).To(HaveLen(2))

				var userGUIDs []string
				for _, user := range list.Users {
					userGUIDs = append(userGUIDs, user.GUID)
				}

				Expect(userGUIDs).To(ConsistOf([]string{user1.GUID, user2.GUID}))
			})

			Context("when the organization does not exist", func() {
				It("returns an error", func() {
					_, err := service.ListUsers("org-does-not-exist", token)
					Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
				})
			})
		})

		Describe("ListBillingManagers", func() {
			BeforeEach(func() {
				err := organization.BillingManagers.Associate(user2.GUID, token)
				Expect(err).NotTo(HaveOccurred())

				err = organization.BillingManagers.Associate(user3.GUID, token)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the billing managers belonging to the organization", func() {
				list, err := service.ListBillingManagers(organization.GUID, token)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.TotalResults).To(Equal(2))
				Expect(list.TotalPages).To(Equal(1))
				Expect(list.Users).To(HaveLen(2))

				var userGUIDs []string
				for _, user := range list.Users {
					userGUIDs = append(userGUIDs, user.GUID)
				}

				Expect(userGUIDs).To(ConsistOf([]string{user2.GUID, user3.GUID}))
			})
		})

		Describe("ListAuditors", func() {
			BeforeEach(func() {
				err := organization.Auditors.Associate(user1.GUID, token)
				Expect(err).NotTo(HaveOccurred())

				err = organization.Auditors.Associate(user3.GUID, token)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the auditors belonging to the organization", func() {
				list, err := service.ListAuditors(organization.GUID, token)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.TotalResults).To(Equal(2))
				Expect(list.TotalPages).To(Equal(1))
				Expect(list.Users).To(HaveLen(2))

				var userGUIDs []string
				for _, user := range list.Users {
					userGUIDs = append(userGUIDs, user.GUID)
				}

				Expect(userGUIDs).To(ConsistOf([]string{user1.GUID, user3.GUID}))
			})
		})

		Describe("ListManagers", func() {
			BeforeEach(func() {
				err := organization.Managers.Associate(user3.GUID, token)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the managers belonging to the organization", func() {
				list, err := service.ListManagers(organization.GUID, token)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.TotalResults).To(Equal(1))
				Expect(list.TotalPages).To(Equal(1))
				Expect(list.Users).To(HaveLen(1))

				Expect(list.Users[0].GUID).To(Equal(user3.GUID))
			})
		})

	})
})

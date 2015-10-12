package rainmaker_test

import (
	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpacesService", func() {
	var (
		config    rainmaker.Config
		service   *rainmaker.SpacesService
		token     string
		spaceName string
		org       rainmaker.Organization
	)

	BeforeEach(func() {
		var err error
		token = "token-asd"
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		service = rainmaker.NewSpacesService(config)
		spaceName = "my-space"

		org, err = rainmaker.NewOrganizationsService(config).Create("org-123", token)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Create/Get", func() {
		It("creates a space and allows it to be fetched from the cloud controller", func() {
			space, err := service.Create(spaceName, org.GUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(space.Name).To(Equal(spaceName))
			Expect(space.OrganizationGUID).To(Equal(org.GUID))

			fetchedSpace, err := service.Get(space.GUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedSpace.GUID).To(Equal(space.GUID))
		})

		Context("when the request errors", func() {
			BeforeEach(func() {
				config.Host = ""
				service = rainmaker.NewSpacesService(config)
			})

			It("returns the error", func() {
				_, err := service.Create("space-name", "org-guid", token)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete", func() {
		It("deletes a space", func() {
			space, err := service.Create(spaceName, org.GUID, token)
			Expect(err).NotTo(HaveOccurred())

			err = service.Delete(space.GUID, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = service.Get(space.GUID, token)
			Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
		})
	})

	Describe("ListUsers", func() {
		var (
			user  rainmaker.User
			space rainmaker.Space
		)

		BeforeEach(func() {
			var err error

			user, err = rainmaker.NewUsersService(config).Create("user-abc", token)
			Expect(err).NotTo(HaveOccurred())

			_, err = rainmaker.NewUsersService(config).Create("user-xyz", token)
			Expect(err).NotTo(HaveOccurred())

			space, err = service.Create(spaceName, org.GUID, token)
			Expect(err).NotTo(HaveOccurred())

			err = space.Developers.Associate(user.GUID, token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the users belonging to the space", func() {
			list, err := service.ListUsers(space.GUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(1))
			Expect(list.TotalPages).To(Equal(1))
			Expect(list.Users).To(HaveLen(1))

			Expect(list.Users[0].GUID).To(Equal(user.GUID))
		})
	})
})

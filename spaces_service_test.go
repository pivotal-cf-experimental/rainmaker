package rainmaker_test

import (
	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpacesService", func() {
	var config rainmaker.Config
	var service *rainmaker.SpacesService

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		service = rainmaker.NewSpacesService(config)
	})

	Describe("Get", func() {
		It("returns the space matching the given GUID", func() {
			space, err := service.Get("space-001", "token-asd")
			Expect(err).NotTo(HaveOccurred())
			Expect(space.GUID).To(Equal("space-001"))
		})
	})

	Describe("ListUsers", func() {
		It("returns the users belonging to the space", func() {
			list, err := service.ListUsers("space-001", "token-abc")
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(2))
			Expect(list.TotalPages).To(Equal(1))
			Expect(list.Users).To(HaveLen(2))

			var userGUIDs []string
			for _, user := range list.Users {
				userGUIDs = append(userGUIDs, user.GUID)
			}

			Expect(userGUIDs).To(ContainElement("user-abc"))
			Expect(userGUIDs).To(ContainElement("user-xyz"))
		})
	})
})

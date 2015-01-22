package rainmaker_test

import (
	"net/url"

	"github.com/pivotal-golang/rainmaker"
	"github.com/pivotal-golang/rainmaker/internal/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpacesService", func() {
	var config rainmaker.Config
	var service *rainmaker.SpacesService
	var token, spaceGUID string

	BeforeEach(func() {
		token = "token-asd"
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		service = rainmaker.NewSpacesService(config)
		spaceGUID = "space-001"

		fakeCloudController.Spaces.Add(fakes.Space{
			GUID:       spaceGUID,
			Developers: fakes.NewUsers(),
		})
	})

	Describe("Get", func() {
		It("returns the space matching the given GUID", func() {
			space, err := service.Get(spaceGUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(space.GUID).To(Equal(spaceGUID))
		})
	})

	Describe("ListUsers", func() {
		BeforeEach(func() {
			list := rainmaker.NewUsersList(config, rainmaker.NewRequestPlan("/v2/users", url.Values{}))
			user, err := list.Create(rainmaker.User{GUID: "user-abc"}, token)
			if err != nil {
				panic(err)
			}

			space, err := service.Get(spaceGUID, token)
			if err != nil {
				panic(err)
			}

			err = space.Developers.Associate(user.GUID, token)
			if err != nil {
				panic(err)
			}

		})

		It("returns the users belonging to the space", func() {
			list, err := service.ListUsers(spaceGUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(1))
			Expect(list.TotalPages).To(Equal(1))
			Expect(list.Users).To(HaveLen(1))

			Expect(list.Users[0].GUID).To(Equal("user-abc"))
		})
	})
})

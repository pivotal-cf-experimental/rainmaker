package rainmaker_test

import (
	"fmt"
	"net/url"

	"github.com/pivotal-cf-experimental/rainmaker"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpacesList", func() {
	var (
		config      rainmaker.Config
		path, token string
		list        rainmaker.SpacesList
		orgGUID     string
	)

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		path = "/v2/spaces"
		token = "token"
		query := url.Values{}
		query.Add("results-per-page", "2")
		query.Add("page", "1")
		list = rainmaker.NewSpacesList(config, rainmaker.NewRequestPlan(path, query))

		orgGUID = "org-abc"
		fakeCloudController.Organizations.Add(domain.Organization{
			GUID:   orgGUID,
			Spaces: domain.NewSpaces(),
		})
	})

	AfterEach(func() {
		fakeCloudController.Organizations.Delete(orgGUID)
	})

	Describe("Create", func() {
		It("adds a space to the list", func() {
			space, err := list.Create(rainmaker.Space{
				GUID:             "space-123",
				OrganizationGUID: orgGUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(space.GUID).To(Equal("space-123"))

			err = list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(list.Spaces).To(HaveLen(1))
			Expect(list.Spaces[0].GUID).To(Equal("space-123"))
		})
	})

	Describe("Next", func() {
		BeforeEach(func() {
			_, err := list.Create(rainmaker.Space{
				GUID:             "space-123",
				OrganizationGUID: orgGUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Space{
				GUID:             "space-456",
				OrganizationGUID: orgGUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Space{
				GUID:             "space-789",
				OrganizationGUID: orgGUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the next SpaceList result for the paginated set", func() {
			err := list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(list.Spaces).To(HaveLen(2))
			Expect(list.HasNextPage()).To(BeTrue())
			Expect(list.HasPrevPage()).To(BeFalse())
			Expect(list.TotalResults).To(Equal(3))
			Expect(list.TotalPages).To(Equal(2))

			nextList, err := list.Next(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(nextList.Spaces).To(HaveLen(1))
			Expect(nextList.HasNextPage()).To(BeFalse())
			Expect(nextList.HasPrevPage()).To(BeTrue())
			Expect(nextList.TotalResults).To(Equal(3))
			Expect(nextList.TotalPages).To(Equal(2))

			var spaces []rainmaker.Space
			spaces = append(spaces, list.Spaces...)
			spaces = append(spaces, nextList.Spaces...)
			Expect(spaces).To(HaveLen(3))

			var guids []string
			for _, space := range spaces {
				guids = append(guids, space.GUID)
			}
			Expect(guids).To(ConsistOf([]string{"space-123", "space-456", "space-789"}))
		})
	})

	Describe("Prev", func() {
		BeforeEach(func() {
			_, err := list.Create(rainmaker.Space{
				GUID:             "space-qwe",
				OrganizationGUID: orgGUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Space{
				GUID:             "space-rty",
				OrganizationGUID: orgGUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Space{
				GUID:             "space-uio",
				OrganizationGUID: orgGUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the previous SpaceList result for the paginated set", func() {
			query := url.Values{}
			query.Set("page", "2")
			query.Set("results-per-page", "2")

			list := rainmaker.NewSpacesList(config, rainmaker.NewRequestPlan(path, query))
			err := list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(list.Spaces).To(HaveLen(1))
			Expect(list.HasNextPage()).To(BeFalse())
			Expect(list.HasPrevPage()).To(BeTrue())
			Expect(list.TotalResults).To(Equal(3))
			Expect(list.TotalPages).To(Equal(2))

			prevList, err := list.Prev(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(prevList.Spaces).To(HaveLen(2))
			Expect(prevList.HasNextPage()).To(BeTrue())
			Expect(prevList.HasPrevPage()).To(BeFalse())
			Expect(prevList.TotalResults).To(Equal(3))
			Expect(prevList.TotalPages).To(Equal(2))

			var spaces []rainmaker.Space
			spaces = append(spaces, list.Spaces...)
			spaces = append(spaces, prevList.Spaces...)
			Expect(spaces).To(HaveLen(3))

			var guids []string
			for _, space := range spaces {
				guids = append(guids, space.GUID)
			}
			Expect(guids).To(ConsistOf([]string{"space-qwe", "space-rty", "space-uio"}))
		})
	})

	Describe("HasNextPage", func() {
		It("indicates whether or not there is a next page of results", func() {
			list.NextURL = "/v2/spaces?page=2"
			Expect(list.HasNextPage()).To(BeTrue())

			list.NextURL = ""
			Expect(list.HasNextPage()).To(BeFalse())
		})
	})

	Describe("HasPrevPage", func() {
		It("indicates whether or not there is a previous page of results", func() {
			list.PrevURL = "/v2/spaces?page=1"
			Expect(list.HasPrevPage()).To(BeTrue())

			list.PrevURL = ""
			Expect(list.HasPrevPage()).To(BeFalse())
		})
	})

	Describe("AllSpaces", func() {
		BeforeEach(func() {
			for i := 0; i < 10; i++ {
				_, err := list.Create(rainmaker.Space{
					GUID:             fmt.Sprintf("space-%d", i),
					OrganizationGUID: orgGUID,
				}, token)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("returns a slice of all of spaces", func() {
			err := list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			list, err = list.Next(token)
			Expect(err).NotTo(HaveOccurred())

			spaces, err := list.AllSpaces(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(spaces).To(HaveLen(10))
			var guids []string
			for _, space := range spaces {
				guids = append(guids, space.GUID)
			}
			Expect(guids).To(ConsistOf([]string{
				"space-0",
				"space-1",
				"space-2",
				"space-3",
				"space-4",
				"space-5",
				"space-6",
				"space-7",
				"space-8",
				"space-9",
			}))
		})
	})
})

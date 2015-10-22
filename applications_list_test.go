package rainmaker_test

import (
	"net/url"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ApplicationsList", func() {
	var (
		config      rainmaker.Config
		client      rainmaker.Client
		path, token string
		list        rainmaker.ApplicationsList
		space       rainmaker.Space
	)

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		path = "/v2/apps"
		token = "token"
		query := url.Values{
			"page":             {"1"},
			"results-per-page": {"2"},
		}
		list = rainmaker.NewApplicationsList(config, rainmaker.NewRequestPlan(path, query))

		client = rainmaker.NewClient(config)
		org, err := client.Organizations.Create("some-org-name", token)
		Expect(err).NotTo(HaveOccurred())

		space, err = client.Spaces.Create("some-space-name", org.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("pagination", func() {
		var (
			app1 rainmaker.Application
			app2 rainmaker.Application
			app3 rainmaker.Application
		)

		BeforeEach(func() {
			var err error

			app1, err = client.Applications.Create(rainmaker.Application{
				Name:      "some-app-name",
				SpaceGUID: space.GUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())

			app2, err = client.Applications.Create(rainmaker.Application{
				Name:      "some-other-app-name",
				SpaceGUID: space.GUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())

			app3, err = client.Applications.Create(rainmaker.Application{
				Name:      "a-different-app-name",
				SpaceGUID: space.GUID,
			}, token)
			Expect(err).NotTo(HaveOccurred())
		})

		Describe("Next", func() {
			It("returns the next ApplicationsList result for the paginated set", func() {
				err := list.Fetch(token)
				Expect(err).NotTo(HaveOccurred())

				Expect(list.Applications).To(HaveLen(2))
				Expect(list.HasNextPage()).To(BeTrue())
				Expect(list.HasPrevPage()).To(BeFalse())
				Expect(list.TotalResults).To(Equal(3))
				Expect(list.TotalPages).To(Equal(2))

				nextList, err := list.Next(token)
				Expect(err).NotTo(HaveOccurred())
				Expect(nextList.Applications).To(HaveLen(1))
				Expect(nextList.HasNextPage()).To(BeFalse())
				Expect(nextList.HasPrevPage()).To(BeTrue())
				Expect(nextList.TotalResults).To(Equal(3))
				Expect(nextList.TotalPages).To(Equal(2))

				var apps []rainmaker.Application
				apps = append(apps, list.Applications...)
				apps = append(apps, nextList.Applications...)
				Expect(apps).To(HaveLen(3))

				var guids []string
				for _, app := range apps {
					guids = append(guids, app.GUID)
				}
				Expect(guids).To(ConsistOf([]string{app1.GUID, app2.GUID, app3.GUID}))
			})
		})

		Describe("Prev", func() {
			It("returns the previous ApplicationsList result for the paginated set", func() {
				query := url.Values{}
				query.Set("page", "2")
				query.Set("results-per-page", "2")

				list := rainmaker.NewApplicationsList(config, rainmaker.NewRequestPlan(path, query))
				err := list.Fetch(token)
				Expect(err).NotTo(HaveOccurred())

				Expect(list.Applications).To(HaveLen(1))
				Expect(list.HasNextPage()).To(BeFalse())
				Expect(list.HasPrevPage()).To(BeTrue())
				Expect(list.TotalResults).To(Equal(3))
				Expect(list.TotalPages).To(Equal(2))

				prevList, err := list.Prev(token)
				Expect(err).NotTo(HaveOccurred())
				Expect(prevList.Applications).To(HaveLen(2))
				Expect(prevList.HasNextPage()).To(BeTrue())
				Expect(prevList.HasPrevPage()).To(BeFalse())
				Expect(prevList.TotalResults).To(Equal(3))
				Expect(prevList.TotalPages).To(Equal(2))

				var apps []rainmaker.Application
				apps = append(apps, list.Applications...)
				apps = append(apps, prevList.Applications...)
				Expect(apps).To(HaveLen(3))

				var guids []string
				for _, app := range apps {
					guids = append(guids, app.GUID)
				}
				Expect(guids).To(ConsistOf([]string{app1.GUID, app2.GUID, app3.GUID}))
			})
		})
	})

	Describe("HasNextPage", func() {
		It("indicates whether or not there is a next page of results", func() {
			list.NextURL = "/v2/apps?page=2"
			Expect(list.HasNextPage()).To(BeTrue())

			list.NextURL = ""
			Expect(list.HasNextPage()).To(BeFalse())
		})
	})

	Describe("HasPrevPage", func() {
		It("indicates whether or not there is a previous page of results", func() {
			list.PrevURL = "/v2/apps?page=1"
			Expect(list.HasPrevPage()).To(BeTrue())

			list.PrevURL = ""
			Expect(list.HasPrevPage()).To(BeFalse())
		})
	})
})

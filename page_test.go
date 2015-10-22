package rainmaker_test

import (
	"encoding/json"
	"net/url"

	"github.com/pivotal-cf-experimental/rainmaker"
	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Page", func() {
	var (
		config      rainmaker.Config
		path, token string
		page        rainmaker.Page
	)

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}

		path = "/v2/users"
		token = "token"
		query := url.Values{
			"results-per-page": {"2"},
			"page":             {"1"},
		}
		page = rainmaker.NewPage(config, rainmaker.NewRequestPlan(path, query))
	})

	Describe("Next", func() {
		BeforeEach(func() {
			usersService := rainmaker.NewUsersService(config)

			_, err := usersService.Create("user-123", token)
			Expect(err).NotTo(HaveOccurred())

			_, err = usersService.Create("user-456", token)
			Expect(err).NotTo(HaveOccurred())

			_, err = usersService.Create("user-789", token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the next page of resources for the paginated set", func() {
			err := page.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(page.Resources).To(HaveLen(2))
			Expect(page.HasNextPage()).To(BeTrue())
			Expect(page.HasPrevPage()).To(BeFalse())
			Expect(page.TotalResults).To(Equal(3))
			Expect(page.TotalPages).To(Equal(2))

			nextPage, err := page.Next(token)
			Expect(err).NotTo(HaveOccurred())

			err = nextPage.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(nextPage.Resources).To(HaveLen(1))
			Expect(nextPage.HasNextPage()).To(BeFalse())
			Expect(nextPage.HasPrevPage()).To(BeTrue())
			Expect(nextPage.TotalResults).To(Equal(3))
			Expect(nextPage.TotalPages).To(Equal(2))

			var resources []json.RawMessage
			resources = append(resources, page.Resources...)
			resources = append(resources, nextPage.Resources...)
			Expect(resources).To(HaveLen(3))

			var guids []string
			for _, resource := range resources {
				var user documents.UserResponse
				err := json.Unmarshal(resource, &user)
				Expect(err).NotTo(HaveOccurred())

				guids = append(guids, user.Metadata.GUID)
			}
			Expect(guids).To(ConsistOf([]string{"user-123", "user-456", "user-789"}))
		})
	})

	Describe("Prev", func() {
		BeforeEach(func() {
			usersService := rainmaker.NewUsersService(config)

			_, err := usersService.Create("user-123", token)
			Expect(err).NotTo(HaveOccurred())

			_, err = usersService.Create("user-456", token)
			Expect(err).NotTo(HaveOccurred())

			_, err = usersService.Create("user-789", token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the next page of resources for the paginated set", func() {
			query := url.Values{
				"results-per-page": {"2"},
				"page":             {"2"},
			}

			page = rainmaker.NewPage(config, rainmaker.NewRequestPlan(path, query))
			err := page.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(page.Resources).To(HaveLen(1))
			Expect(page.HasNextPage()).To(BeFalse())
			Expect(page.HasPrevPage()).To(BeTrue())
			Expect(page.TotalResults).To(Equal(3))
			Expect(page.TotalPages).To(Equal(2))

			prevPage, err := page.Prev(token)
			Expect(err).NotTo(HaveOccurred())

			err = prevPage.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(prevPage.Resources).To(HaveLen(2))
			Expect(prevPage.HasNextPage()).To(BeTrue())
			Expect(prevPage.HasPrevPage()).To(BeFalse())
			Expect(prevPage.TotalResults).To(Equal(3))
			Expect(prevPage.TotalPages).To(Equal(2))

			var resources []json.RawMessage
			resources = append(resources, prevPage.Resources...)
			resources = append(resources, page.Resources...)
			Expect(resources).To(HaveLen(3))

			var guids []string
			for _, resource := range resources {
				var user documents.UserResponse
				err := json.Unmarshal(resource, &user)
				Expect(err).NotTo(HaveOccurred())

				guids = append(guids, user.Metadata.GUID)
			}
			Expect(guids).To(ConsistOf([]string{"user-123", "user-456", "user-789"}))
		})
	})

	Describe("HasNextPage", func() {
		It("indicates whether or not there is a next page of results", func() {
			page.NextURL = "/v2/users?page=2"
			Expect(page.HasNextPage()).To(BeTrue())

			page.NextURL = ""
			Expect(page.HasNextPage()).To(BeFalse())
		})
	})

	Describe("HasPrevPage", func() {
		It("indicates whether or not there is a previous page of results", func() {
			page.PrevURL = "/v2/users?page=1"
			Expect(page.HasPrevPage()).To(BeTrue())

			page.PrevURL = ""
			Expect(page.HasPrevPage()).To(BeFalse())
		})
	})
})

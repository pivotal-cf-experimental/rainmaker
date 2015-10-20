package rainmaker_test

import (
	"fmt"
	"net/url"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrganizationsList", func() {
	var (
		config      rainmaker.Config
		path, token string
		list        rainmaker.OrganizationsList
	)

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		path = "/v2/organizations"
		token = "token"
		query := url.Values{}
		query.Add("page", "1")
		query.Add("results-per-page", "2")
		list = rainmaker.NewOrganizationsList(config, rainmaker.NewRequestPlan(path, query))
	})

	Describe("Create", func() {
		It("adds an organization to the list", func() {
			org, err := list.Create(rainmaker.Organization{
				GUID: "org-123",
			}, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(org.GUID).To(Equal("org-123"))

			err = list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(list.Organizations).To(HaveLen(1))
			Expect(list.Organizations[0].GUID).To(Equal("org-123"))
		})
	})

	Describe("Next", func() {
		BeforeEach(func() {
			_, err := list.Create(rainmaker.Organization{
				GUID: "org-123",
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Organization{
				GUID: "org-456",
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Organization{
				GUID: "org-789",
			}, token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the next OrganizationList result for the paginated set", func() {
			err := list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(list.Organizations).To(HaveLen(2))
			Expect(list.HasNextPage()).To(BeTrue())
			Expect(list.HasPrevPage()).To(BeFalse())
			Expect(list.TotalResults).To(Equal(3))
			Expect(list.TotalPages).To(Equal(2))

			nextList, err := list.Next(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(nextList.Organizations).To(HaveLen(1))
			Expect(nextList.HasNextPage()).To(BeFalse())
			Expect(nextList.HasPrevPage()).To(BeTrue())
			Expect(nextList.TotalResults).To(Equal(3))
			Expect(nextList.TotalPages).To(Equal(2))

			var orgs []rainmaker.Organization
			orgs = append(orgs, list.Organizations...)
			orgs = append(orgs, nextList.Organizations...)
			Expect(orgs).To(HaveLen(3))

			var guids []string
			for _, org := range orgs {
				guids = append(guids, org.GUID)
			}
			Expect(guids).To(ConsistOf([]string{"org-123", "org-456", "org-789"}))
		})
	})

	Describe("Prev", func() {
		BeforeEach(func() {
			_, err := list.Create(rainmaker.Organization{
				GUID: "org-qwe",
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Organization{
				GUID: "org-rty",
			}, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = list.Create(rainmaker.Organization{
				GUID: "org-uio",
			}, token)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the previous OrganizationList result for the paginated set", func() {
			query := url.Values{}
			query.Set("page", "2")
			query.Set("results-per-page", "2")

			list := rainmaker.NewOrganizationsList(config, rainmaker.NewRequestPlan(path, query))
			err := list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(list.Organizations).To(HaveLen(1))
			Expect(list.HasNextPage()).To(BeFalse())
			Expect(list.HasPrevPage()).To(BeTrue())
			Expect(list.TotalResults).To(Equal(3))
			Expect(list.TotalPages).To(Equal(2))

			prevList, err := list.Prev(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(prevList.Organizations).To(HaveLen(2))
			Expect(prevList.HasNextPage()).To(BeTrue())
			Expect(prevList.HasPrevPage()).To(BeFalse())
			Expect(prevList.TotalResults).To(Equal(3))
			Expect(prevList.TotalPages).To(Equal(2))

			var orgs []rainmaker.Organization
			orgs = append(orgs, list.Organizations...)
			orgs = append(orgs, prevList.Organizations...)
			Expect(orgs).To(HaveLen(3))

			var guids []string
			for _, org := range orgs {
				guids = append(guids, org.GUID)
			}
			Expect(guids).To(ConsistOf([]string{"org-qwe", "org-rty", "org-uio"}))
		})
	})

	Describe("HasNextPage", func() {
		It("indicates whether or not there is a next page of results", func() {
			list.NextURL = "/v2/organizations?page=2"
			Expect(list.HasNextPage()).To(BeTrue())

			list.NextURL = ""
			Expect(list.HasNextPage()).To(BeFalse())
		})
	})

	Describe("HasPrevPage", func() {
		It("indicates whether or not there is a previous page of results", func() {
			list.PrevURL = "/v2/organizations?page=1"
			Expect(list.HasPrevPage()).To(BeTrue())

			list.PrevURL = ""
			Expect(list.HasPrevPage()).To(BeFalse())
		})
	})

	Describe("AllOrganizations", func() {
		BeforeEach(func() {
			for i := 0; i < 10; i++ {
				_, err := list.Create(rainmaker.Organization{
					GUID: fmt.Sprintf("org-%d", i),
				}, token)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("returns a slice of all of organizations", func() {
			err := list.Fetch(token)
			Expect(err).NotTo(HaveOccurred())

			list, err = list.Next(token)
			Expect(err).NotTo(HaveOccurred())

			orgs, err := list.AllOrganizations(token)
			Expect(err).NotTo(HaveOccurred())

			Expect(orgs).To(HaveLen(10))
			var guids []string
			for _, org := range orgs {
				guids = append(guids, org.GUID)
			}
			Expect(guids).To(ConsistOf([]string{
				"org-0",
				"org-1",
				"org-2",
				"org-3",
				"org-4",
				"org-5",
				"org-6",
				"org-7",
				"org-8",
				"org-9",
			}))
		})
	})
})

package rainmaker_test

import (
	"github.com/pivotal-golang/rainmaker"
	"github.com/pivotal-golang/rainmaker/internal/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrganizationsService", func() {
	var config rainmaker.Config
	var service *rainmaker.OrganizationsService

	BeforeEach(func() {
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		client := rainmaker.NewClient(config)
		service = client.Organizations
	})

	Describe("Get", func() {
		BeforeEach(func() {
			fakeCloudController.Organizations.Add(fakes.Organization{
				GUID: "org-001",
			})
		})

		It("returns the organization matching the given GUID", func() {
			organization, err := service.Get("org-001", "token-123")
			Expect(err).NotTo(HaveOccurred())

			Expect(organization.GUID).To(Equal("org-001"))
		})
	})

	Describe("ListUsers", func() {
		It("returns the users belonging to the organization", func() {
			list, err := service.ListUsers("org-001", "token-456")
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(4))
			Expect(list.TotalPages).To(Equal(3))
			Expect(list.Users).To(HaveLen(2))
			var userGUIDs []string
			for _, user := range list.Users {
				userGUIDs = append(userGUIDs, user.GUID)
			}

			Expect(userGUIDs).To(ContainElement("user-123"))
			Expect(userGUIDs).To(ContainElement("user-456"))
		})
	})

	Describe("ListBillingManagers", func() {
		It("returns the billing managers belonging to the organization", func() {
			list, err := service.ListBillingManagers("org-002", "token-987")
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(2))
			Expect(list.TotalPages).To(Equal(1))
			Expect(list.Users).To(HaveLen(2))

			var userGUIDs []string
			for _, user := range list.Users {
				userGUIDs = append(userGUIDs, user.GUID)
			}

			Expect(userGUIDs).To(ContainElement("user-987"))
			Expect(userGUIDs).To(ContainElement("user-654"))
		})
	})

	Describe("ListAuditors", func() {
		It("returns the auditors belonging to the organization", func() {
			list, err := service.ListAuditors("org-003", "token-012")
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(2))
			Expect(list.TotalPages).To(Equal(1))
			Expect(list.Users).To(HaveLen(2))

			var userGUIDs []string
			for _, user := range list.Users {
				userGUIDs = append(userGUIDs, user.GUID)
			}

			Expect(userGUIDs).To(ContainElement("user-asd"))
			Expect(userGUIDs).To(ContainElement("user-jkl"))
		})
	})

	Describe("ListManagers", func() {
		It("returns the managers belonging to the organization", func() {
			list, err := service.ListManagers("org-004", "token-345")
			Expect(err).NotTo(HaveOccurred())
			Expect(list.TotalResults).To(Equal(2))
			Expect(list.TotalPages).To(Equal(1))
			Expect(list.Users).To(HaveLen(2))
			var userGUIDs []string
			for _, user := range list.Users {
				userGUIDs = append(userGUIDs, user.GUID)
			}

			Expect(userGUIDs).To(ContainElement("user-aaa"))
			Expect(userGUIDs).To(ContainElement("user-bbb"))
		})
	})
})

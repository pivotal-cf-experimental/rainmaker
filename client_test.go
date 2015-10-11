package rainmaker_test

import (
	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var client rainmaker.Client

	BeforeEach(func() {
		client = rainmaker.NewClient(rainmaker.Config{
			Host: "http://example.com",
		})
	})

	It("has an organizations service", func() {
		Expect(client.Organizations).To(BeAssignableToTypeOf(&rainmaker.OrganizationsService{}))
		Expect(client.Organizations).NotTo(BeNil())
	})

	It("has an spaces service", func() {
		Expect(client.Spaces).To(BeAssignableToTypeOf(&rainmaker.SpacesService{}))
		Expect(client.Spaces).NotTo(BeNil())
	})

	It("has a service instance service", func() {
		Expect(client.ServiceInstances).To(BeAssignableToTypeOf(&rainmaker.ServiceInstancesService{}))
		Expect(client.ServiceInstances).NotTo(BeNil())
	})
})

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

	It("has a organizations service", func() {
		Expect(client.Organizations).To(BeAssignableToTypeOf(rainmaker.OrganizationsService{}))
	})

	It("has a spaces service", func() {
		Expect(client.Spaces).To(BeAssignableToTypeOf(rainmaker.SpacesService{}))
	})

	It("has a users service", func() {
		Expect(client.Users).To(BeAssignableToTypeOf(rainmaker.UsersService{}))
	})

	It("has a service instance service", func() {
		Expect(client.ServiceInstances).To(BeAssignableToTypeOf(rainmaker.ServiceInstancesService{}))
	})
})

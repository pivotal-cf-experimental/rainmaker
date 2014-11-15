package rainmaker_test

import (
	"github.com/pivotal-golang/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var client rainmaker.Client
	var host string

	BeforeEach(func() {
		host = "hostURL"
		client = rainmaker.NewClient(rainmaker.Config{Host: host})
	})

	It("has an organizations service", func() {
		Expect(client.Organizations).To(BeAssignableToTypeOf(&rainmaker.OrganizationsService{}))
		Expect(client.Organizations).NotTo(BeNil())
	})

	It("has an spaces service", func() {
		Expect(client.Spaces).To(BeAssignableToTypeOf(&rainmaker.SpacesService{}))
		Expect(client.Spaces).NotTo(BeNil())
	})

})

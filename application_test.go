package rainmaker_test

import (
	"github.com/pivotal-cf-experimental/rainmaker"
	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application", func() {
	var config rainmaker.Config

	Describe("NewApplicationFromResponse", func() {
		BeforeEach(func() {
			config = rainmaker.Config{
				Host: fakeCloudController.URL(),
			}
		})

		It("converts a response into an application", func() {
			document := documents.ApplicationResponse{}
			document.Metadata.GUID = "cool-app"
			document.Entity.Name = "my cool app"
			document.Entity.SpaceGUID = "space-123"
			document.Entity.Diego = true

			application := rainmaker.NewApplicationFromResponse(config, document)
			expectedApplication := rainmaker.NewApplication(config, "cool-app")
			expectedApplication.Name = "my cool app"
			expectedApplication.SpaceGUID = "space-123"
			expectedApplication.Diego = true

			Expect(application).To(Equal(expectedApplication))
		})
	})
})

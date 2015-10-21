package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application", func() {
	var (
		token  string
		client rainmaker.Client
		org    rainmaker.Organization
		space  rainmaker.Space
		app    rainmaker.Application
	)

	BeforeEach(func() {
		token = os.Getenv("UAA_TOKEN")
		client = rainmaker.NewClient(rainmaker.Config{
			Host:          os.Getenv("CC_HOST"),
			SkipVerifySSL: true,
		})
	})

	AfterEach(func() {
		err := client.Spaces.Delete(space.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		err = client.Organizations.Delete(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})

	It("fetches the app", func() {
		By("creating an organization", func() {
			var err error
			org, err = client.Organizations.Create(NewGUID("org"), token)
			Expect(err).NotTo(HaveOccurred())
		})

		By("creating a space", func() {
			var err error
			space, err = client.Spaces.Create(NewGUID("space"), org.GUID, token)
			Expect(err).NotTo(HaveOccurred())
		})

		By("creating an app", func() {
			var err error
			app, err = client.Applications.Create(rainmaker.Application{
				Name:      "some-test-app",
				SpaceGUID: space.GUID,
				Diego:     true,
			}, token)
			Expect(err).NotTo(HaveOccurred())
		})

		By("getting the app", func() {
			fetchedApp, err := client.Applications.Get(app.GUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(fetchedApp.GUID).To(Equal(app.GUID))
			Expect(fetchedApp.Name).To(Equal(app.Name))
			Expect(fetchedApp.SpaceGUID).To(Equal(app.SpaceGUID))
			Expect(fetchedApp.Diego).To(BeTrue())
		})

		By("deleting the app", func() {
			Expect(client.Applications.Delete(app.GUID, token)).To(Succeed())
		})
	})
})

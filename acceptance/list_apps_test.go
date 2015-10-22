package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List all apps", func() {
	var (
		token  string
		client rainmaker.Client
		org    rainmaker.Organization
		space  rainmaker.Space
	)

	BeforeEach(func() {
		token = os.Getenv("UAA_TOKEN")
		client = rainmaker.NewClient(rainmaker.Config{
			Host:          os.Getenv("CC_HOST"),
			SkipVerifySSL: true,
		})

		var err error
		org, err = client.Organizations.Create(NewGUID("org"), token)
		Expect(err).NotTo(HaveOccurred())

		space, err = client.Spaces.Create(NewGUID("space"), org.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		err := client.Spaces.Delete(space.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		err = client.Organizations.Delete(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})

	It("lists all apps", func() {
		app, err := client.Applications.Create(rainmaker.Application{
			Name:      NewGUID("app"),
			SpaceGUID: space.GUID,
		}, token)
		Expect(err).NotTo(HaveOccurred())

		list, err := client.Applications.List(token)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(list.Applications)).To(BeNumerically(">", 0))

		var guids []string
		for _, app := range list.Applications {
			guids = append(guids, app.GUID)
		}
		Expect(guids).To(ContainElement(app.GUID))

		err = client.Applications.Delete(app.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})
})

package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fetch a space", func() {
	var (
		token  string
		client rainmaker.Client
		org    rainmaker.Organization
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

	})

	It("fetches the space", func() {
		space, err := client.Spaces.Create(NewGUID("space"), org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		fetchedSpace, err := client.Spaces.Get(space.GUID, token)
		Expect(err).NotTo(HaveOccurred())
		Expect(fetchedSpace).To(Equal(space))

		err = client.Spaces.Delete(space.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})
})

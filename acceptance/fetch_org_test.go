package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fetch an organization", func() {
	var (
		token  string
		client rainmaker.Client
	)

	BeforeEach(func() {
		token = os.Getenv("UAA_TOKEN")
		client = rainmaker.NewClient(rainmaker.Config{
			Host:          os.Getenv("CC_HOST"),
			SkipVerifySSL: true,
		})
	})

	It("fetches the organization", func() {
		org, err := client.Organizations.Create(NewGUID("org"), token)
		Expect(err).NotTo(HaveOccurred())

		fetchedOrg, err := client.Organizations.Get(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())
		Expect(fetchedOrg).To(Equal(org))

		err = client.Organizations.Delete(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		_, err = client.Organizations.Get(org.GUID, token)
		Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
	})
})

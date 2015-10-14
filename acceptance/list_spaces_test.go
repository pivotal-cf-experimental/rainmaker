package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List all spaces", func() {
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

	AfterEach(func() {
		err := client.Organizations.Delete(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})

	It("lists all spaces", func() {
		space, err := client.Spaces.Create(NewGUID("space"), org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		list, err := client.Spaces.List(token)
		Expect(err).NotTo(HaveOccurred())

		// This seems unneccessarily vague, but it is hard to be more specific
		// when testing on anything but a local bosh-lite environment.
		Expect(list.Spaces).NotTo(HaveLen(0))

		err = client.Spaces.Delete(space.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})
})

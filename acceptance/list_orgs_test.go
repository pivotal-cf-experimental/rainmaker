package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List all organizations", func() {
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

	It("lists all organizations", func() {
		org, err := client.Organizations.Create(NewGUID("org"), token)
		Expect(err).NotTo(HaveOccurred())

		list, err := client.Organizations.List(token)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(list.Organizations)).To(BeNumerically(">", 0))

		var guids []string
		for _, org := range list.Organizations {
			guids = append(guids, org.GUID)
		}
		Expect(guids).To(ContainElement(org.GUID))

		err = client.Organizations.Delete(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())
	})
})

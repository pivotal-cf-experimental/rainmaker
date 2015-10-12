package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fetch all the auditors of an organization", func() {
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

	It("fetches the user records of all auditors associated with an organization", func() {
		user, err := client.Users.Create(NewGUID("user"), token)
		Expect(err).NotTo(HaveOccurred())

		err = org.Auditors.Associate(user.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		list, err := client.Organizations.ListAuditors(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Users).To(HaveLen(1))
		Expect(list.Users[0].GUID).To(Equal(user.GUID))
	})
})

package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fetch all the spaces of an organization", func() {
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

	It("fetches the space records of all spaces associated with an organization", func() {
		space, err := client.Spaces.Create(NewGUID("space"), org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		list, err := client.Organizations.ListSpaces(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Spaces).To(HaveLen(1))
		Expect(list.Spaces[0].GUID).To(Equal(space.GUID))

		err = client.Spaces.Delete(space.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		list, err = client.Organizations.ListSpaces(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Spaces).To(HaveLen(0))
	})

	It("fetches paginated results of spaces associated with an organization", func() {
		spacenames := make(chan string, 150)
		for i := 0; i < 150; i++ {
			spacenames <- NewGUID("space")
		}

		pool := NewWorkPool(10, func() error {
			name := <-spacenames

			_, err := client.Spaces.Create(name, org.GUID, token)
			if err != nil {
				return err
			}

			return nil
		})

		for i := 0; i < 150; i++ {
			r := <-pool.Results
			Expect(r.Error).NotTo(HaveOccurred())
		}

		list, err := client.Organizations.ListSpaces(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		Expect(list.TotalResults).To(Equal(150))
		Expect(list.TotalPages).To(Equal(3))
		Expect(list.Spaces).To(HaveLen(50))

		spaces := list.Spaces
		for list.HasNextPage() {
			list, err = list.Next(token)
			Expect(err).NotTo(HaveOccurred())

			spaces = append(spaces, list.Spaces...)
		}

		Expect(spaces).To(HaveLen(150))

		cnt := 0
		for _, space := range spaces {
			err = client.Spaces.Delete(space.GUID, token)
			Expect(err).NotTo(HaveOccurred())

			cnt = cnt + 1
		}
		Expect(cnt).To(Equal(150))

		list, err = client.Organizations.ListSpaces(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		Expect(list.TotalResults).To(Equal(0))
	})
})

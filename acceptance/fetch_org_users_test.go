package acceptance

import (
	"os"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fetch all the users of an organization", func() {
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

	It("fetches the user records of all users associated with an organization", func() {
		user, err := client.Users.Create(NewGUID("user"), token)
		Expect(err).NotTo(HaveOccurred())

		err = org.Users.Associate(user.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		list, err := client.Organizations.ListUsers(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Users).To(HaveLen(1))
		Expect(list.Users[0].GUID).To(Equal(user.GUID))
	})

	It("fetches paginated results of users associated with an organization", func() {
		usernames := make(chan string, 150)
		for i := 0; i < 150; i++ {
			usernames <- NewGUID("user")
		}

		pool := NewWorkPool(10, func() error {
			name := <-usernames

			user, err := client.Users.Create(name, token)
			if err != nil {
				return err
			}

			err = org.Users.Associate(user.GUID, token)
			if err != nil {
				return err
			}

			return nil
		})

		for i := 0; i < 150; i++ {
			r := <-pool.Results
			Expect(r.Error).NotTo(HaveOccurred())
		}

		list, err := client.Organizations.ListUsers(org.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		Expect(list.TotalResults).To(Equal(150))
		Expect(list.TotalPages).To(Equal(3))
		Expect(list.Users).To(HaveLen(50))

		users, err := list.AllUsers(token)
		Expect(err).NotTo(HaveOccurred())

		Expect(users).To(HaveLen(150))
	})
})

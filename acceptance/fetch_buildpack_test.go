package acceptance

import (
	"fmt"
	"os"
	"time"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Buildpack lifecycle", func() {
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

	It("can create/get a new buildpack", func() {
		createdBuildpack, err := client.Buildpacks.Create("rainmaker-buildpack", token, nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(createdBuildpack.GUID).NotTo(BeEmpty())
		Expect(createdBuildpack.URL).To(Equal(fmt.Sprintf("/v2/buildpacks/%s", createdBuildpack.GUID)))
		Expect(createdBuildpack.CreatedAt).To(BeTemporally("~", time.Now().Truncate(time.Second).UTC(), 2*time.Second))
		Expect(createdBuildpack.UpdatedAt).To(Equal(time.Time{}))
		Expect(createdBuildpack.Name).To(Equal("rainmaker-buildpack"))
		Expect(createdBuildpack.Position).To(Equal(1))
		Expect(createdBuildpack.Enabled).To(BeTrue())
		Expect(createdBuildpack.Locked).To(BeFalse())
		Expect(createdBuildpack.Filename).To(BeEmpty())

		fetchedBuildpack, err := client.Buildpacks.Get(createdBuildpack.GUID, token)
		Expect(err).NotTo(HaveOccurred())
		Expect(fetchedBuildpack.GUID).NotTo(BeEmpty())
		Expect(fetchedBuildpack.URL).To(Equal(fmt.Sprintf("/v2/buildpacks/%s", fetchedBuildpack.GUID)))
		Expect(fetchedBuildpack.CreatedAt).To(Equal(createdBuildpack.CreatedAt))
		Expect(fetchedBuildpack.UpdatedAt).To(Equal(time.Time{}))
		Expect(fetchedBuildpack.Name).To(Equal("rainmaker-buildpack"))
		Expect(fetchedBuildpack.Position).To(Equal(1))
		Expect(fetchedBuildpack.Enabled).To(BeTrue())
		Expect(fetchedBuildpack.Locked).To(BeFalse())
		Expect(fetchedBuildpack.Filename).To(BeEmpty())
	})
})

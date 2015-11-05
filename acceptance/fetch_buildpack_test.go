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

	It("can create a new buildpack", func() {
		buildpack, err := client.Buildpacks.Create("rainmaker-buildpack", token)
		Expect(err).NotTo(HaveOccurred())
		Expect(buildpack.GUID).NotTo(BeEmpty())
		Expect(buildpack.URL).To(Equal(fmt.Sprintf("/v2/buildpacks/%s", buildpack.GUID)))
		Expect(buildpack.CreatedAt).To(BeTemporally("~", time.Now().Truncate(time.Second).UTC(), 2*time.Second))
		Expect(buildpack.UpdatedAt).To(Equal(time.Time{}))
		Expect(buildpack.Name).To(Equal("rainmaker-buildpack"))
		Expect(buildpack.Position).To(Equal(1))
		Expect(buildpack.Enabled).To(BeTrue())
		Expect(buildpack.Locked).To(BeFalse())
		Expect(buildpack.Filename).To(BeEmpty())
	})
})

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
		token     string
		client    rainmaker.Client
		buildpack rainmaker.Buildpack
	)

	BeforeEach(func() {
		token = os.Getenv("UAA_TOKEN")
		client = rainmaker.NewClient(rainmaker.Config{
			Host:          os.Getenv("CC_HOST"),
			SkipVerifySSL: true,
		})
	})

	AfterEach(func() {
		_, err := client.Buildpacks.Get(buildpack.GUID, token)
		if err == nil {
			err := client.Buildpacks.Delete(buildpack.GUID, token)
			Expect(err).NotTo(HaveOccurred())
		}
	})

	It("can create/get a new buildpack", func() {
		var err error
		buildpack, err = client.Buildpacks.Create("rainmaker-buildpack", token, nil)
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

		fetchedBuildpack, err := client.Buildpacks.Get(buildpack.GUID, token)
		Expect(err).NotTo(HaveOccurred())
		Expect(fetchedBuildpack.GUID).NotTo(BeEmpty())
		Expect(fetchedBuildpack.URL).To(Equal(fmt.Sprintf("/v2/buildpacks/%s", fetchedBuildpack.GUID)))
		Expect(fetchedBuildpack.CreatedAt).To(Equal(buildpack.CreatedAt))
		Expect(fetchedBuildpack.UpdatedAt).To(Equal(time.Time{}))
		Expect(fetchedBuildpack.Name).To(Equal("rainmaker-buildpack"))
		Expect(fetchedBuildpack.Position).To(Equal(1))
		Expect(fetchedBuildpack.Enabled).To(BeTrue())
		Expect(fetchedBuildpack.Locked).To(BeFalse())
		Expect(fetchedBuildpack.Filename).To(BeEmpty())
	})

	It("can create/delete a new buildpack", func() {
		buildpack, err := client.Buildpacks.Create("rainmaker-buildpack", token, nil)
		Expect(err).NotTo(HaveOccurred())

		err = client.Buildpacks.Delete(buildpack.GUID, token)
		Expect(err).NotTo(HaveOccurred())

		_, err = client.Buildpacks.Get(buildpack.GUID, token)
		Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
	})
})

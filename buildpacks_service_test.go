package rainmaker_test

import (
	"fmt"
	"time"

	"github.com/pivotal-cf-experimental/rainmaker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuildpacksService", func() {
	var (
		token   string
		service rainmaker.BuildpacksService
	)

	BeforeEach(func() {
		token = "token"
		service = rainmaker.NewBuildpacksService(rainmaker.Config{
			Host: fakeCloudController.URL(),
		})
	})

	Describe("Create", func() {
		It("creates a buildpack with the given name", func() {
			buildpack, err := service.Create("my-buildpack", token)
			Expect(err).NotTo(HaveOccurred())
			Expect(buildpack.GUID).NotTo(BeEmpty())
			Expect(buildpack.URL).To(Equal(fmt.Sprintf("/v2/buildpacks/%s", buildpack.GUID)))
			Expect(buildpack.CreatedAt).To(BeTemporally("~", time.Now().UTC()))
			Expect(buildpack.UpdatedAt).To(Equal(time.Time{}))
			Expect(buildpack.Name).To(Equal("my-buildpack"))
			Expect(buildpack.Position).To(Equal(1))
			Expect(buildpack.Enabled).To(BeTrue())
			Expect(buildpack.Locked).To(BeFalse())
			Expect(buildpack.Filename).To(BeEmpty())
		})
	})
})

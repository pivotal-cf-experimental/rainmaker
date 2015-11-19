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
			buildpack, err := service.Create("my-buildpack", token, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(buildpack.GUID).NotTo(BeEmpty())
			Expect(buildpack.URL).To(Equal(fmt.Sprintf("/v2/buildpacks/%s", buildpack.GUID)))
			Expect(buildpack.CreatedAt).To(Equal(time.Time{}.UTC()))
			Expect(buildpack.UpdatedAt).To(Equal(time.Time{}))
			Expect(buildpack.Name).To(Equal("my-buildpack"))
			Expect(buildpack.Position).To(Equal(0))
			Expect(buildpack.Enabled).To(BeFalse())
			Expect(buildpack.Locked).To(BeFalse())
			Expect(buildpack.Filename).To(BeEmpty())
		})

		It("creates a buildpack with optional values", func() {
			buildpack, err := service.Create("my-buildpack", token, &rainmaker.CreateBuildpackOptions{
				Position: rainmaker.IntPtr(3),
				Enabled:  rainmaker.BoolPtr(true),
				Locked:   rainmaker.BoolPtr(true),
				Filename: rainmaker.StringPtr("some-file"),
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(buildpack.GUID).NotTo(BeEmpty())
			Expect(buildpack.URL).To(Equal(fmt.Sprintf("/v2/buildpacks/%s", buildpack.GUID)))
			Expect(buildpack.CreatedAt).To(Equal(time.Time{}.UTC()))
			Expect(buildpack.UpdatedAt).To(Equal(time.Time{}))
			Expect(buildpack.Name).To(Equal("my-buildpack"))
			Expect(buildpack.Position).To(Equal(3))
			Expect(buildpack.Enabled).To(BeTrue())
			Expect(buildpack.Locked).To(BeTrue())
			Expect(buildpack.Filename).To(Equal("some-file"))
		})

		Context("when the request errors", func() {
			PIt("returns the error", func() {})
		})

		Context("when the response cannot be unmarshalled", func() {
			PIt("returns the error", func() {})
		})
	})

	Describe("Get", func() {
		var (
			bp rainmaker.Buildpack
		)

		BeforeEach(func() {
			var err error
			bp, err = service.Create("my-buildpack", token, &rainmaker.CreateBuildpackOptions{
				Position: rainmaker.IntPtr(1),
				Enabled:  rainmaker.BoolPtr(true),
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("finds the buildpack", func() {
			buildpack, err := service.Get(bp.GUID, token)
			Expect(err).NotTo(HaveOccurred())
			Expect(buildpack).To(Equal(rainmaker.Buildpack{
				GUID:      bp.GUID,
				URL:       fmt.Sprintf("/v2/buildpacks/%s", bp.GUID),
				CreatedAt: time.Time{}.UTC(),
				Name:      "my-buildpack",
				Position:  1,
				Enabled:   true,
			}))
		})

		Context("when the buildpack does not exist", func() {
			PIt("returns an error", func() {})
		})

		Context("when the request errors", func() {
			PIt("returns the error", func() {})
		})

		Context("when the response cannot be unmarshalled", func() {
			PIt("returns the error", func() {})
		})
	})
})

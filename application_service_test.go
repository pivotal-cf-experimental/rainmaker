package rainmaker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf-experimental/rainmaker"
)

var _ = Describe("ApplicationService", func() {
	var (
		config  rainmaker.Config
		token   string
		service rainmaker.ApplicationsService
		app     rainmaker.Application
	)

	BeforeEach(func() {
		token = "token"
		config = rainmaker.Config{
			Host: fakeCloudController.URL(),
		}
		client := rainmaker.NewClient(config)
		service = client.Applications
	})

	Describe("Create", func() {
		Context("when the request errors", func() {
			BeforeEach(func() {
				config.Host = ""
				service = rainmaker.NewApplicationsService(config)
			})

			It("returns the error", func() {
				_, err := service.Create(rainmaker.Application{}, token)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Get", func() {
		It("returns the app", func() {
			var appGUID string

			By("creating an app", func() {
				var err error

				app, err = service.Create(rainmaker.Application{
					Name:      "some-app",
					SpaceGUID: "some-space-guid",
				}, token)
				Expect(err).NotTo(HaveOccurred())

				appGUID = app.GUID
			})

			By("retrieving the app", func() {
				retrievedApp, err := service.Get(appGUID, token)
				Expect(err).NotTo(HaveOccurred())
				Expect(retrievedApp).To(Equal(app))
			})
		})

		Context("when the request errors", func() {
			BeforeEach(func() {
				config.Host = ""
				service = rainmaker.NewApplicationsService(config)
			})

			It("returns the error", func() {
				_, err := service.Get("whoops-guid", token)
				Expect(err).To(BeAssignableToTypeOf(rainmaker.Error{}))
			})
		})

		Context("when unmarshalling fails", func() {
			It("returns an error", func() {
				_, err := service.Get("very-bad-guid", token)
				Expect(err).To(BeAssignableToTypeOf(rainmaker.Error{}))
			})
		})
	})

	Describe("Delete", func() {
		var appGUID string

		BeforeEach(func() {
			var err error

			app, err = service.Create(rainmaker.Application{
				Name:      "some-app",
				SpaceGUID: "some-space-guid",
			}, token)
			Expect(err).NotTo(HaveOccurred())

			appGUID = app.GUID
		})

		It("deletes an application given an appGUID", func() {
			err := service.Delete(appGUID, token)
			Expect(err).NotTo(HaveOccurred())

			_, err = service.Get(appGUID, token)
			Expect(err).To(BeAssignableToTypeOf(rainmaker.NotFoundError{}))
		})

		Context("when the request errors", func() {
			BeforeEach(func() {
				config.Host = ""
				service = rainmaker.NewApplicationsService(config)
			})

			It("returns the error", func() {
				err := service.Delete("whoops-guid", token)
				Expect(err).To(BeAssignableToTypeOf(rainmaker.Error{}))
			})
		})
	})
})

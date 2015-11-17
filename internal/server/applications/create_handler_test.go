package applications_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/applications"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("createHandler", func() {
	var (
		router                 http.Handler
		recorder               *httptest.ResponseRecorder
		request                *http.Request
		applicationsCollection *domain.Applications
	)

	BeforeEach(func() {
		recorder = httptest.NewRecorder()

		requestBody, err := json.Marshal(map[string]interface{}{
			"name":       "some-app",
			"space_guid": "some-space",
			"diego":      true,
		})
		Expect(err).NotTo(HaveOccurred())

		request, err = http.NewRequest("POST", "/v2/apps", bytes.NewBuffer(requestBody))
		Expect(err).NotTo(HaveOccurred())

		guidGenerator := func(string) string { return "some-app-guid" }

		applicationsCollection = domain.NewApplications()
		router = applications.NewRouter(guidGenerator, applicationsCollection)
	})

	It("creates applications", func() {
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusCreated))

		var responseBody map[string]interface{}
		err := json.NewDecoder(recorder.Body).Decode(&responseBody)
		Expect(err).NotTo(HaveOccurred())
		Expect(responseBody).To(Equal(map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid":       "some-app-guid",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z",
			},
			"entity": map[string]interface{}{
				"name":                 "some-app",
				"space_guid":           "some-space",
				"diego":                true,
				"events_url":           "/v2/apps/some-app-guid/events",
				"routes_url":           "/v2/apps/some-app-guid/routes",
				"service_bindings_url": "/v2/apps/some-app-guid/service_bindings",
				"space_url":            "/v2/spaces/some-space",
				"stack_url":            "/v2/stacks/some-not-implemented-stack-guid",
			},
		}))
	})

	It("stores the new app in the applications collection", func() {
		_, ok := applicationsCollection.Get("some-app-guid")
		Expect(ok).To(BeFalse())

		router.ServeHTTP(recorder, request)
		app, ok := applicationsCollection.Get("some-app-guid")
		Expect(ok).To(BeTrue())
		Expect(app).To(Equal(domain.Application{
			GUID:      "some-app-guid",
			Name:      "some-app",
			SpaceGUID: "some-space",
			Diego:     true,
		}))
	})
})

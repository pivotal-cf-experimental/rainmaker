package buildpacks_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/buildpacks"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("createHandler", func() {
	var (
		router               http.Handler
		recorder             *httptest.ResponseRecorder
		request              *http.Request
		buildpacksCollection *domain.Buildpacks
	)

	BeforeEach(func() {
		guidGenerator := func(string) string { return "some-buildpack-guid" }
		buildpacksCollection = domain.NewBuildpacks()
		router = buildpacks.NewRouter(guidGenerator, buildpacksCollection)
		recorder = httptest.NewRecorder()

		requestBody, err := json.Marshal(map[string]interface{}{
			"name":     "some-buildpack",
			"position": 2,
			"enabled":  false,
			"locked":   true,
			"filename": "some-file",
		})
		Expect(err).NotTo(HaveOccurred())

		request, err = http.NewRequest("POST", "/v2/buildpacks", bytes.NewBuffer(requestBody))
		Expect(err).NotTo(HaveOccurred())
	})

	It("creates the buildpack", func() {
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusCreated))
		Expect(recorder.Body).To(MatchJSON(`{
			"metadata": {
				"guid": "some-buildpack-guid",
				"url": "/v2/buildpacks/some-buildpack-guid",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			},
			"entity": {
				"name": "some-buildpack",
				"position": 2,
				"enabled": false,
				"locked": true,
				"filename": "some-file"
			}
		}`))
	})

	It("stores the buildpack in the collection", func() {
		router.ServeHTTP(recorder, request)
		Expect(buildpacksCollection.Get("some-buildpack-guid")).To(Equal(domain.Buildpack{
			GUID:     "some-buildpack-guid",
			Name:     "some-buildpack",
			Position: 2,
			Enabled:  false,
			Locked:   true,
			Filename: "some-file",
		}))
	})

	It("handles invalid JSON request bodies", func() {
		request.Body = ioutil.NopCloser(strings.NewReader("%%%"))

		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
		Expect(recorder.Body).To(MatchJSON(`{
			"code": 1001,
			"description": "Request invalid due to parse error",
			"error_code": "CF-MessageParseError"
		}`))
	})

	It("handles optional request body fields", func() {
		requestBody, err := json.Marshal(map[string]interface{}{
			"name": "some-buildpack",
		})
		Expect(err).NotTo(HaveOccurred())

		request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusCreated))
		Expect(recorder.Body).To(MatchJSON(`{
			"metadata": {
				"guid": "some-buildpack-guid",
				"url": "/v2/buildpacks/some-buildpack-guid",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			},
			"entity": {
				"name": "some-buildpack",
				"position": 0,
				"enabled": false,
				"locked": false,
				"filename": ""
			}
		}`))

	})
})

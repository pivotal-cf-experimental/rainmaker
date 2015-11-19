package buildpacks_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/buildpacks"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("getHandler", func() {
	var (
		router   http.Handler
		recorder *httptest.ResponseRecorder
		request  *http.Request
	)

	BeforeEach(func() {
		guidGenerator := func(string) string { return "some-buildpack-guid" }

		buildpack := domain.Buildpack{
			GUID:     "some-buildpack-guid",
			Name:     "some-buildpack",
			Position: 3,
			Enabled:  false,
			Locked:   true,
			Filename: "some-file",
		}
		buildpacksCollection := domain.NewBuildpacks()
		buildpacksCollection.Add(buildpack)

		router = buildpacks.NewRouter(guidGenerator, buildpacksCollection)
		recorder = httptest.NewRecorder()

		var err error
		request, err = http.NewRequest("GET", "/v2/buildpacks/some-buildpack-guid", nil)
		Expect(err).NotTo(HaveOccurred())
	})

	It("fetches the requested buildpack", func() {
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusOK))
		Expect(recorder.Body).To(MatchJSON(`{
			"metadata": {
				"guid": "some-buildpack-guid",
				"url": "/v2/buildpacks/some-buildpack-guid",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			},
			"entity": {
			    "name": "some-buildpack",
			    "position": 3,
			    "enabled": false,
			    "locked": true,
			    "filename": "some-file"
			}
		}`))
	})
})

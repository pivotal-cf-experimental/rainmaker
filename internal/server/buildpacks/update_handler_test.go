package buildpacks_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/buildpacks"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("updateHandler", func() {
	var (
		router               http.Handler
		recorder             *httptest.ResponseRecorder
		request              *http.Request
		buildpacksCollection *domain.Buildpacks
		buildpack            domain.Buildpack
	)

	BeforeEach(func() {
		guidGenerator := func(string) string { return "some-buildpack-guid" }
		buildpacksCollection = domain.NewBuildpacks()
		router = buildpacks.NewRouter(guidGenerator, buildpacksCollection)
		buildpack = domain.NewBuildpack("some-buildpack-guid")
		buildpack.Name = "some-buildpack"
		buildpack.Position = 2
		buildpack.Enabled = false
		buildpack.Locked = false
		buildpack.Filename = "some-filename"
		buildpacksCollection.Add(buildpack)
		recorder = httptest.NewRecorder()

		body := `{
			"name" : "new_buildpack_name",
			"enabled": true,
			"locked": true,
			"position": 4,
			"filename": "some-other-filename"
		}`

		var err error
		request, err = http.NewRequest("PUT", "/v2/buildpacks/some-buildpack-guid", strings.NewReader(body))
		Expect(err).NotTo(HaveOccurred())
	})

	It("updates the buildpack", func() {
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
			    "name": "new_buildpack_name",
			    "position": 4,
			    "enabled": true,
			    "locked": true,
			    "filename": "some-other-filename"
			}
		}`))

		updatedBuildpack, ok := buildpacksCollection.Get(buildpack.GUID)
		Expect(ok).To(BeTrue())
		Expect(updatedBuildpack.Enabled).To(BeTrue())
	})

	Context("when the buildpack does not exist", func() {
		It("returns a 404", func() {
			request.URL.Path = "/v2/buildpacks/some-missing-guid"
			router.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusNotFound))
		})
	})

	Context("when the json is unproperly formatted", func() {
		It("it returns a 400", func() {
			body := `not_json`

			var err error
			request, err = http.NewRequest("PUT", "/v2/buildpacks/some-buildpack-guid", strings.NewReader(body))
			Expect(err).NotTo(HaveOccurred())

			router.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body).To(MatchJSON(`{
 				"code": 1001,
  				"description": "Request invalid due to parse error",
  				"error_code": "CF-MessageParseError"
			}`))
		})
	})
})

package buildpacks_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/pivotal-cf-experimental/rainmaker/internal/server/buildpacks"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("deleteHandler", func() {
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
		buildpacksCollection.Add(buildpack)
		recorder = httptest.NewRecorder()

		var err error
		request, err = http.NewRequest("DELETE", "/v2/buildpacks/some-buildpack-guid", nil)
		Expect(err).NotTo(HaveOccurred())
	})

	It("deletes the buildpack", func() {
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusNoContent))

		_, ok := buildpacksCollection.Get(buildpack.GUID)
		Expect(ok).To(BeFalse())
	})
})

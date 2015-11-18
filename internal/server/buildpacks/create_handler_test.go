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
		router   http.Handler
		recorder *httptest.ResponseRecorder
		request  *http.Request
	)

	BeforeEach(func() {
		guidGenerator := func(string) string { return "some-buildpack-guid" }
		router = buildpacks.NewRouter(guidGenerator, &domain.Buildpacks{})
		recorder = httptest.NewRecorder()

		requestBody, err := json.Marshal(map[string]interface{}{
			"name": "some-buildpack",
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
				"guid":       "some-buildpack-guid",
				"url":        "/v2/buildpacks/some-buildpack-guid"
			},
			"entity": {
				"name":     "some-buildpack",
				"position": 1,
				"enabled":  true,
				"locked":   false,
				"filename": ""
			}
		}`))
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
})

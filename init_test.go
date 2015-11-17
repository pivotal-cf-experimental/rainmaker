package rainmaker_test

import (
	"testing"

	"github.com/pivotal-cf-experimental/rainmaker/testserver"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var fakeCloudController *testserver.CloudController

func TestRainmakerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "rainmaker")
}

var _ = BeforeSuite(func() {
	fakeCloudController = testserver.NewCloudController()
	fakeCloudController.Start()
})

var _ = AfterSuite(func() {
	fakeCloudController.Close()
})

var _ = BeforeEach(func() {
	fakeCloudController.Reset()
})

package applications_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApplicationsSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/server/applications")
}

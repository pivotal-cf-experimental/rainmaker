package organizations_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrganizationsSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/server/organizations")
}

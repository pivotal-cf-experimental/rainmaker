package buildpacks_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuildpacksSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/server/buildpacks")
}

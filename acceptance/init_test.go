package acceptance

import (
	"fmt"
	"testing"

	"github.com/nu7hatch/gouuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAcceptanceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance")
}

func NewGUID(prefix string) string {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	if prefix != "" {
		return fmt.Sprintf("warrant-%s-%s", prefix, guid.String())
	}

	return guid.String()
}

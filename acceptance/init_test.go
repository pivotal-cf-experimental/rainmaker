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
		return fmt.Sprintf("rainmaker-%s-%s", prefix, guid.String())
	}

	return guid.String()
}

type WorkPool struct {
	Results <-chan result
}

type result struct {
	Error error
}

func NewWorkPool(count int, fn func() error) WorkPool {
	results := make(chan result)

	for i := 0; i < count; i++ {
		go func() {
			for {
				err := fn()
				results <- result{err}
			}
		}()
	}

	return WorkPool{results}
}

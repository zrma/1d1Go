package warm_up_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWarmUp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WarmUp Suite")
}

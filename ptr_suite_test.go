package ptr_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPtr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ptr Suite")
}

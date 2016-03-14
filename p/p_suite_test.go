package p_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P Suite")
}

package latte_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLatte(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Latte Suite")
}

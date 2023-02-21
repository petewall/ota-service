package internal_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

func TestLib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Internal unit test suite")
}

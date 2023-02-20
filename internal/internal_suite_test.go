package internal_test

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FailingReader struct {
	Message string
}

func (r *FailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New(r.Message)
}

func TestLib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Internal unit test suite")
}

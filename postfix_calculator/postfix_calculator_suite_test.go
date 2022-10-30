package calculator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPostfixCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postfix Calculator Suite")
}

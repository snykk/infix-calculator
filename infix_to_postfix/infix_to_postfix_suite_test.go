package infix_to_postfix_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInfixToPostfix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infix to Postfix Suite")
}

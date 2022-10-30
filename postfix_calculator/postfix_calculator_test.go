package calculator_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	postfixCalc "github.com/snykk/infix-calculator/postfix_calculator"
)

var _ = Describe("Postfix Calculate", func() {
	It("Should return error when equation is empty (\"\")", func() {
		c, err := postfixCalc.PostfixCalculate("")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("equation cannot be empty"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error when equation is only whitespace", func() {
		c, err := postfixCalc.PostfixCalculate("  ")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("equation cannot be empty"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error when operand is invalid (a + b)", func() {
		c, err := postfixCalc.PostfixCalculate("a + b")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("invalid operand"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error 1 + calculation", func() {
		c, err := postfixCalc.PostfixCalculate("2 +")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("missing required operand"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error + 1 calculation", func() {
		c, err := postfixCalc.PostfixCalculate("+ 2")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("missing required operand"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return 1 + 1 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("1 + 1")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(2.0))
	})

	It("Should return \"  1 + 1  \" calculation result correctly with whitespace", func() {
		c, err := postfixCalc.PostfixCalculate("  1 + 1  ")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(2.0))
	})

	It("Should return 1 + 2 * 3 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("1 + 2 * 3")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(7.0))
	})

	It("Should return 2 * 3 + 4 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("2 * 3 + 4")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(10.0))
	})

	It("Should return 2 + 3 / 4 * 5 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("2 + 3 / 4 * 5")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(5.75))
	})

	It("Should return 2 * 3 - 4 / 5 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("2 * 3 - 4 / 5")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(5.2))
	})

	It("Should return 2 ^ 3 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("2 ^ 3")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(8.0))
	})

	It("Should return 2 ^ 2 ^ 2 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("2 ^ 2 ^ 2")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(16.0))
	})

	It("Should return 2 ^ 2 ^ 2 ^ 2 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("2 ^ 2 ^ 2 ^ 2")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(65536.0))
	})
	It("Should return 2 ^ 3 + 2 - 10 / 5 * 29 - 29 * 2 + 3 ^ 8 calculation result correctly", func() {
		c, err := postfixCalc.PostfixCalculate("2 ^ 3 + 2 - 10 / 5 * 29 - 29 * 2 + 3 ^ 8")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(6455.0))
	})
})

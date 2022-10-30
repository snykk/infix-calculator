package infix_to_postfix_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	infToPost "github.com/snykk/infix-calculator/infix_to_postfix"
)

var _ = Describe("Infix to Postfix", func() {
	It("Should convert infix \"\" to postfix correctly", func() {
		r := infToPost.InfixToPostfix("")
		Expect(r).To(Equal(""))
	})

	It("Should convert infix 21 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("21")
		Expect(r).To(Equal("21"))
	})

	It("Should convert infix + 11 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("+ 11")
		Expect(r).To(Equal("11 +"))
	})

	It("Should convert infix 29 + to postfix correctly", func() {
		r := infToPost.InfixToPostfix("29 +")
		Expect(r).To(Equal("29 +"))
	})

	It("Should convert infix 3 + 9 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("3 + 9")
		Expect(r).To(Equal("3 9 +"))
	})

	It("Should convert infix 1 + 3 - 4 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 + 3 - 4")
		Expect(r).To(Equal("1 3 + 4 -"))
	})

	It("Should convert infix 1 * 2 + 3 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 * 2 + 3")
		Expect(r).To(Equal("1 2 * 3 +"))
	})

	It("Should convert infix 1 + 2 * 3 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 + 2 * 3")
		Expect(r).To(Equal("1 2 3 * +"))
	})

	It("Should convert infix 1 / 2 * 3 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 / 2 * 3")
		Expect(r).To(Equal("1 2 / 3 *"))
	})

	It("Should convert infix 1 + 2 * 3 / 4 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 + 2 * 3 / 4")
		Expect(r).To(Equal("1 2 3 * 4 / +"))
	})

	It("Should convert infix 1 ^ 2 ^ 3 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 ^ 2 ^ 3")
		Expect(r).To(Equal("1 2 3 ^ ^"))
	})

	It("Should convert infix 1 + 2 ^ 3 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 + 2 ^ 3")
		Expect(r).To(Equal("1 2 3 ^ +"))
	})

	It("Should convert infix 1 * 2 ^ 3 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 * 2 ^ 3")
		Expect(r).To(Equal("1 2 3 ^ *"))
	})

	It("Should convert infix 1 ^ 2 * 3 to postfix correctly", func() {
		r := infToPost.InfixToPostfix("1 ^ 2 * 3")
		Expect(r).To(Equal("1 2 ^ 3 *"))
	})
})

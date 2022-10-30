package stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snykk/infix-calculator/stack"
)

var _ = Describe("Stack", func() {
	Describe("String stack", func() {
		It("Should create a new stack correctly", func() {
			s := stack.NewStringStack()
			Expect(len(s)).To(Equal(0))
		})

		It("Should add item to end of stack with push", func() {
			s := stack.NewStringStack()
			s.Push("1")
			s.Push("+")
			s.Push("2")

			Expect(s[0]).To(Equal("1"))
			Expect(s[1]).To(Equal("+"))
			Expect(s[2]).To(Equal("2"))
		})

		It("Should remove item from end of stack and return it with pop", func() {
			s := stack.NewStringStack()
			s.Push("1")
			s.Push("20")

			item := s.Pop()
			Expect(item).To(Equal("20"))
			Expect(len(s)).To(Equal(1))
			Expect(s[0]).To(Equal("1"))
		})
	})

	Describe("Float stack", func() {
		It("Should create a new stack correctly", func() {
			s := stack.NewFloatStack()
			Expect(len(s)).To(Equal(0))
		})

		It("Should add item to end of stack with push", func() {
			s := stack.NewFloatStack()
			s.Push(1.0)
			s.Push(2.9)

			Expect(s[0]).To(Equal(1.0))
			Expect(s[1]).To(Equal(2.9))
		})

		It("Should remove item from end of stack and return it with pop", func() {
			s := stack.NewFloatStack()
			s.Push(1.0)
			s.Push(2.9)

			item := s.Pop()
			Expect(item).To(Equal(2.9))
			Expect(len(s)).To(Equal(1))
			Expect(s[0]).To(Equal(1.0))
		})
	})
})

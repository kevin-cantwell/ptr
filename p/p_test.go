package p_test

import (
	"time"

	"github.com/kevin-cantwell/ptr/p"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Addressers", func() {
	Describe("Bool", func() {
		It("Should return the pointer to a bool", func() {
			Expect(true).To(Equal(*p.Bool(true)))
		})
	})
	Describe("Byte", func() {
		It("Should return the pointer to a byte", func() {
			Expect(byte('!')).To(Equal(*p.Byte(33)))
		})
	})
	Describe("Complex64", func() {
		It("Should return the pointer to a complex64", func() {
			Expect(complex64(123)).To(Equal(*p.Complex64(123)))
		})
	})
	Describe("Complex128", func() {
		It("Should return the pointer to a complex128", func() {
			Expect(complex128(123)).To(Equal(*p.Complex128(123)))
		})
	})
	Describe("Time", func() {
		It("Should return the pointer to a time.Time", func() {
			v := time.Now()
			Expect(v).To(Equal(*p.Time(v)))
		})
	})
})

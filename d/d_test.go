package d_test

import (
	"time"

	"github.com/kevin-cantwell/ptr/d"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dereferencers", func() {
	Describe("Bool", func() {
		It("Should return the value of a *bool", func() {
			v := true
			Expect(v).To(Equal(d.Bool(&v)))
		})
	})
	Describe("Byte", func() {
		It("Should return the value of a *byte", func() {
			v := byte('!')
			Expect(v).To(Equal(d.Byte(&v)))
		})
	})
	Describe("Complex64", func() {
		It("Should return the value of a *complex64", func() {
			v := complex64(123)
			Expect(v).To(Equal(d.Complex64(&v)))
		})
	})
	Describe("Complex128", func() {
		It("Should return the value of a *complex128", func() {
			v := complex128(123)
			Expect(v).To(Equal(d.Complex128(&v)))
		})
	})
	Describe("Time", func() {
		It("Should return the value of a *time.Time", func() {
			v := time.Now()
			Expect(v).To(Equal(d.Time(&v)))
		})
	})
})

package cmd_test

import (
	. "github.com/lgug2z/tdee/cmd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Root", func() {
	var f Flags

	BeforeEach(func() {
		f = Flags{}
	})

	Describe("When run without a unit flag", func() {
		It("Should throw an error", func() {
			err := Root(f)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrNoUnitsSelected))
		})
	})

	Describe("When run without any required information", func() {
		It("Should throw an error", func() {
			f.Metric = true
			err := Root(f)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrMissingInformation))
		})
	})

	Describe("When run missing a single piece of required information", func() {
		It("Should throw an error", func() {
			f.Metric = true
			f.Height = 172
			f.Weight = 63
			f.Sex = "male"
			f.Age = 29
			err := Root(f)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrMissingInformation))
		})
	})

	Describe("When given a sex string that is not male or female", func() {
		It("Should throw an error", func() {
			f.Metric = true
			f.Height = 172
			f.Weight = 63
			f.Sex = "notmalenotfemale"
			f.Age = 29
			f.Lifestyle = 1.375
			err := Root(f)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrInvalidSex))
		})
	})

	Describe("When given an invalid lifestyle modifier", func() {
		It("Should throw an error", func() {
			f.Metric = true
			f.Height = 172
			f.Weight = 63
			f.Sex = "male"
			f.Age = 29
			f.Lifestyle = 1.4
			err := Root(f)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrInvalidLifestyleModifier))
		})
	})

	Describe("When passed the raw flag", func() {
		It("Should output the raw number without 'kcal' appended", func() {
			f.Metric = true
			f.Raw = true
			f.Height = 172
			f.Weight = 63
			f.Sex = "male"
			f.Age = 29
			f.Lifestyle = 1.375

			output := captureStdout(func (){
				Expect(Root(f)).To(Succeed())
			})

			Expect(output).ToNot(ContainSubstring("kcal"))
		})
	})
})

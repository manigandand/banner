package banner_test

import (
	"github.com/m-rec/banner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banner Test Suite", func() {
	Context("basic test init", func() {
		It("should return the err", func() {
			res, err := banner.NewBanner("invalid/timezone")
			Expect(err.Error()).To(Equal("unknown time zone invalid/timezone"))
			Ω(res).To(BeNil())
		})
		It("should return the banner object", func() {
			res, err := banner.NewBanner("Asia/Kolkata")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(res).ToNot(BeNil())
		})
	})
})

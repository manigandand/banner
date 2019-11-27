package banner_test

import (
	"github.com/m-rec/banner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banner Test Suite", func() {
	Context("basic test init", func() {
		It("should return the banner object", func() {
			res := banner.NewBanner()
			Ω(res).ToNot(BeNil())
		})
	})
})

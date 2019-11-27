package banner_test

import (
	"os"
	"testing"

	"github.com/m-rec/banner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	MainSetup()
	defer MainTearDown()
	os.Exit(m.Run())
}

var _ = Describe("Banner Test Suite", func() {
	Context("basic test init", func() {
		It("should return the banner object", func() {
			res := banner.NewBanner()
			// Ω(err).ShouldNot(HaveOccurred())
			Ω(res).ToNot(BeNil())
		})
	})
})

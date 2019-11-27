package banner_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

var _ = BeforeSuite(func() {
	Setup()
})

var _ = AfterSuite(func() {
})

func Setup() {

}

// Load mock banner data
func MainSetup() {

}

// flush all the mock banner data
func MainTearDown() {

}

package banner_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBanner(t *testing.T) {
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

func getFormatedTime(t time.Time, addH, addM int) string {
	var hour, min string
	h := t.Hour() + addH
	m := t.Minute() + addM
	hour = strconv.Itoa(h)
	min = strconv.Itoa(m)
	if h < 10 {
		hour = fmt.Sprintf("0%d", h)
	}
	if m < 10 {
		min = fmt.Sprintf("0%d", m)
	}

	return fmt.Sprintf("%s:%s", hour, min)
}

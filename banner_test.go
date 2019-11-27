package banner_test

import (
	"os"
	"testing"
	"time"

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
	var adaptor banner.Adaptor
	BeforeEach(func() {
		adaptor, _ = banner.NewBanner("Asia/Kolkata")
	})

	Context("banner plugin add method test", func() {
		It("should return name required error", func() {
			newBanner := &banner.Banner{
				Name: "",
			}
			err := adaptor.Add(newBanner)
			Expect(err.Error()).To(Equal("name required"))
		})
		It("should return url required error", func() {
			newBanner := &banner.Banner{
				Name: "A test banner",
				URL:  "",
			}
			err := adaptor.Add(newBanner)
			Expect(err.Error()).To(Equal("url required"))
		})
		It("should return start time required error", func() {
			newBanner := &banner.Banner{
				Name: "A test banner",
				URL:  "https://gopherhut.com/logo.png",
				DisplayPeriod: banner.DisplayPeriod{
					Start:    "",
					Duration: 1200,
					TimeZone: "UTC",
				},
			}
			err := adaptor.Add(newBanner)
			Expect(err.Error()).To(Equal("start time required"))
		})
		It("should return timezone required error", func() {
			newBanner := &banner.Banner{
				Name: "A test banner",
				URL:  "https://gopherhut.com/logo.png",
				DisplayPeriod: banner.DisplayPeriod{
					Start:    "01:55",
					Duration: 0,
					TimeZone: "",
				},
			}
			err := adaptor.Add(newBanner)
			Expect(err.Error()).To(Equal("timezone required"))
		})
		It("should return duration required error", func() {
			newBanner := &banner.Banner{
				Name: "A test banner",
				URL:  "https://gopherhut.com/logo.png",
				DisplayPeriod: banner.DisplayPeriod{
					Start:    "01:55",
					Duration: 0,
					TimeZone: "UTC",
				},
			}
			err := adaptor.Add(newBanner)
			Expect(err.Error()).To(Equal("duration required"))
		})
		It("should return duration required error", func() {
			newBanner := &banner.Banner{
				Name: "A test banner",
				URL:  "https://gopherhut.com/logo.png",
				DisplayPeriod: banner.DisplayPeriod{
					Start:    "01:55",
					Duration: 688,
					TimeZone: "invalid/timezone",
				},
			}
			err := adaptor.Add(newBanner)
			Expect(err.Error()).To(Equal("unknown time zone invalid/timezone"))
		})
		It("should return invalid start time", func() {
			newBanner := &banner.Banner{
				ID:     123,
				Name:   "A new banner",
				URL:    "https://media.buyee.jp/campaign/mercari191120/assets/img/title_text_en.png",
				Width:  500,
				Height: 200,
				Repeat: 10,
				DisplayPeriod: banner.DisplayPeriod{
					Start:    "01-1-bla",
					Duration: 1200,
					TimeZone: "UTC",
				},
			}
			err := adaptor.Add(newBanner)
			Expect(err.Error()).To(Equal(`parsing time "01-1-bla" as "15:04": cannot parse "-1-bla" as ":"`))
		})
		It("should return duration required error", func() {
			newBanner := &banner.Banner{
				ID:     123,
				Name:   "A new banner",
				URL:    "https://media.buyee.jp/campaign/mercari191120/assets/img/title_text_en.png",
				Width:  500,
				Height: 200,
				Repeat: 10,
				DisplayPeriod: banner.DisplayPeriod{
					Start:    "01:55",
					Duration: 1200,
					TimeZone: "UTC",
				},
			}
			err := adaptor.Add(newBanner)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Context("banner plugin get method test", func() {
		It("should return no active banners", func() {
			adaptor.Clear()
			ban, err := adaptor.Get()
			Expect(err.Error()).To(Equal("no active banners"))
			Ω(ban).To(BeNil())
		})

		It("should return banner 1", func() {
			adaptor.Clear()
			now := time.Now()
			newBanner := &banner.Banner{
				ID:     123,
				Name:   "Banner 1",
				URL:    "https://media.buyee.jp/campaign/mercari191120/assets/img/title_text_en.png",
				Width:  500,
				Height: 200,
				Repeat: 10,
				DisplayPeriod: banner.DisplayPeriod{
					Start:    getFormatedTime(now, 0, 0),
					Duration: 1200,
					TimeZone: "UTC",
				},
			}
			err := adaptor.Add(newBanner)
			Ω(err).ShouldNot(HaveOccurred())

			ban, err := adaptor.Get()
			Ω(err).ShouldNot(HaveOccurred())
			Expect(ban.ID).To(Equal(uint(123)))
			Expect(ban.Name).To(Equal("Banner 1"))
			Expect(ban.IsActive).To(BeTrue())
			Expect(ban.Start).To(Equal(getFormatedTime(now, 0, 0)))
			Expect(ban.Duration).To(Equal(time.Duration(1200)))
			Expect(ban.TimeZone).To(Equal("UTC"))
		})

		It("should return no active banners 2", func() {
			adaptor.Clear()
			now := time.Now()
			startTime := getFormatedTime(now, 0, 0)
			startTime2 := getFormatedTime(now, 0, 0)

			newBanner := &banner.Banner{
				ID:     123,
				Name:   "Banner 1",
				URL:    "https://media.buyee.jp/campaign/mercari191120/assets/img/title_text_en.png",
				Width:  500,
				Height: 200,
				Repeat: 10,
				DisplayPeriod: banner.DisplayPeriod{
					Start:    startTime,
					Duration: 1200,
					TimeZone: "UTC",
				},
			}
			err := adaptor.Add(newBanner)
			Ω(err).ShouldNot(HaveOccurred())

			newBanner2 := &banner.Banner{
				ID:     456,
				Name:   "Banner 2",
				URL:    "https://media.buyee.jp/campaign/mercari191120/assets/img/title_text_en.png",
				Width:  500,
				Height: 200,
				Repeat: 10,
				DisplayPeriod: banner.DisplayPeriod{
					Start:    startTime2,
					Duration: 1100,
					TimeZone: "UTC",
				},
			}
			err = adaptor.Add(newBanner2)
			Ω(err).ShouldNot(HaveOccurred())

			ban, err := adaptor.Get()
			Ω(err).ShouldNot(HaveOccurred())
			Expect(ban.ID).To(Equal(uint(456)))
			Expect(ban.Name).To(Equal("Banner 2"))
			Expect(ban.IsActive).To(BeTrue())
			Expect(ban.Start).To(Equal(startTime2))
			Expect(ban.Duration).To(Equal(time.Duration(1100)))
			Expect(ban.TimeZone).To(Equal("UTC"))
		})
	})
})

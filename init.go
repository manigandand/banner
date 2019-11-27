package banner

// Adaptor interface expose Get banner method
type Adaptor interface {
	Add(banner *Banner) error
	Get() (*Banner, error)
	Clear()
}

// NewBanner return the banner interface
func NewBanner(timeZone string) (Adaptor, error) {
	loc, err := loadLocation(timeZone)
	if err != nil {
		return nil, err
	}

	banners, err := loadBannersFromStub()
	if err != nil {
		return nil, err
	}

	adaptor := &Plugin{
		clientTimeZone: timeZone,
		timeZone:       loc,
		banners:        banners,
		bannersMap:     make(map[string]*Banner),
	}
	adaptor.toMap()

	return adaptor, nil
}

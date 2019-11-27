package banner

// Adaptor interface expose Get banner method
type Adaptor interface {
	Add()
	Get()
}

// NewBanner return the banner interface
func NewBanner() Adaptor {
	return &banner{}
}

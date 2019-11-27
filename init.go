package banner

// Banner interface expose Get banner method
type Banner interface {
	Add()
	Get()
}

// NewBanner return the banner interface
func NewBanner() Banner {
	return &banner{}
}

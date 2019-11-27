package banner

import (
	"sync"
	"time"
)

// Banner struct act as local store, which stores all the banners
type Banner struct {
	mu             sync.RWMutex
	clientTimeZone string
	banners        map[string]*banner
}

// banner struct holds all the properties of the
type banner struct {
	bannerMu sync.RWMutex
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	// Not Implemented: we can restrict how many time the banner can visible
	Repeat       int `json:"repeat"`
	displayCount int
	displayPeriod
}

// displayPeriod holds all the information of banner display time
type displayPeriod struct {
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	Duration  time.Duration `json:"duration"`
	TimeZone  string        `json:"time_zone"`
}

// Add method implements the banner Adaptor interface
// adds a new banner into the available list
func (b *banner) Add() {

}

// Get method implements the banner Adaptor interface
// returns the first active banner
func (b *banner) Get() {

}

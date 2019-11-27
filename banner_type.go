package banner

import (
	"sync"
	"time"
)

// Banner struct holds all the properties of the
type Banner struct {
	bannerMu     sync.RWMutex
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Repeat       int    `json:"repeat"`
	IsActive     bool   `json:"is_active"`
	displayCount int
	DisplayPeriod
}

// DisplayPeriod holds all the information of banner display time
type DisplayPeriod struct {
	// represents the time to start the banner in 24 hours format. Ex: "15:04"
	Start    string        `json:"start"`
	Duration time.Duration `json:"duration"` // in sec, indicates how long this
	// will be active from the start time
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	StartTimeSec int64     `json:"start_time_sec"`
	EndTimeSec   int64     `json:"end_time_sec"`
	TimeZone     string    `json:"time_zone"`
}

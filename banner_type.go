package banner

import (
	"errors"
	"strings"
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
	TimeZone string `json:"time_zone"`

	startTime    time.Time
	endTime      time.Time
	startTimeSec int64
	endTimeSec   int64
}

func (b *Banner) ok() error {
	switch {
	case strings.TrimSpace(b.Name) == "":
		return errors.New("name required")
	case strings.TrimSpace(b.URL) == "":
		return errors.New("url required")
	case strings.TrimSpace(b.DisplayPeriod.Start) == "":
		return errors.New("start time required")
	case strings.TrimSpace(b.DisplayPeriod.TimeZone) == "":
		return errors.New("timezone required")
	case b.DisplayPeriod.Duration == 0:
		return errors.New("duration required")
	}

	if _, err := loadLocation(b.DisplayPeriod.TimeZone); err != nil {
		return err
	}

	return nil
}

// loadBannerTimes loads the banner times from the given start time in UTC
func (b *Banner) loadBannerTimes() error {
	t, err := time.Parse("15:04", b.Start)
	if err != nil {
		return err
	}
	startTime := t.In(utcLoc)
	endTime := startTime.Add(b.Duration * time.Second)

	b.startTime = startTime
	b.startTimeSec = timeToSec(startTime)
	b.endTime = endTime
	b.endTimeSec = timeToSec(endTime)

	return nil
}

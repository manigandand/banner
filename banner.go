package banner

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type bannerSortByEndTime []*Banner

func (a bannerSortByEndTime) Len() int           { return len(a) }
func (a bannerSortByEndTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bannerSortByEndTime) Less(i, j int) bool { return a[i].EndTimeSec < a[j].EndTimeSec }

// Plugin struct act as local store, which stores all the banners
type Plugin struct {
	mu             sync.RWMutex
	clientTimeZone string
	timeZone       *time.Location
	banners        []*Banner
	bannersMap     map[string]*Banner
}

// Add method implements the banner Adaptor interface
// adds a new banner into the available list
func (b *Plugin) Add() {

}

// Get method implements the banner Adaptor interface
// returns the first active banner
func (b *Plugin) Get() (*Banner, error) {
	currentTime := b.getClientTimeSec()

	var activeBanners bannerSortByEndTime
	for _, ban := range b.banners {
		if ban.StartTimeSec <= currentTime && currentTime <= ban.EndTimeSec {
			activeBanners = append(activeBanners, ban)
		}
	}
	// sort and return the banner which has earilest endtime
	sort.Sort(activeBanners)
	if len(activeBanners) == 0 {
		return nil, errors.New("no active banners")
	}

	return activeBanners[0], nil
}

// getClientTimeSec returns the current time in sec by client timezone
func (b *Plugin) getClientTimeSec() int64 {
	t := time.Now()
	if b.timeZone != nil {
		t = t.In(b.timeZone)
	}
	return int64((t.Hour() * 3600) + (t.Minute() * 60))
}

func (b *Plugin) toMap() {
	if len(b.banners) == 0 || b.banners == nil {
		b.bannersMap = make(map[string]*Banner)
		return
	}

	for _, ban := range b.banners {
		b.mu.Lock()
		b.bannersMap[ban.Name] = ban
		b.mu.Unlock()
	}
}

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
func (a bannerSortByEndTime) Less(i, j int) bool { return a[i].endTimeSec < a[j].endTimeSec }

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
func (b *Plugin) Add(banner *Banner) error {
	if err := banner.ok(); err != nil {
		return err
	}

	if err := banner.loadBannerTimes(); err != nil {
		return err
	}
	banner.IsActive = true
	b.banners = append(b.banners, banner)
	b.mu.Lock()
	b.bannersMap[banner.Name] = banner
	b.mu.Unlock()

	return nil
}

// Get method implements the banner Adaptor interface
// returns the first active banner
func (b *Plugin) Get() (*Banner, error) {
	currentTime := b.getClientTimeSec()
	var activeBanners bannerSortByEndTime
	isInternalIP := checkIsInternalIP()

	for _, ban := range b.banners {
		if isInternalIP {
			if currentTime <= ban.endTimeSec {
				activeBanners = append(activeBanners, ban)
			}
			continue
		}
		if ban.startTimeSec <= currentTime && currentTime <= ban.endTimeSec {
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

// Clear resets the data
func (b *Plugin) Clear() {
	b.banners = []*Banner{}
	b.bannersMap = make(map[string]*Banner)
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

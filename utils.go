package banner

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

var utcLoc *time.Location

func init() {
	utcLoc, _ = loadLocation("UTC")
}

// loadLocation validates the given timezone string and returns the location
// object for valid timezone
func loadLocation(timeZone string) (*time.Location, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}
	return loc, nil
}

// loadBannersFromStub loads the banner data from the stub file
func loadBannersFromStub() ([]*Banner, error) {
	var banners []*Banner
	bytestream, err := ioutil.ReadFile("stub/banners.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytestream, &banners); err != nil {
		return nil, err
	}

	for _, b := range banners {
		b.IsActive = true
		if err := b.loadBannerTimes(); err != nil {
			return nil, err
		}
	}

	return banners, nil
}

// FIXME: can be improved, handle the zero case
func timeToSec(t time.Time) int64 {
	return int64((t.Hour() * 3600) + (t.Minute() * 60))
}

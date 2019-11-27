package banner

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

var (
	utcLoc                                                  *time.Location
	private16BitBlock, private20BitBlock, private24BitBlock *net.IPNet
)

func init() {
	utcLoc, _ = loadLocation("UTC")
	_, private24BitBlock, _ = net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ = net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ = net.ParseCIDR("192.168.0.0/16")
}

func checkIsInternalIP() bool {
	ip := getExternalIP()
	if strings.TrimSpace(ip) == "" {
		return false
	}

	return isInternalIP(ip)
}

func getExternalIP() string {
	ipadd := ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ipadd
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipadd = ipnet.IP.String()
			}
		}
	}

	return ipadd
}

// isInternalIP checks if the given ip is internal/private or not
func isInternalIP(ip string) bool {
	private := false
	IP := net.ParseIP(ip)
	if IP == nil {
		return private
	}
	private = private24BitBlock.Contains(IP) ||
		private20BitBlock.Contains(IP) || private16BitBlock.Contains(IP)

	return private
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

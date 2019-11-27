package main

import (
	"fmt"
	"log"
	"time"

	"github.com/m-rec/banner"
)

func main() {
	t := time.Now()
	date := fmt.Sprintf("%d", t.Hour())
	fmt.Println(date)
	return
	b, err := banner.NewBanner("Asia/Kolkata")
	if err != nil {
		log.Fatal(err)
	}
	// add a new banner item
	newBan := &banner.Banner{
		ID:     123,
		Name:   "A new banner",
		URL:    "https://media.buyee.jp/campaign/mercari191120/assets/img/title_text_en.png",
		Width:  500,
		Height: 200,
		Repeat: 10,
		DisplayPeriod: banner.DisplayPeriod{
			Start:    "01:55",
			Duration: 1200,
			TimeZone: "UTC",
		},
	}
	if err := b.Add(newBan); err != nil {
		log.Fatal(err)
	}

	// get the first active banner
	ban, err := b.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Active banner: ", ban.Name, " URL: ", ban.URL, ban.Start)
}

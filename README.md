# mercari banner library

See [SKILL_TEST.en.md](/SKILL_TEST.en.md) for the task details.

## how to test

```shell
go test -cover ./...
```

## how to use

```shell
go get github.com/m-rec/45b44fc1b115311e1be0b63b0f6ae90751bf6f74
```

```golang
package main

import (
	"fmt"
    "log"
    // mercari "github.com/m-rec/45b44fc1b115311e1be0b63b0f6ae90751bf6f74"
    mercari "github.com/m-rec/banner"
)

func main() {
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

```

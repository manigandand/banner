# mercari banner library

See [SKILL_TEST.en.md](/SKILL_TEST.en.md) for the task details.

## how to test

```shell
go test -cover ./...
```

or

```shell
./test.sh
```

```shell
Running Suite: Api Suite
========================
Random Seed: 1574892063
Will run 13 of 13 specs

•••••••••••••
Ran 13 of 13 Specs in 0.005 seconds
SUCCESS! -- 13 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
coverage: 92.5% of statements
ok  	github.com/manigandand/banner	0.129s
```

## how to use

```shell
go get github.com/manigandand/45b44fc1b115311e1be0b63b0f6ae90751bf6f74
```

```golang
package main

import (
	"fmt"
	"log"

    mercari "github.com/manigandand/banner"
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

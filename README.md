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
	b.Add()

	// get the first active banner
	ban, err := b.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Active banner: ", ban.Name, " URL: ", ban.URL)
}

```

# mercari banner library

# how to test

# how to use

```golang
package main

import (
    "log"
    // mercari "github.com/m-rec/45b44fc1b115311e1be0b63b0f6ae90751bf6f74"
    mercari "github.com/m-rec/banner"
)

func main() {
    banner := mercari.NewBanner()
    // add a new banner item
    b.Add()

    // get the first active banner
    ban, err := banner.Get()
    if err != nil {
        log.Fatal(err)
    }

    log.Println(ban)
}

```

package main

import (
	"fmt"
	"log"

	"github.com/m-rec/banner"
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

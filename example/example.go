package main

import (
	"fmt"

	"github.com/m-rec/banner"
)

func main() {
	fmt.Println("init")
	b := banner.NewBanner()
	b.Add()
	b.Get()
	fmt.Println("exit")
}

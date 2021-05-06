package main

import (
	"log"
	"os"

	"github.com/abibby/mangadexv5"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	c := mangadexv5.NewClient()
	err := c.Login(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(c.UserFlolowsManga(1, 0))
}

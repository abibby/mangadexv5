package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/abibby/mangadexv5"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	c := mangadexv5.NewClient()

	token, err := ioutil.ReadFile("./token.txt")

	if err != nil {
		err := c.Login(os.Args[1], os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile("./token.txt", []byte(c.Token()), 0755)
	} else {
		c.SetToken(string(token))
	}
	// spew.Dump(c.UserFlolowsManga(1, 0))
	spew.Dump(c.ChapterList(nil))
}

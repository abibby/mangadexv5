package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/abibby/mangadexv5"
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

	chapters, _, err := c.UserFeedChapters(&mangadexv5.UserFeedChaptersRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range chapters {
		fmt.Printf("%s | %s V%d #%s\n", "", c.Title, c.Volume, c.Chapter)
	}

	// manga, _, err := c.UserFlolowsManga(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, m := range manga {
	// 	chapters, _, err := c.ChapterList(&mangadexv5.ChapterListRequest{
	// 		MangaID: m.ID,
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for _, c := range chapters {
	// 		fmt.Printf("%s | %s V%d #%s\n", m.Title, c.Title, c.Volume, c.Chapter)
	// 	}
	// }
}

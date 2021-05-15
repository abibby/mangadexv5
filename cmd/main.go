package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abibby/mangadexv5"
)

func main() {
	c := mangadexv5.NewClient()

	err := c.Authenticate(os.Args[1], os.Args[2], "./token.json")
	if err != nil {
		log.Fatal(err)
	}

	chapters, _, err := c.UserFeedChapters(&mangadexv5.UserFeedChaptersRequest{
		Limit:   50,
		Locales: []string{"en"},
		// CreatedAtSince: time.Now().Add(-24 * 60 * time.Hour),
	})
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = c.AttachManga(chapters)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, c := range chapters {
		fmt.Printf("%s %s | %s V%d #%s, %s\n", c.Manga().ID, c.Manga().Title, c.Title, c.Volume, c.Chapter, c.PublishAt)
	}

}

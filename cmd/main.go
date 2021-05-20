package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abibby/mangadexv5"
)

func main() {
	c := mangadexv5.NewClient()

	err := c.Authenticate(os.Args[1], os.Args[2], "./token.json")
	if err != nil {
		log.Fatal(err)
	}
	var response *mangadexv5.PaginatedResponse

	request := &mangadexv5.UserFeedChaptersRequest{
		Limit:   100,
		Locales: []string{"en"},
		CreatedAtSince: time.Now().
			Add(-24 * 60 * time.Hour).
			Format("2006-01-02T15:04:05"),
	}

	for mangadexv5.EachPage(request, response) {
		var chapters []*mangadexv5.Chapter
		chapters, response, err = c.UserFeedChapters(request)
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

}

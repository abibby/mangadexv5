package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abibby/mangadexv5"
)

func main() {
	c := mangadexv5.NewClient("./token.json")

	// _, err := ioutil.ReadFile("./token.txt")

	err := c.Authenticate(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// manga, _, err := c.UserFlolowsManga(&mangadexv5.UserFlolowsMangaRequest{
	// 	Limit: 100,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	chapters, _, err := c.UserFeedChapters(&mangadexv5.UserFeedChaptersRequest{
		Limit:          50,
		OrderCreatedAt: "asc",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.AttachManga(chapters)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range chapters {
		fmt.Printf("%s | %s V%d #%s, %s\n", c.Manga().Title, c.Title, c.Volume, c.Chapter, c.CreatedAt)
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

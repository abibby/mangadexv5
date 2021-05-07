package mangadexv5

import (
	"fmt"

	"github.com/pkg/errors"
)

type Manga struct {
	Model

	Title                  map[string]string   `json:"title"`
	AltTitles              []map[string]string `json:"altTitles"`
	Description            map[string]string   `json:"description"`
	IsLocked               bool                `json:"isLocked"`
	Links                  []string            `json:"links"`
	OriginalLanguage       string              `json:"originalLanguage"`
	LastVolume             string              `json:"lastVolume"`
	LastChapter            string              `json:"lastChapter"`
	PublicationDemographic string              `json:"publicationDemographic"`
	Status                 string              `json:"status"`
	Year                   int                 `json:"year"`
	ContentRating          string              `json:"contentRating"`
	Tags                   []*Tag              `json:"tags"`
	Version                int                 `json:"version"`
	CreatedAt              string              `json:"createdAt"`
	UpdatedAt              string              `json:"updatedAt"`
}

type Tag struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes *TagAttributes `json:"attributes"`
}

type TagAttributes struct {
	Name map[string]string `json:"name"`
}

// UserFlolowsManga
//
// API Link https://api.mangadex.org/docs.html#operation/get-user-follows-manga
func (c *Client) UserFlolowsManga(limit, offset int) ([]*Manga, *PaginatedResponse, error) {
	resp := &PaginatedResponse{}
	err := c.get(fmt.Sprintf("/user/follows/manga?limit=%d&offset=%d", limit, offset), resp)
	if err != nil {
		return nil, nil, err
	}

	manga := []*Manga{}
	err = resp.loadResults(&manga)
	if err != nil {
		return nil, resp, errors.Wrap(err, "failed to load manga from response")
	}

	return manga, resp, nil

}

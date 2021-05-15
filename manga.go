package mangadexv5

import (
	"github.com/pkg/errors"
)

type Manga struct {
	Model

	Title                  LangMap     `json:"title"`
	AltTitles              []LangMap   `json:"altTitles"`
	Description            LangMap     `json:"description"`
	IsLocked               bool        `json:"isLocked"`
	Links                  *MangaLinks `json:"links"`
	OriginalLanguage       string      `json:"originalLanguage"`
	LastVolume             string      `json:"lastVolume"`
	LastChapter            string      `json:"lastChapter"`
	PublicationDemographic string      `json:"publicationDemographic"`
	Status                 string      `json:"status"`
	Year                   int         `json:"year"`
	ContentRating          string      `json:"contentRating"`
	Tags                   []*Tag      `json:"tags"`
	Version                int         `json:"version"`
	CreatedAt              string      `json:"createdAt"`
	UpdatedAt              string      `json:"updatedAt"`
}

// https://api.mangadex.org/docs.html#section/Static-data/Manga-links-data
type MangaLinks struct {
	AnilistID string `json:"al"`
}

type Tag struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes *TagAttributes `json:"attributes"`
}

type TagAttributes struct {
	Name map[string]string `json:"name"`
}

type MangaListRequest struct {
	Limit  int      `qstring:"limit,omitempty"`
	Offset int      `qstring:"offset,omitempty"`
	Title  string   `qstring:"title,omitempty"`
	IDs    []string `qstring:"ids[],omitempty"`
}

// MangaList
//
// API Link https://api.mangadex.org/docs.html#operation/get-search-manga
func (c *Client) MangaList(request *MangaListRequest) ([]*Manga, *PaginatedResponse, error) {
	resp := &PaginatedResponse{}
	err := c.get("/manga", request, resp)
	if err != nil {
		return nil, nil, errors.Wrap(err, "request failed")
	}

	manga := []*Manga{}
	err = resp.loadResults(&manga)
	if err != nil {
		return nil, resp, errors.Wrap(err, "failed to load manga from response")
	}

	return manga, resp, nil

}

type UserFlolowsMangaRequest struct {
	Limit  int `qstring:"limit,omitempty"`
	Offset int `qstring:"offset,omitempty"`
}

// UserFlolowsManga
//
// API Link https://api.mangadex.org/docs.html#operation/get-user-follows-manga
func (c *Client) UserFlolowsManga(request *UserFlolowsMangaRequest) ([]*Manga, *PaginatedResponse, error) {
	resp := &PaginatedResponse{}
	err := c.get("/user/follows/manga", request, resp)
	if err != nil {
		return nil, nil, errors.Wrap(err, "request failed")
	}

	manga := []*Manga{}
	err = resp.loadResults(&manga)
	if err != nil {
		return nil, resp, errors.Wrap(err, "failed to load manga from response")
	}

	return manga, resp, nil

}

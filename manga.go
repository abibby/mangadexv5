package mangadexv5

import "fmt"

type Tag struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes *TagAttributes `json:"attributes"`
}

type TagAttributes struct {
	Name map[string]string `json:"name"`
}

type UserFlolowsMangaResponse struct {
	Results []*UserFlolowsMangaResult `json:"results"`
	Limit   int                       `json:"limit"`
	Offset  int                       `json:"offset"`
	Total   int                       `json:"total"`
}

type UserFlolowsMangaResult struct {
	Result string                `json:"result"`
	Data   *UserFlolowsMangaData `json:"data"`
}

type UserFlolowsMangaData struct {
	ID         string                      `json:"id"`
	Type       string                      `json:"type"`
	Attributes *UserFlolowsMangaAttributes `json:"attributes"`
}
type UserFlolowsMangaAttributes struct {
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

// UserFlolowsManga
//
// API Link https://api.mangadex.org/docs.html#operation/get-user-follows-manga
func (c *Client) UserFlolowsManga(limit, offset int) (*UserFlolowsMangaResponse, error) {
	result := &UserFlolowsMangaResponse{}
	err := c.get(fmt.Sprintf("/user/follows/manga?limit=%d&offset=%d", limit, offset), result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

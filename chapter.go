package mangadexv5

import (
	"fmt"
	"path"
	"time"

	"github.com/abibby/nulls"
	"github.com/pkg/errors"
)

type Chapter struct {
	Model

	Title              string        `json:"title"`              // <= 255 characters
	Volume             *nulls.String `json:"volume"`             // or null
	Chapter            *nulls.String `json:"chapter"`            // or null <= 8 characters
	Pages              int           `json:"pages"`              // Count of readable images for this chapter
	TranslatedLanguage string        `json:"translatedLanguage"` // ^[a-z]{2}(-[a-z]{2})?$
	Uploader           string        `json:"uploader"`           // <uuid>
	ExternalURL        *nulls.String `json:"externalUrl"`        // or null <= 512 characters ^https?:// Denotes a chapter that links to an external source.
	Version            int           `json:"version"`            // >= 1
	CreatedAt          time.Time     `json:"createdAt"`
	UpdatedAt          time.Time     `json:"updatedAt"`
	PublishAt          time.Time     `json:"publishAt"`
	ReadableAt         time.Time     `json:"readableAt"`

	manga *Manga
}

func (c *Chapter) Manga() *Manga {
	if c.manga == nil {
		return &Manga{}
	}
	return c.manga
}

func (c *Chapter) PageURL(atHomeServer *AtHomeServerResponse, page int) string {
	return atHomeServer.BaseURL + "/" + path.Join("data", atHomeServer.Chapter.Hash, atHomeServer.Chapter.Data[page])
}

func (c *Chapter) PageURLDataSaver(atHomeServer *AtHomeServerResponse, page int) string {
	return atHomeServer.BaseURL + "/" + path.Join("data-saver", atHomeServer.Chapter.Hash, atHomeServer.Chapter.DataSaver[page])
}

type ChapterListRequest struct {
	Limit              int       `qstring:"limit,omitempty"`
	Offset             int       `qstring:"offset,omitempty"`
	IDs                []string  `qstring:"ids[],omitempty"`
	Title              string    `qstring:"title,omitempty"`
	GroupIDs           []string  `qstring:"groups[],omitempty"`
	UploaderID         string    `qstring:"uploader,omitempty"`
	MangaID            string    `qstring:"manga,omitempty"`
	VolumeIDs          string    `qstring:"volume[],omitempty"`
	ChapterIDs         string    `qstring:"chapter[],omitempty"`
	TranslatedLanguage []string  `qstring:"translatedLanguage[],omitempty"`
	CreatedAtSince     time.Time `qstring:"createdAtSince,omitempty"`
	UpdatedAtSince     time.Time `qstring:"updatedAtSince,omitempty"`
	PublishAtSince     time.Time `qstring:"publishAtSince,omitempty"`
	OrderCreatedAt     Order     `qstring:"order[createdAt],omitempty"`
	OrderUpdatedAt     Order     `qstring:"order[updatedAt],omitempty"`
	OrderPublishAt     Order     `qstring:"order[publishAt],omitempty"`
	OrderVolume        Order     `qstring:"order[volume],omitempty"`
	OrderChapter       Order     `qstring:"order[chapter],omitempty"`
}

func (r *ChapterListRequest) SetOffset(offset int) {
	r.Offset = offset
}

func (r *ChapterListRequest) SetLimit(limit int) {
	r.Limit = limit
}

var _ PaginatedRequest = &ChapterListRequest{}

// ChapterList
//
// API Link https://api.mangadex.org/docs.html#operation/get-chapter
func (c *Client) ChapterList(request *ChapterListRequest) ([]*Chapter, *PaginatedResponse, error) {
	resp := &PaginatedResponse{}
	err := c.get("/chapter", request, resp)
	if err != nil {
		return nil, nil, err
	}

	chapters := []*Chapter{}

	err = resp.loadResults(&chapters)
	if err != nil {
		return nil, resp, errors.Wrap(err, "failed to load chapters from response")
	}

	return chapters, resp, err

}

type UserFeedChaptersRequest struct {
	Limit              int       `qstring:"limit,omitempty"`
	Offset             int       `qstring:"offset,omitempty"`
	TranslatedLanguage []string  `qstring:"translatedLanguage[],omitempty"`
	CreatedAtSince     time.Time `qstring:"createdAtSince,omitempty"`
	UpdatedAtSince     time.Time `qstring:"updatedAtSince,omitempty"`
	PublishAtSince     string    `qstring:"publishAtSince,omitempty"`
	OrderCreatedAt     Order     `qstring:"order[createdAt],omitempty"`
	OrderUpdatedAt     Order     `qstring:"order[updatedAt],omitempty"`
	OrderPublishAt     Order     `qstring:"order[publishAt],omitempty"`
	OrderVolume        Order     `qstring:"order[volume],omitempty"`
	OrderChapter       Order     `qstring:"order[chapter],omitempty"`
}

func (r *UserFeedChaptersRequest) SetOffset(offset int) {
	r.Offset = offset
}

func (r *UserFeedChaptersRequest) SetLimit(limit int) {
	r.Limit = limit
}

var _ PaginatedRequest = &UserFeedChaptersRequest{}

// UserFeedChapters
//
// API Link https://api.mangadex.org/docs.html#operation/get-user-follows-manga-feed
func (c *Client) UserFeedChapters(request *UserFeedChaptersRequest) ([]*Chapter, *PaginatedResponse, error) {
	resp := &PaginatedResponse{}
	err := c.get("/user/follows/manga/feed", request, resp)
	if err != nil {
		return nil, nil, errors.Wrap(err, "request failed")
	}

	chapters := []*Chapter{}
	err = resp.loadResults(&chapters)
	if err != nil {
		fmt.Printf("%s\n", resp.Data[0].Attributes)
		return nil, resp, errors.Wrap(err, "failed to load chapters from response")
	}

	return chapters, resp, nil

}

func (c *Client) AttachManga(chapters []*Chapter) error {
	manga := map[string]*Manga{}

	for _, chapter := range chapters {
		manga[chapter.Relationships.Get("manga")] = nil
	}

	if len(manga) == 0 {
		return nil
	}

	mangaIDs := []string{}
	for id := range manga {
		mangaIDs = append(mangaIDs, id)
	}

	var mangaList []*Manga
	var response *PaginatedResponse
	var err error

	request := &MangaListRequest{
		Limit: 100,
		IDs:   mangaIDs,
	}

	for EachPage(request, response) {
		mangaList, response, err = c.MangaList(request)
		if err != nil {
			return err
		}

		for _, m := range mangaList {
			manga[m.ID] = m
		}

		for _, chapter := range chapters {
			chapter.manga = manga[chapter.Relationships.Get("manga")]
		}
	}
	return nil
}

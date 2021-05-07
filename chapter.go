package mangadexv5

import (
	"github.com/pkg/errors"
)

type Chapter struct {
	Model

	Title              string   `json:"title"`
	Volume             int      `json:"volume"`
	Chapter            string   `json:"chapter"`
	TranslatedLanguage string   `json:"translatedLanguage"`
	Hash               string   `json:"hash"`
	Data               []string `json:"data"`
	DataSaver          []string `json:"dataSaver"`
	Uploader           string   `json:"uploader"`
	Version            int      `json:"version"`
	CreatedAt          string   `json:"createdAt"`
	UpdatedAt          string   `json:"updatedAt"`
	PublishAt          string   `json:"publishAt"`
}

type ChapterListRequest struct {
	Limit              int      `qstring:"limit,omitempty"`
	Offset             int      `qstring:"offset,omitempty"`
	IDs                []string `qstring:"ids,omitempty"`
	Title              string   `qstring:"title,omitempty"`
	GroupIDs           []string `qstring:"groups,omitempty"`
	UploaderID         string   `qstring:"uploader,omitempty"`
	MangaID            string   `qstring:"manga,omitempty"`
	Volume             string   `qstring:"volume,omitempty"`
	Chapter            string   `qstring:"chapter,omitempty"`
	TranslatedLanguage string   `qstring:"translatedLanguage,omitempty"`
	CreatedAtSince     string   `qstring:"createdAtSince,omitempty"`
	UpdatedAtSince     string   `qstring:"updatedAtSince,omitempty"`
	PublishAtSince     string   `qstring:"publishAtSince,omitempty"`
}

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

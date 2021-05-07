package mangadexv5

import (
	"github.com/dyninc/qstring"
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
	Limit              int      `json:"limit"`
	Offset             int      `json:"offset"`
	IDs                []string `json:"ids"`
	Title              string   `json:"title"`
	GroupIDs           []string `json:"groups"`
	UploaderID         string   `json:"uploader"`
	MangaID            string   `json:"manga"`
	Volume             string   `json:"volume"`
	Chapter            string   `json:"chapter"`
	TranslatedLanguage string   `json:"translatedLanguage"`
	CreatedAtSince     string   `json:"createdAtSince"`
	UpdatedAtSince     string   `json:"updatedAtSince"`
	PublishAtSince     string   `json:"publishAtSince"`
}

// ChapterList
//
// API Link https://api.mangadex.org/docs.html#operation/get-chapter
func (c *Client) ChapterList(request *ChapterListRequest) ([]*Chapter, *PaginatedResponse, error) {
	q, err := qstring.MarshalString(request)
	if err != nil {
		return nil, nil, err
	}

	resp := &PaginatedResponse{}
	err = c.get("/chapter?"+q, resp)
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

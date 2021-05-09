package mangadexv5

import (
	"path"

	"github.com/pkg/errors"
)

type AtHomeServerResponse struct {
	BaseURL string `json:"baseUrl"`
}

func (c *Client) AtHomeServer(chapterID string) (*AtHomeServerResponse, error) {
	response := &AtHomeServerResponse{}
	err := c.get(path.Join("/at-home/server", chapterID), nil, response)
	if err != nil {
		return nil, errors.Wrap(err, "could not find at home server")
	}
	return response, nil
}

package mangadexv5

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

type PaginatedResponse struct {
	Results []*Result `json:"results"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
	Total   int       `json:"total"`
}

type Result struct {
	Result        string          `json:"result"`
	Data          *ResultData     `json:"data"`
	Relationships []*Relationship `json:"relationships"`
}

type ResultData struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes json.RawMessage `json:"attributes"`
}

type Relationship struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type SetIDer interface {
	SetID(id string)
}

type Model struct {
	ID string
}

func (m *Model) SetID(id string) {
	m.ID = id
}

func (r *PaginatedResponse) loadResults(v interface{}) error {
	rv := reflect.ValueOf(v).Elem()

	for _, result := range r.Results {
		element := reflect.New(rv.Type().Elem().Elem()).Interface()

		if ider, ok := element.(SetIDer); ok {
			ider.SetID(result.Data.ID)
		}

		err := json.Unmarshal(result.Data.Attributes, element)
		if err != nil {
			return errors.Wrap(err, "could not parse attributes json")
		}
		rv = reflect.Append(rv, reflect.ValueOf(element))
	}

	reflect.ValueOf(v).Elem().Set(rv)

	return nil
}

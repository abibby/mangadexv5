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

type RelationshipList []*Relationship

func (l RelationshipList) Get(relationship string) string {
	for _, r := range l {
		if r.Type == relationship {
			return r.ID
		}
	}
	return ""
}

type Modeler interface {
	update(id string, relationships RelationshipList)
}

type Model struct {
	ID            string
	Relationships RelationshipList
}

func (m *Model) update(id string, relationships RelationshipList) {
	m.ID = id
	m.Relationships = relationships
}

var _ Modeler = &Model{}

func (r *PaginatedResponse) loadResults(v interface{}) error {
	rv := reflect.ValueOf(v).Elem()

	for _, result := range r.Results {
		element := reflect.New(rv.Type().Elem().Elem()).Interface()

		if ider, ok := element.(Modeler); ok {
			ider.update(result.Data.ID, result.Relationships)
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

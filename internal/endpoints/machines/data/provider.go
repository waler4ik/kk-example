// Code generated by kk; BUT FEEL FREE TO EDIT.

package data

import (
	"fmt"
	"github.com/waler4ik/kk-example/internal/api"
)

type Provider struct {
}

func NewProvider(a *api.API) Provider {
	return Provider{}
}

// QueryParameter implements github.com/gorilla/schema
type QueryParameter struct {
	//SomeNiceQueryParameter         string `schema:"queryparametername,required"`
}

func (p Provider) GetData(qp QueryParameter) ([]*Datum, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p Provider) GetDatum(id string) (*Datum, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p Provider) CreateDatum(c *Datum) error {
	return fmt.Errorf("not implemented")
}

func (p Provider) UpdateDatum(c *Datum) error {
	return fmt.Errorf("not implemented")
}

func (p Provider) DeleteDatum(id string) error {
	return fmt.Errorf("not implemented")
}
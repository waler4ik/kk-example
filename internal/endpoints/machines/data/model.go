// Code generated by kk; BUT FEEL FREE TO EDIT.

package data

import (
	"net/http"

	"github.com/google/uuid"
)

type Datum struct {
	UUID uuid.UUID
}

func (c *Datum) Bind(r *http.Request) error {
	return nil
}

func (c *Datum) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

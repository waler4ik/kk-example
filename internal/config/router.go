// Code generated by kk; DO NOT EDIT.

package config

import (
	"github.com/waler4ik/kk-example/internal/api"
	"github.com/waler4ik/kk-example/internal/endpoints/machines/data"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ConfigureRouter(a *api.API) http.Handler {
	r := chi.NewRouter()

	ConfigureMiddleware(r)
	data.ConfigureRouter(a, r)

	return r
}

// Code generated by kk; DO NOT EDIT.

package config

import (
	"github.com/waler4ik/kk-example/internal/api"
	"github.com/waler4ik/kk-example/internal/endpoints/ws"
)

func configureWebsockets(a *api.Websockets) {
	//websockets
	a.Ws = ws.NewWebsocket()
}
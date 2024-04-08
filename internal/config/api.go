// Code generated by kk; DO NOT EDIT.

package config

import (
	"fmt"

	"github.com/waler4ik/kk-example/internal/api"
)

func ConfigureAPI(a *api.API) error {
	configureBasicAPI(&a.Basic)
	configureWebsockets(&a.Websockets)
	if err := configureProvider(a.Logger, &a.Provider); err != nil {
		return fmt.Errorf("configureProvider: %w", err)
	}
	if err := configureUserManagedAPI(a.Logger, &a.UserManagedAPI); err != nil {
		return fmt.Errorf("configureUserManagedAPI: %w", err)
	}
	return nil
}
package main

import (
	"context"

	"parking/api/storage/postgres"
	"parking/utility/logging"
)

func Bootstrap(store *postgres.Storage) error {
	log := logging.FromContext(context.Background()).WithField("method", "api.Bootstrap")
	log.Trace("request received")

	if err := BootstrapInitialUsers(store); err != nil {
		log.Info(err)
		return nil
	}
	return nil
}

func BootstrapInitialUsers(store *postgres.Storage) error {
	return nil
}

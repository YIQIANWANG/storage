//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"storage/service"
)

func InitContainer() (*Container, error) {
	wire.Build(
		// container
		NewContainer,

		// service
		service.NewChunkService,
		service.NewHashService,
		service.NewHeartbeatService,
		service.NewSyncService,
	)
	return &Container{}, nil
}

package app

import "storage/service"

var Default *Container

func InitDefault() {
	var err error
	Default, err = InitContainer()
	if err != nil {
		panic(err)
	}
}

type Container struct {
	chunkService     *service.ChunkService
	heartbeatService *service.HeartbeatService
}

func NewContainer(
	chunkService *service.ChunkService,
	heartbeatService *service.HeartbeatService,
) *Container {
	return &Container{
		chunkService:     chunkService,
		heartbeatService: heartbeatService,
	}
}

func (c *Container) GetChunkService() *service.ChunkService {
	return c.chunkService
}

func (c *Container) GetHeartbeatService() *service.HeartbeatService {
	return c.heartbeatService
}

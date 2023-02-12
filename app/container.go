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
	hashService      *service.HashService
	heartbeatService *service.HeartbeatService
}

func NewContainer(
	chunkService *service.ChunkService,
	hashService *service.HashService,
	heartbeatService *service.HeartbeatService,
) *Container {
	return &Container{
		chunkService:     chunkService,
		hashService:      hashService,
		heartbeatService: heartbeatService,
	}
}

func (c *Container) GetChunkService() *service.ChunkService {
	return c.chunkService
}

func (c *Container) GetHashService() *service.HashService {
	return c.hashService
}

func (c *Container) GetHeartbeatService() *service.HeartbeatService {
	return c.heartbeatService
}

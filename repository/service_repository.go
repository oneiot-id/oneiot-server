package repository

import (
	"context"
	"database/sql"
	"oneiot-server/model/entity"
)

type IServiceRepository interface {
	CreateService(ctx context.Context, service entity.Service) (entity.Service, error)
}

type ServiceRepository struct {
	db sql.DB
}

func (s *ServiceRepository) CreateService(ctx context.Context, service entity.Service) (entity.Service, error) {

	return entity.Service{}, nil
}

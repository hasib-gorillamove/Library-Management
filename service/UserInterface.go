package service

import (
	"context"
	"github.com/golden-infotech/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, data entity.UserRegistration) error
	GetAllUser(ctx context.Context, filter entity.UserFilter) ([]entity.UserRegistration, int, error)
	GetAUser(ctx context.Context, id string) (entity.UserRegistration, error)
	UpdateUser(ctx context.Context, data entity.UserRegistration, id string) error
	DeleteUser(ctx context.Context, id string) error
}

package service

import (
	"context"
	"github.com/golden-infotech/entity"
	"github.com/golden-infotech/entity/httpentity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, data httpentity.CreateUserRegistration) error
	GetAllUser(ctx context.Context, filter entity.UserFilter) ([]entity.UserRegistration, int, error)
	GetAUser(ctx context.Context, id string) (entity.UserRegistration, error)
	UpdateUser(ctx context.Context, data entity.UserRegistration, id string) error
	DeleteUser(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (entity.UserRegistration, error)
}

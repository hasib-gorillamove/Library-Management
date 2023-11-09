package service

import (
	"context"
	"github.com/golden-infotech/entity"
)

type WriterRepository interface {
	Create(ctx context.Context, data entity.Writer) error
	ListAllWriter(ctx context.Context, filter entity.WriterFilter) ([]entity.Writer, int, error)
	GetAWriter(ctx context.Context, id string) (entity.Writer, error)
	UpdateWriter(ctx context.Context, data entity.Writer, id string) error
	Delete(ctx context.Context, id string) error
}

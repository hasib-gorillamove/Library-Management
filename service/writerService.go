package service

import (
	"context"
	"github.com/golden-infotech/entity"
)

type WriterService struct {
	WriterRepository WriterRepository
}

func NewWriterService(writerRepository WriterRepository) *WriterService {
	return &WriterService{
		WriterRepository: writerRepository,
	}
}
func (s *WriterService) Create(ctx context.Context, data entity.Writer) error {
	err := s.WriterRepository.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *WriterService) ListAllWriter(ctx context.Context, filter entity.WriterFilter) ([]entity.Writer, int, error) {
	res, count, err := s.WriterRepository.ListAllWriter(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil

}

func (s *WriterService) GetAWriter(ctx context.Context, id string) (entity.Writer, error) {
	res, err := s.WriterRepository.GetAWriter(ctx, id)
	if err != nil {
		return entity.Writer{}, err
	}
	return res, nil
}
func (s *WriterService) UpdateWriter(ctx context.Context, data entity.Writer, id string) error {
	err := s.WriterRepository.UpdateWriter(ctx, data, id)

	if err != nil {
		return err
	}
	return nil
}

func (s *WriterService) Delete(ctx context.Context, id string) error {
	err := s.WriterRepository.Delete(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"
	"errors"
	"github.com/golden-infotech/entity"
	"github.com/uptrace/bun"
)

type WriterRepository struct {
	DB *bun.DB
}

func NewWriterRepository(db *bun.DB) *WriterRepository {
	return &WriterRepository{
		DB: db,
	}
}
func (repo *WriterRepository) Create(ctx context.Context, data entity.Writer) error {
	_, err := repo.DB.NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *WriterRepository) ListAllWriter(ctx context.Context, filter entity.WriterFilter) ([]entity.Writer, int, error) {
	var data []entity.Writer

	query := repo.DB.NewSelect().Model(&data)

	if filter.Keyword != "" {
		query.Where("title Ilike ?", "%"+filter.Keyword+"%")
	}
	count, err := query.Limit(filter.GetLimit()).Offset(filter.GetOffset()).ScanAndCount(ctx)
	if err != nil {
		return []entity.Writer{}, 0, err
	}
	return data, count, nil

}

func (repo *WriterRepository) GetAWriter(ctx context.Context, id string) (entity.Writer, error) {
	var data entity.Writer

	err := repo.DB.NewSelect().Model(&data).Where("author_id =?", id).Relation("Books").Scan(ctx)

	if err != nil {
		return entity.Writer{}, err
	}

	return data, nil
}
func (repo *WriterRepository) UpdateWriter(ctx context.Context, data entity.Writer, id string) error {
	_, err := repo.DB.NewUpdate().Model(&data).
		ExcludeColumn("created_at").
		ExcludeColumn("created_by").
		ExcludeColumn("deleted_at").
		ExcludeColumn("updated_by").
		Set("updated_at=NOW()").
		Set("name=?", data.Name).
		Set("nationality=?", data.Nationality).
		Set("address", data.Address).
		Where("id=?", id).Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (repo *WriterRepository) Delete(ctx context.Context, id string) error {
	var data entity.Writer
	res, err := repo.DB.NewDelete().Model(&data).Where("id=?", id).Exec(ctx)
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no data matched")
	}
	return err

}
